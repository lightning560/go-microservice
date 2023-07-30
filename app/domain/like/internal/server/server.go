package server

import (
	"context"
	"miopkg/application"
	grpcsvr "miopkg/grpc/server"
	"miopkg/log"
	likev1 "redbook/api/domain/like/v1"
	"redbook/app/domain/like/internal/biz"
	"redbook/app/domain/like/internal/data"
)

type Engine struct {
	application.Application
}

func NewEngine() *Engine {
	eng := &Engine{}

	if err := eng.Startup(
		eng.NewGrpcServer,
	); err != nil {
		log.Panic("like grpcsvt startup", log.Any("err", err))
	}
	return eng
}

type LikeServer struct {
	likev1.UnimplementedLikeServer
	server    *grpcsvr.Server
	likebz    *biz.LikeBiz
	dislikebz *biz.DislikeBiz
}

// 这里手动配置组装data的repo和biz、grpc的具体server
func BuildLikeServer(ctx context.Context, svr *grpcsvr.Server) *LikeServer {
	d, _, err := data.NewData(ctx)
	if err != nil {
		log.Error("account grpc NewAccount NewData()")
		return nil
	}
	//new 符合biz接口的repo
	likeRepoInterface := data.NewLikeRepo(d)
	dislikeRepoInterface := data.NewDislikeRepo(d)
	return &LikeServer{
		likebz:    biz.NewLikeBiz(likeRepoInterface),
		dislikebz: biz.NewDislikeBiz(dislikeRepoInterface),
		server:    svr,
	}
}
