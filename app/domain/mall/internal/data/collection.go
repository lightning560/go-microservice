package data

import (
	"context"
	"fmt"
	"miopkg/util/snowflake"
	"redbook/app/domain/mall/internal/biz"
	"redbook/app/domain/mall/model"
	"redbook/app/domain/mall/model/entity"
	"time"

	"miopkg/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type collectionRepo struct {
	data *Data
}

func NewCollectionRepo(data *Data) biz.ICollectionRepo {
	return &collectionRepo{data: data}
}

func (r *collectionRepo) CreateCollection(ctx context.Context, m *model.CreateCollection) (int64, error) {
	id := snowflake.GenInt64()
	et := entity.Collection{
		Id:        id,
		State:     1,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),

		ShopId:     m.ShopId,
		Name:       m.Name,
		Title:      m.Title,
		Cover:      m.Cover,
		Video:      m.Video,
		SkusOnlyId: m.SkusOnlyId.ListToEntity(),
		Tags:       m.Tags,
	}
	_, err := r.data.collectionMG.InsertOne(ctx, et)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (r *collectionRepo) GetCollectionById(ctx context.Context, id int64) (res *entity.Collection, err error) {
	sr := r.data.collectionMG.FindOne(ctx, bson.M{"id": id})
	if sr.Err() != nil {
		if sr.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrapf(sr.Err(), "GetCollectionById failed, id: %v", id)
	}
	res = new(entity.Collection)
	if err = sr.Decode(&res); err != nil {
		return nil, err
	}
	return
}

func (r *collectionRepo) ListCollectionByIds(ctx context.Context, ids []int64) ([]*entity.Collection, error) {
	fmt.Println("ids:", ids)
	cursor, err := r.data.collectionMG.Find(ctx, bson.M{"id": bson.M{"$in": ids}})
	if err != nil {
		return nil, err
	}
	res := make([]*entity.Collection, 0)
	err = cursor.All(ctx, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *collectionRepo) ListCollectionByShopId(ctx context.Context, shopId int64, state int32, offset, limit int32, sortBy, sortOrder string) ([]*entity.Collection, error) {

	skip := int64(offset)
	lm := int64(limit)
	sort := int32(1)
	if sortOrder == "desc" {
		sort = -1
	}
	options := &options.FindOptions{
		Sort: bson.D{
			bson.E{Key: sortBy, Value: sort},
		},
		Skip:  &skip,
		Limit: &lm,
	}
	cursor, err := c.data.collectionMG.Find(ctx, bson.M{"shop_id": shopId, "state": 1}, options)
	if err != nil {
		return nil, err
	}
	res := make([]*entity.Collection, 0)
	err = cursor.All(ctx, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *collectionRepo) UpdateCollectionState(ctx context.Context, id int64, state int32) error {
	rv := r.data.collectionMG.FindOneAndUpdate(ctx, bson.M{"id": id}, bson.M{"$set": bson.M{"state": state}})
	if rv.Err() != nil {
		return rv.Err()
	}
	return nil
}

func (r *collectionRepo) UpdateCollectionItem(ctx context.Context, m *model.Collection) error {
	rv := r.data.collectionMG.FindOneAndUpdate(ctx, bson.M{"id": m.Id}, bson.M{"$set": bson.M{"skus_only_id": m.Skus, "updated_at": time.Now().Unix()}})
	if rv.Err() != nil {
		return errors.Wrapf(errors.FromError(rv.Err()), "UpdateCollectionItem failed, id: %v", m.Id)
	}
	return nil
}
