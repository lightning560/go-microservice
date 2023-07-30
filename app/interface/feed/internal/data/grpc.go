package data

import (
	grpcclt "miopkg/grpc/client"
	feedv1 "redbook/api/domain/feed/v1"
	userv1 "redbook/api/domain/user/v1"
)

func NewFeedRpcClient() feedv1.FeedRpcClient {
	conn := grpcclt.StdConfig("directserver.feed").Build()
	client := feedv1.NewFeedRpcClient(conn)
	return client
}

func NewUserRpcClient() userv1.UserRpcClient {
	conn := grpcclt.StdConfig("directserver.user").Build()
	client := userv1.NewUserRpcClient(conn)
	return client
}
