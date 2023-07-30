package biz

import (
	"context"
	"fmt"
	"redbook/app/domain/mall/model"
	"redbook/app/domain/mall/model/entity"
	"redbook/common/basemodel"
)

type IShopRepo interface {
	CreateShop(ctx context.Context, uid int64, name, sign string, logo *basemodel.Image) (int64, error)
	GetShopById(ctx context.Context, id int64) (*entity.Shop, error)
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
func (b *ShopBiz) GetShopById(ctx context.Context, id int64) (res *model.Shop, err error) {
	var rv *entity.Shop
	if rv, err = b.repo.GetShopById(ctx, id); err != nil {
		return nil, err
	}
	res = &model.Shop{}
	fmt.Println("GetShopById rv:", rv)
	res.FromEntity(rv)
	return res, nil
}
