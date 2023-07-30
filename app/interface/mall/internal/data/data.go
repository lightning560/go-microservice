package data

import (
	mallv1 "redbook/api/domain/mall/v1"
	userv1 "redbook/api/domain/user/v1"
)

type Data struct {
	mallRpc mallv1.MallRpcClient
	userRpc userv1.UserRpcClient
}

func NewData(mallRpc mallv1.MallRpcClient, userRpc userv1.UserRpcClient) *Data {
	return &Data{
		mallRpc: mallRpc,
		userRpc: userRpc,
	}
}
