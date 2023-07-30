package server

import (
	"context"
	grpcsvr "miopkg/grpc/server"
	commentv1 "redbook/api/domain/comment/v1"
)

func (eng *Engine) NewGrpcServer() error {
	server := grpcsvr.StdConfig("grpc.comment").MustBuild()
	comment := BuildCommentServer(context.Background(), server)
	commentv1.RegisterCommentRpcServer(server.Server, comment)
	return eng.Serve(server)
}
