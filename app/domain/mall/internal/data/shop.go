package data

import (
	"context"
	"fmt"
	"miopkg/errors"
	"miopkg/util/snowflake"
	"time"

	"redbook/app/domain/mall/internal/biz"
	"redbook/app/domain/mall/model"
	"redbook/app/domain/mall/model/entity"
	"redbook/common/basemodel"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type shopRepo struct {
	data *Data
}

func NewShopRepo(data *Data) biz.IShopRepo {
	return &shopRepo{data: data}
}
func (r *shopRepo) CreateShop(ctx context.Context, uid int64, name, sign string, logo *basemodel.Image) (int64, error) {
	existed, err := r.IsShopExistsByUid(ctx, uid)
	if err != nil {
		return -1, err
	}
	if existed {
		return -1, errors.Wrapf(errors.ErrAlreadyExists, "CreateShop IsShopExistsByUid failed, uid: %v", uid)
	}
	id := snowflake.GenInt64()
	shop := &entity.Shop{
		Id:   id,
		Uid:  uid,
		Name: name,
		Sign: sign,
		Logo: logo,

		State:     1,
		CreatedAt: time.Now().Unix(),
	}
	_, err = r.data.shopMG.InsertOne(ctx, shop)
	if err != nil {
		return -1, errors.Wrapf(errors.FromError(err), "CreateShop failed, uid: %v", uid)
	}
	return id, nil
}

func (r *shopRepo) IsShopExistsByUid(ctx context.Context, uid int64) (bool, error) {
	rv := r.data.shopMG.FindOne(ctx, bson.M{"uid": uid})
	if rv.Err() != nil {
		if rv.Err() == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, errors.Wrapf(errors.FromError(rv.Err()), "IsShopExistsByUid failed, uid: %v", uid)
	}
	return true, nil
}

func (r *shopRepo) UpdateShop(ctx context.Context, m *model.Shop) (int64, error) {
	et := m.ToEntity()
	rv, err := r.data.shopMG.UpdateOne(ctx, bson.M{"id": m.Id}, bson.M{"$set": et})
	if err != nil {
		return -1, err
	}
	return rv.ModifiedCount, nil
}

func (r *shopRepo) GetShopById(ctx context.Context, id int64) (res *entity.Shop, err error) {
	rv := r.data.shopMG.FindOne(ctx, bson.M{"id": id})
	if rv.Err() != nil {
		if rv.Err() == mongo.ErrNoDocuments {
			return nil, errors.Wrapf(errors.ErrNoDocuments, "GetShopById failed, id: %v", id)
		}
		return nil, errors.Wrapf(errors.FromError(rv.Err()), "GetShopById failed, id: %v", id)
	}
	res = &entity.Shop{}
	if err = rv.Decode(res); err != nil {
		fmt.Println(err)
		return nil, errors.Wrapf(errors.FromError(err), "GetShopById Decode failed, id: %v", id)
	}
	return
}
