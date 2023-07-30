package server

import (
	"context"
	"miopkg/application"
	grpcsvr "miopkg/grpc/server"
	"miopkg/log"
	userv1 "redbook/api/domain/user/v1"
	"redbook/app/domain/user/internal/biz"
	"redbook/app/domain/user/internal/data"
)

type Engine struct {
	application.Application
}

func NewEngine() *Engine {
	eng := &Engine{}

	if err := eng.Startup(
		eng.NewGrpcServer,
		eng.NewServerGovernor,
	); err != nil {
		log.Panic("startup", log.Any("err", err))
	}
	return eng
}

type UserServer struct {
	userv1.UnimplementedUserRpcServer
	server     *grpcsvr.Server
	profilebiz *biz.ProfileBiz
	userbiz    *biz.UserBiz
}

// 这里手动配置组装data的repo和biz、grpc的具体server
func BuildUserServer(ctx context.Context, svr *grpcsvr.Server) *UserServer {
	d, _, err := data.NewData(ctx)
	if err != nil {
		log.Error("user grpc BuildUser NewData()")
		return nil
	}
	//new 符合biz接口的repo
	IProfileRepo := data.NewProfileRepo(d)
	IUserRepo := data.NewUserRepo(d)
	return &UserServer{
		profilebiz: biz.NewProfileBiz(IProfileRepo),
		userbiz:    biz.NewUserBiz(IUserRepo),
		server:     svr,
	}
}
