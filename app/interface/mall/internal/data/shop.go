package data

import (
	"context"
	"fmt"
	mallv1 "redbook/api/domain/mall/v1"
	"redbook/app/interface/mall/internal/biz"
	"redbook/common/basemodel"

	"google.golang.org/grpc/status"
)

type shopRepo struct {
	data *Data
}

func NewShopRepo(data *Data) biz.IShopRepo {
	return &shopRepo{data: data}
}

func (r *shopRepo) CreateShop(ctx context.Context, uid int64, name, sign string, logo *basemodel.Image) (int64, error) {
	req := &mallv1.CreateShopReq{
		Uid:  uid,
		Name: name,
		Sign: sign,
		Logo: logo.ToProto(),
	}
	rv, err := r.data.mallRpc.CreateShop(ctx, req)
	if err != nil {
		fmt.Println("CreateShop grpc client err:", err)
		if ne, ok := status.FromError(err); ok {
			fmt.Println("CreateShop grpc client status.code:", ne.Code())
			fmt.Println("CreateShop grpc client status.Message:", ne.Message())
		}
		return -1, err
	}
	return rv.Id, nil
}

func (r *shopRepo) GetShopById(ctx context.Context, id int64) (*mallv1.Shop, error) {
	req := &mallv1.GetShopByIdReq{
		Id: id,
	}
	resp, err := r.data.mallRpc.GetShopById(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Shop, nil
}
