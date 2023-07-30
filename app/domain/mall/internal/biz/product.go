package biz

import (
	"context"
	"redbook/app/domain/mall/model"
	"redbook/app/domain/mall/model/entity"
)

type IProductRepo interface {
	CreateProduct(ctx context.Context, m *model.CreateProduct) (int64, error)
	GetProductById(ctx context.Context, id int64) (*entity.Product, error)
	ListProductByIds(ctx context.Context, ids []int64) ([]*entity.Product, error)
	UpdateProductState(ctx context.Context, productId int64, state int32) error
	// DeleteProduct(ctx context.Context, id int64) error
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
	var rv *entity.Product
	if rv, err = b.repo.GetProductById(ctx, id); err != nil {
		return nil, err
	}
	res = &model.Product{}
	res.FromEntity(rv)
	return res, nil
}
func (b *ProductBiz) ListProductByIds(ctx context.Context, ids []int64) (res model.Products, err error) {
	var rv []*entity.Product
	if rv, err = b.repo.ListProductByIds(ctx, ids); err != nil {
		return nil, err
	}
	res = model.Products{}
	res.ListFromEntity(rv)
	return res, nil
}

func (b *ProductBiz) UpdateProductState(ctx context.Context, productId int64, state int32) error {
	return b.repo.UpdateProductState(ctx, productId, state)
}
