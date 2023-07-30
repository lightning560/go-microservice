package server

import (
	"context"
	"miopkg/application"
	grpcsvr "miopkg/grpc/server"
	"miopkg/log"
	feedv1 "redbook/api/domain/feed/v1"
	"redbook/app/domain/feed/internal/biz"
	"redbook/app/domain/feed/internal/data"
)

type Engine struct {
	application.Application
}

func NewEngine() *Engine {
	eng := &Engine{}

	if err := eng.Startup(
		eng.NewGrpcServer,
	); err != nil {
		log.Panic("feed grpcsvr startup", log.Any("err", err))
	}
	return eng
}

type FeedServer struct {
	feedv1.UnimplementedFeedRpcServer
	server      *grpcsvr.Server
	feedbiz     *biz.FeedBiz
	relationbiz *biz.RelationBiz
	likebiz     *biz.LikeBiz
	favorbiz    *biz.FavorBiz
	sharebiz    *biz.ShareBiz
}

func BuildFeedServer(ctx context.Context, svr *grpcsvr.Server) *FeedServer {
	d, _, err := data.NewData(ctx)
	if err != nil {
		log.Error("feed grpc BuildFeedServer NewData()")
		return nil
	}
	IFeedRepo := data.NewFeedRepo(d)
	IRelationRepo := data.NewRelationRepo(d)
	ILikeRepo := data.NewLikeRepo(d)
	IFavorRepo := data.NewFavorRepo(d)
	IShareRepo := data.NewShareRepo(d)
	return &FeedServer{
		server:      svr,
		feedbiz:     biz.NewFeedBiz(IFeedRepo),
		relationbiz: biz.NewRelationBiz(IRelationRepo),
		likebiz:     biz.NewLikeBiz(ILikeRepo),
		favorbiz:    biz.NewFavorBiz(IFavorRepo),
		sharebiz:    biz.NewShareBiz(IShareRepo),
	}
}
