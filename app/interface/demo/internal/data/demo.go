package data

import (
	"context"
	"miopkg/log"
	"redbook/app/interface/demo/internal/biz"

	demov1 "redbook/api/domain/demo/v1"

	"github.com/davecgh/go-spew/spew"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type demoRepo struct {
	data *Data
}

func NewDemoRepo() biz.IDemoRepo {
	data := NewData(NewDemoRpcClient())
	return &demoRepo{data: data}
}

func (dr demoRepo) CreateDemo(ctx context.Context, demo *biz.Demo) (*biz.Demo, error) {
	var headers metadata.MD
	var trailers metadata.MD
	resp, err := dr.data.demoRpc.SayHello(context.Background(), &demov1.DemoReq{
		Oid: 1,
	}, grpc.Header(&headers), grpc.Trailer(&trailers))
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Info("receive response", log.String("resp", resp.String()))
	}
	spew.Dump(headers)
	spew.Dump(trailers)
	// fmt.println(resp)
	r := &biz.Demo{
		Age:  11,
		Name: "name",
	}
	return r, err
}
