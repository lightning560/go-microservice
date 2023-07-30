package data

import (
	grpcclt "miopkg/grpc/client"
	demov1 "redbook/api/domain/demo/v1"
)

type Data struct {
	demoRpc demov1.GreeterClient
}

func NewData(demoRpc demov1.GreeterClient) *Data {
	return &Data{
		demoRpc: demoRpc,
	}
}
func NewDemoRpcClient() demov1.GreeterClient {
	conn := grpcclt.StdConfig("directserver").Build()
	client := demov1.NewGreeterClient(conn)
	return client
}
