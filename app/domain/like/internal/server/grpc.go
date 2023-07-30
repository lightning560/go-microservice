package server

import (
	"context"
	grpcsvr "miopkg/grpc/server"
	likev1 "redbook/api/domain/like/v1"
)

func (eng *Engine) NewGrpcServer() error {
	server := grpcsvr.StdConfig("like").MustBuild()
	like := BuildLikeServer(context.Background(), server)
	likev1.RegisterLikeServer(server.Server, like)
	return eng.Serve(server)
}
