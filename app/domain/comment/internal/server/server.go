package server

import (
	"context"
	"miopkg/application"
	grpcsvr "miopkg/grpc/server"
	"miopkg/log"
	commentv1 "redbook/api/domain/comment/v1"
	"redbook/app/domain/comment/internal/biz"
	"redbook/app/domain/comment/internal/data"
)

type Engine struct {
	application.Application
}

func NewEngine() *Engine {
	eng := &Engine{}

	if err := eng.Startup(
		eng.NewGrpcServer,
	); err != nil {
		log.Panic("comment grpcsvt startup", log.Any("err", err))
	}
	return eng
}

type CommentServer struct {
	commentv1.UnimplementedCommentRpcServer
	server     *grpcsvr.Server
	replybiz   *biz.ReplyBiz
	subjectbiz *biz.SubjectBiz
	likebiz    *biz.LikeBiz
}

func BuildCommentServer(ctx context.Context, svr *grpcsvr.Server) *CommentServer {

	mongoDB := data.NewMongoClient(ctx, "comment")
	replyMG := data.NewMongoCollection(ctx, mongoDB, "reply")
	subjectMG := data.NewMongoCollection(ctx, mongoDB, "subject")
	likeMG := data.NewMongoCollection(ctx, mongoDB, "like")
	redisClient := data.NewRedisClient(ctx)

	fanout := data.NewFanout()
	logger := log.MioLogger.With(log.FieldMod("domain.comment"))

	d, _, err := data.NewData(ctx, replyMG, subjectMG, likeMG, redisClient, fanout, logger)
	if err != nil {
		log.Error("comment grpc BuildCommentServer NewData()")
		return nil
	}
	replyRepo := data.NewReplyRepo(d)
	subjectRepo := data.NewSubjectRepo(d)
	likeRepo := data.NewLikeRepo(d)
	return &CommentServer{
		replybiz:   biz.NewReplyBiz(replyRepo),
		subjectbiz: biz.NewSubjectBiz(subjectRepo),
		likebiz:    biz.NewLikeBiz(likeRepo),
		server:     svr,
	}
}
