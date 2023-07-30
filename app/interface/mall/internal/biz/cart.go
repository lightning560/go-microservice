package biz

import (
	"context"
	"fmt"
	domainmallv1 "redbook/api/domain/mall/v1"
	"redbook/app/interface/mall/model"
)

type ICartRepo interface {
	CreateCartItem(ctx context.Context, uid, collectionId, productId int64, quantity int32) (int64, error)
	// GetCartItem(id int64) (domainmallv1.CartItem, error)
	ListCartItemByUid(ctx context.Context, uid int64) ([]*domainmallv1.CartItem, error)
	UpdateCartItemQuantity(ctx context.Context, id int64, quantity int32) error
	// DeleteCartItem(id int64) error
}

type CartBiz struct {
	repo ICartRepo
}

func NewCartBiz(repo ICartRepo) *CartBiz {
	return &CartBiz{repo: repo}
}

func (b *CartBiz) CreateCartItem(ctx context.Context, uid, collectionId, productId int64, quantity int32) (int64, error) {
	return b.repo.CreateCartItem(ctx, uid, collectionId, productId, quantity)
}

func (b *CartBiz) ListCartItemByUid(ctx context.Context, uid int64) (model.CartItems, error) {
	rv, err := b.repo.ListCartItemByUid(ctx, uid)
	fmt.Println("biz.ListCartItemByUid.rv:", rv[0])
	if err != nil {
		return nil, err
	}
	var cartItems model.CartItems
	cartItems.ListFromDomainProto(rv)
	fmt.Println("biz.ListCartItemByUid.cartItems", cartItems[0])
	return cartItems, nil
}

func (b *CartBiz) UpdateCartItemQuantity(ctx context.Context, id int64, quantity int32) error {
	return b.repo.UpdateCartItemQuantity(ctx, id, quantity)
}
