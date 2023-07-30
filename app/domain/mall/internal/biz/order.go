package biz

import (
	"context"
	"redbook/app/domain/mall/internal/data/ent"
	"redbook/app/domain/mall/model"
)

type IOrderRepo interface {
	CreateOrder(ctx context.Context, m *model.Order) (int64, error)
	GetOrderById(ctx context.Context, id int64) (*ent.Order, error)
	ListOrderByIds(ctx context.Context, ids []int64) ([]*ent.Order, error)
	ListOrderByUid(ctx context.Context, uid int64, offset, limit int32, sortBy, sortOrder string) ([]*ent.Order, error)
	UpdateOrderState(ctx context.Context, id int64, state int32) error
}

type OrderBiz struct {
	repo IOrderRepo
}

func NewOrderBiz(repo IOrderRepo) *OrderBiz {
	return &OrderBiz{repo: repo}
}

func (b *OrderBiz) CreateOrder(ctx context.Context, m *model.Order) (int64, error) {
	return b.repo.CreateOrder(ctx, m)
}

func (b *OrderBiz) GetOrderById(ctx context.Context, id int64) (*model.Order, error) {
	rv, err := b.repo.GetOrderById(ctx, id)
	if err != nil {
		return nil, err
	}
	req := &model.Order{}
	req.FromEntity(rv)
	return req, nil
}

func (b *OrderBiz) ListOrderByIds(ctx context.Context, ids []int64) ([]*model.Order, error) {
	rv, err := b.repo.ListOrderByIds(ctx, ids)
	if err != nil {
		return nil, err
	}
	var req []*model.Order
	for _, r := range rv {
		m := &model.Order{}
		m.FromEntity(r)
		req = append(req, m)
	}
	return req, nil
}

func (b *OrderBiz) ListOrderByUid(ctx context.Context, uid int64, offset, limit int32, sortBy, sortOrder string) ([]*model.Order, error) {
	rv, err := b.repo.ListOrderByUid(ctx, uid, offset, limit, sortBy, sortOrder)
	if err != nil {
		return nil, err
	}
	var req []*model.Order
	for _, r := range rv {
		m := &model.Order{}
		m.FromEntity(r)
		req = append(req, m)
	}
	return req, nil
}

func (b *OrderBiz) UpdateOrderState(ctx context.Context, id int64, state int32) error {
	return b.repo.UpdateOrderState(ctx, id, state)
}
