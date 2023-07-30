package data

import (
	grpcclt "miopkg/grpc/client"
	mallv1 "redbook/api/domain/mall/v1"
	userv1 "redbook/api/domain/user/v1"
)

func NewMallRpcClient() mallv1.MallRpcClient {
	conn := grpcclt.StdConfig("directserver.mall").Build()
	client := mallv1.NewMallRpcClient(conn)
	return client
}

func NewUserRpcClient() userv1.UserRpcClient {
	conn := grpcclt.StdConfig("directserver.user").Build()
	client := userv1.NewUserRpcClient(conn)
	return client
}
