package biz

import (
	"context"
	"redbook/app/domain/mall/model"
	"redbook/app/domain/mall/model/entity"
)

type ICollectionRepo interface {
	CreateCollection(ctx context.Context, m *model.CreateCollection) (int64, error)
	GetCollectionById(ctx context.Context, id int64) (*entity.Collection, error)
	ListCollectionByIds(ctx context.Context, ids []int64) ([]*entity.Collection, error)
	ListCollectionByShopId(ctx context.Context, shopId int64, state int32, offset, limit int32, by, order string) ([]*entity.Collection, error)
	UpdateCollectionState(ctx context.Context, id int64, state int32) error
	UpdateCollectionItem(ctx context.Context, m *model.Collection) error
}
type CollectionBiz struct {
	repo ICollectionRepo
}

func NewCollectionBiz(repo ICollectionRepo) *CollectionBiz {
	return &CollectionBiz{repo: repo}
}

func (b *CollectionBiz) CreateCollection(ctx context.Context, m *model.CreateCollection) (int64, error) {
	return b.repo.CreateCollection(ctx, m)
}

func (b *CollectionBiz) GetCollectionById(ctx context.Context, id int64) (res *model.Collection, err error) {
	et, err := b.repo.GetCollectionById(ctx, id)
	if err != nil {
		return nil, err
	}
	res = new(model.Collection)
	res.FromEntity(et)
	return
}

func (b *CollectionBiz) ListCollectionByIds(ctx context.Context, ids []int64) (res model.Collections, err error) {
	ets, err := b.repo.ListCollectionByIds(ctx, ids)
	if err != nil {
		return nil, err
	}
	res = model.Collections{}
	res.ListFromEntity(ets)
	return res, nil
}

func (b *CollectionBiz) ListCollectionByShopId(ctx context.Context, shopId int64, state int32, offset, limit int32, by, order string) (res model.Collections, err error) {
	ets, err := b.repo.ListCollectionByShopId(ctx, shopId, state, offset, limit, by, order)
	if err != nil {
		return nil, err
	}
	res = model.Collections{}
	res.ListFromEntity(ets)
	return
}

func (b *CollectionBiz) UpdateCollectionState(ctx context.Context, id int64, state int32) error {
	return b.repo.UpdateCollectionState(ctx, id, state)
}

func (b *CollectionBiz) UpdateCollectionItem(ctx context.Context, m *model.Collection) error {
	return b.repo.UpdateCollectionItem(ctx, m)
}
