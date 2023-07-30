package server

import (
	"context"
	grpcsvr "miopkg/grpc/server"
	feedv1 "redbook/api/domain/feed/v1"
)

func (eng *Engine) NewGrpcServer() error {
	server := grpcsvr.StdConfig("grpc.feed").MustBuild()
	feed := BuildFeedServer(context.Background(), server)
	feedv1.RegisterFeedRpcServer(server.Server, feed)
	return eng.Serve(server)
}
