package biz

import (
	"context"
	"fmt"
	"redbook/app/domain/mall/model"
	"redbook/app/domain/mall/model/entity"
)

type ICartRepo interface {
	// 获取购物车
	CreateCartItem(ctx context.Context, uid, productId, collectionId int64, quantity int32) (int64, error)
	ListCartItemByUid(ctx context.Context, uid int64) ([]*entity.CartItem, error)
	UpdateCartItemQuantity(ctx context.Context, id int64, quantity int32) error
	DeleteCartItem(ctx context.Context, id int64) error
}

type CartBiz struct {
	repo ICartRepo
}

func NewCartBiz(repo ICartRepo) *CartBiz {
	return &CartBiz{repo: repo}
}

func (b *CartBiz) CreateCartItem(ctx context.Context, uid, productId, collectionId int64, quantity int32) (int64, error) {
	return b.repo.CreateCartItem(ctx, uid, productId, collectionId, quantity)
}

func (b *CartBiz) UpdateCartItemQuantity(ctx context.Context, id int64, quantity int32) error {
	return b.repo.UpdateCartItemQuantity(ctx, id, quantity)
}

func (b *CartBiz) DeleteCartItem(ctx context.Context, id int64) error {
	return b.repo.DeleteCartItem(ctx, id)
}

func (b *CartBiz) ListCartItemByUid(ctx context.Context, uid int64) (model.CartItems, error) {
	ets, err := b.repo.ListCartItemByUid(ctx, uid)
	fmt.Println("biz#ListCartItemByUid ets:", ets)
	if err != nil {
		return nil, err
	}
	res := model.CartItems{}
	res.ListFromEntity(ets)
	fmt.Println("biz#ListCartItemByUid res:", res)
	return res, nil
}
