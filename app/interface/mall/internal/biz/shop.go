package biz

import (
	"context"
	domainmallv1 "redbook/api/domain/mall/v1"
	"redbook/app/interface/mall/model"
	"redbook/common/basemodel"
)

type IShopRepo interface {
	CreateShop(ctx context.Context, uid int64, name, sign string, logo *basemodel.Image) (int64, error)
	GetShopById(ctx context.Context, id int64) (*domainmallv1.Shop, error)
}

type ShopBiz struct {
	repo IShopRepo
}

func NewShopBiz(repo IShopRepo) *ShopBiz {
	return &ShopBiz{repo: repo}
}
func (b *ShopBiz) CreateShop(ctx context.Context, uid int64, name, sign string, logo *basemodel.Image) (int64, error) {
	return b.repo.CreateShop(ctx, uid, name, sign, logo)
}
func (b *ShopBiz) GetShopById(ctx context.Context, id int64) (*model.Shop, error) {
	rv, err := b.repo.GetShopById(ctx, id)
	if err != nil {
		return nil, err
	}
	var shop model.Shop
	shop.FromDomainProto(rv)
	return &shop, nil
}
