package data

import (
	"context"
	"miopkg/errors"
	"miopkg/util/snowflake"
	"redbook/app/domain/mall/internal/biz"
	"redbook/app/domain/mall/model"
	"redbook/app/domain/mall/model/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type productRepo struct {
	data *Data
}

func NewProductRepo(data *Data) biz.IProductRepo {
	return &productRepo{data: data}
}
func (r *productRepo) CreateProduct(ctx context.Context, m *model.CreateProduct) (int64, error) {
	id := snowflake.GenInt64()
	et := entity.Product{
		Id:        id,
		State:     1,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),

		ShopId:    m.ShopId,
		Title:     m.Title,
		Name:      m.Name,
		Price:     m.Price,
		Inventory: m.Inventory,
		Thumb:     m.Thumb,
		Images:    m.Images,
		Video:     m.Video,
		Overview:  m.Overview,
		Specs:     m.Specs,
		Tags:      m.Tags,
	}

	_, err := r.data.productMG.InsertOne(ctx, et)
	if err != nil {
		return -1, errors.Wrapf(errors.FromError(err), "CreateProduct failed, entity: %v", et)
	}
	return id, nil
}

func (r *productRepo) GetProductById(ctx context.Context, id int64) (res *entity.Product, err error) {
	rv := r.data.productMG.FindOne(ctx, bson.M{"id": id})
	if rv.Err() != nil {
		if rv.Err() == mongo.ErrNoDocuments {
			return nil, errors.Wrapf(errors.ErrNoDocuments, "GetProductById failed, id: %v", id)
		}
	}
	res = &entity.Product{}
	if err = rv.Decode(&res); err != nil {
		return nil, errors.Wrapf(errors.FromError(err), "GetProductById failed, id: %v", id)
	}
	return
}

func (r *productRepo) ListProductByIds(ctx context.Context, ids []int64) (res []*entity.Product, err error) {
	rv, err := r.data.productMG.Find(ctx, bson.M{"id": bson.M{"$in": ids}})
	if err != nil {
		return nil, errors.Wrapf(errors.FromError(err), "ListProductByIds failed, ids: %v", ids)
	}
	res = make([]*entity.Product, 0, len(ids))
	if err = rv.All(ctx, &res); err != nil {
		return nil, errors.Wrapf(errors.FromError(err), "ListProductByIds Decode failed, ids: %v", ids)
	}
	return
}
func (r *productRepo) UpdateProductState(ctx context.Context, productId int64, state int32) error {
	sr := r.data.productMG.FindOneAndUpdate(ctx, bson.M{"id": productId}, bson.M{"$set": bson.M{"state": state}})
	if sr.Err() != nil {
		return errors.Wrapf(errors.FromError(sr.Err()), "UpdateProductState failed, productId: %v, state: %v", productId, state)
	}
	return nil
}
