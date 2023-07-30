package biz

import (
	"context"
	domainmallv1 "redbook/api/domain/mall/v1"
	"redbook/app/interface/mall/model"
)

type ICollectionRepo interface {
	// CreateCollection 创建收藏
	CreateCollection(ctx context.Context, collection *model.CreateCollection) (int64, error)
	// DeleteCollection 删除收藏
	// DeleteCollection(ctx context.Context, collection *model.Collection) error
	// GetCollection 获取收藏
	GetCollectionById(ctx context.Context, id int64) (*domainmallv1.Collection, error)
	GetCollectionCardById(ctx context.Context, id int64) (*domainmallv1.CollectionCard, error)
	// ListCollections 获取收藏列表
	ListCollectionCardByIds(ctx context.Context, Ids []int64) ([]*domainmallv1.CollectionCard, error)
	ListCollectionCardByShopId(ctx context.Context, shopId int64, offset, limit int32, by, order string) ([]*domainmallv1.CollectionCard, error)
	// UpdateCollection 更新收藏
	UpdateCollectionSku(ctx context.Context, id int64, skusOnlyId *model.SkusOnlyId) error
	UpdateCollectionState(ctx context.Context, id int64, state int32) error
}

type CollectionBiz struct {
	repo ICollectionRepo
}

func NewCollectionBiz(repo ICollectionRepo) *CollectionBiz {
	return &CollectionBiz{repo: repo}
}

func (c *CollectionBiz) CreateCollection(ctx context.Context, createCollection *model.CreateCollection) (int64, error) {
	return c.repo.CreateCollection(ctx, createCollection)
}

func (c *CollectionBiz) GetCollectionById(ctx context.Context, id int64) (*model.Collection, error) {
	rv, err := c.repo.GetCollectionById(ctx, id)
	if err != nil {
		return nil, err
	}
	res := &model.Collection{}
	res.FromDomainProto(rv)
	return res, nil
}

func (c *CollectionBiz) GetCollectionCardById(ctx context.Context, id int64) (*model.CollectionCard, error) {
	rv, err := c.repo.GetCollectionCardById(ctx, id)
	if err != nil {
		return nil, err
	}
	res := &model.CollectionCard{}
	res.FromDomainProto(rv)
	return res, nil
}

func (c *CollectionBiz) ListCollectionCardByIds(ctx context.Context, ids []int64) (*model.CollectionCards, error) {
	rv, err := c.repo.ListCollectionCardByIds(ctx, ids)
	if err != nil {
		return nil, err
	}
	res := &model.CollectionCards{}
	res.ListFromDomainProto(rv)
	return res, nil
}

func (c *CollectionBiz) ListCollectionCardByShopId(ctx context.Context, shopId int64, offset, limit int32, by, order string) (*model.CollectionCards, error) {
	rv, err := c.repo.ListCollectionCardByShopId(ctx, shopId, offset, limit, by, order)
	if err != nil {
		return nil, err
	}
	res := &model.CollectionCards{}
	res.ListFromDomainProto(rv)
	return res, nil
}

func (c *CollectionBiz) UpdateCollectionSku(ctx context.Context, id int64, skusOnlyId *model.SkusOnlyId) error {
	return c.repo.UpdateCollectionSku(ctx, id, skusOnlyId)
}

func (c *CollectionBiz) UpdateCollectionState(ctx context.Context, id int64, state int32) error {
	return c.repo.UpdateCollectionState(ctx, id, state)
}
