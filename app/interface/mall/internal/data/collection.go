package data

import (
	"context"
	domainmallv1 "redbook/api/domain/mall/v1"
	basev1 "redbook/api/redbookpb/v1"
	"redbook/app/interface/mall/internal/biz"
	"redbook/app/interface/mall/model"
)

type collectionRepo struct {
	data *Data
}

func NewCollectionRepo(data *Data) biz.ICollectionRepo {
	return &collectionRepo{data: data}
}

func (r *collectionRepo) CreateCollection(ctx context.Context, m *model.CreateCollection) (int64, error) {
	rv, err := r.data.mallRpc.CreateCollection(ctx, &domainmallv1.CreateCollectionReq{
		CreateCollection: m.ToDomainProto(),
	})
	if err != nil {
		return -1, err
	}
	return rv.Id, nil
}

func (r *collectionRepo) GetCollectionById(ctx context.Context, id int64) (*domainmallv1.Collection, error) {
	resp, err := r.data.mallRpc.GetCollectionById(ctx, &domainmallv1.GetCollectionByIdReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	return resp.Collection, nil
}

func (r *collectionRepo) GetCollectionCardById(ctx context.Context, id int64) (*domainmallv1.CollectionCard, error) {
	resp, err := r.data.mallRpc.GetCollectionCardById(ctx, &domainmallv1.GetCollectionCardByIdReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	return resp.CollectionCard, nil
}

func (r *collectionRepo) ListCollectionCardByIds(ctx context.Context, ids []int64) ([]*domainmallv1.CollectionCard, error) {
	resp, err := r.data.mallRpc.MapCollectionCardByIds(ctx, &domainmallv1.MapCollectionCardByIdsReq{
		Ids: ids,
	})
	if err != nil {
		return nil, err
	}
	res := make([]*domainmallv1.CollectionCard, 0, len(resp.CollectionCards))
	for _, v := range resp.CollectionCards {
		res = append(res, v)
	}
	return res, nil
}

func (r *collectionRepo) ListCollectionCardByShopId(ctx context.Context, id int64, offset, limit int32, by, order string) ([]*domainmallv1.CollectionCard, error) {
	resp, err := r.data.mallRpc.MapCollectionCardByShopId(ctx, &domainmallv1.MapCollectionCardByShopIdReq{
		Id: id,
		Cursor: &basev1.Cursor{
			Offset: offset,
			Limit:  limit,
		},
		Sort: &basev1.Sort{
			By:    by,
			Order: order,
		},
	})
	if err != nil {
		return nil, err
	}
	res := make([]*domainmallv1.CollectionCard, 0, len(resp.CollectionCards))
	for _, v := range resp.CollectionCards {
		res = append(res, v)
	}
	return res, nil
}

func (r *collectionRepo) UpdateCollectionSku(ctx context.Context, id int64, skusOnlyId *model.SkusOnlyId) error {
	_, err := r.data.mallRpc.UpdateCollectionSku(ctx, &domainmallv1.UpdateCollectionSkuReq{
		Id:         id,
		SkusOnlyId: skusOnlyId.ListToDomainProto(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *collectionRepo) UpdateCollectionState(ctx context.Context, id int64, state int32) error {
	_, err := r.data.mallRpc.UpdateCollectionState(ctx, &domainmallv1.UpdateCollectionStateReq{
		Id:    id,
		State: state,
	})
	if err != nil {
		return err
	}
	return nil
}
