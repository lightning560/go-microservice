package server

import (
	"context"
	grpcsvr "miopkg/grpc/server"
	userv1 "redbook/api/domain/user/v1"
)

func (eng *Engine) NewGrpcServer() error {
	server := grpcsvr.StdConfig("grpc.user").MustBuild()
	user := BuildUserServer(context.Background(), server)
	userv1.RegisterUserRpcServer(server.Server, user)
	return eng.Serve(server)
}
