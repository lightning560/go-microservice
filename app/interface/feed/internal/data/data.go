package data

import (
	feedv1 "redbook/api/domain/feed/v1"
	userv1 "redbook/api/domain/user/v1"
)

type Data struct {
	feedRpc feedv1.FeedRpcClient
	userRpc userv1.UserRpcClient
}

func NewData(feedRpc feedv1.FeedRpcClient, userRpc userv1.UserRpcClient) *Data {
	return &Data{
		feedRpc: feedRpc,
		userRpc: userRpc,
	}
}
