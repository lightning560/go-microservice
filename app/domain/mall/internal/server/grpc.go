package server

import (
	"context"
	grpcsvr "miopkg/grpc/server"
	mallv1 "redbook/api/domain/mall/v1"
)

func (eng *Engine) NewGrpcServer() error {
	server := grpcsvr.StdConfig("grpc.mall").MustBuild()
	mallGrpcServer := BuildMallServer(context.Background(), server)
	mallv1.RegisterMallRpcServer(server.Server, mallGrpcServer)
	return eng.Serve(server)
}
