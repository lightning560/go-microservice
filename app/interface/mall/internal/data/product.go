package data

import (
	"context"
	domainmallv1 "redbook/api/domain/mall/v1"
	"redbook/app/interface/mall/internal/biz"
	"redbook/app/interface/mall/model"
)

type productRepo struct {
	data *Data
}

func NewProductRepo(data *Data) biz.IProductRepo {
	return &productRepo{data: data}
}

func (r *productRepo) CreateProduct(ctx context.Context, m *model.CreateProduct) (int64, error) {
	rv, err := r.data.mallRpc.CreateProduct(ctx, &domainmallv1.CreateProductReq{
		CreateProduct: m.ToDomainProto(),
	})
	if err != nil {
		return -1, err
	}
	return rv.Id, nil
}

func (r *productRepo) GetProductById(ctx context.Context, id int64) (*domainmallv1.Product, error) {
	resp, err := r.data.mallRpc.GetProductById(ctx, &domainmallv1.GetProductByIdReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	return resp.Product, nil
}

func (r *productRepo) MapProductByIds(ctx context.Context, ids []int64) (map[int64]*domainmallv1.Product, error) {
	resp, err := r.data.mallRpc.MapProductByIds(ctx, &domainmallv1.MapProductByIdsReq{
		Ids: ids,
	})
	if err != nil {
		return nil, err
	}
	return resp.Products, nil
}
