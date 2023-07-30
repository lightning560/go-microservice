package biz

import (
	"context"
	domainmallv1 "redbook/api/domain/mall/v1"
	"redbook/app/interface/mall/model"
)

type IProductRepo interface {
	CreateProduct(ctx context.Context, m *model.CreateProduct) (int64, error)
	GetProductById(ctx context.Context, id int64) (*domainmallv1.Product, error)
	MapProductByIds(ctx context.Context, ids []int64) (map[int64]*domainmallv1.Product, error)
}

type ProductBiz struct {
	repo IProductRepo
}

func NewProductBiz(repo IProductRepo) *ProductBiz {
	return &ProductBiz{repo: repo}
}
func (b *ProductBiz) CreateProduct(ctx context.Context, m *model.CreateProduct) (int64, error) {
	return b.repo.CreateProduct(ctx, m)
}
func (b *ProductBiz) GetProductById(ctx context.Context, id int64) (res *model.Product, err error) {
	rv, err := b.repo.GetProductById(ctx, id)
	if err != nil {
		return nil, err
	}
	res = new(model.Product)
	res.FromDomainProto(rv)
	return
}
func (b *ProductBiz) MapProductByIds(ctx context.Context, ids []int64) (map[int64]*model.Product, error) {
	rv, err := b.repo.MapProductByIds(ctx, ids)
	if err != nil {
		return nil, err
	}
	res := make(map[int64]*model.Product)

	for k, v := range rv {
		m := &model.Product{}
		m.FromDomainProto(v)
		res[k] = m
	}
	return res, nil
}
