package data

import (
	"context"
	"fmt"
	"redbook/app/domain/mall/internal/biz"
	"redbook/app/domain/mall/model/entity"
	"time"

	"miopkg/errors"
	"miopkg/util/snowflake"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type cartRepo struct {
	data *Data
}

func NewCartRepo(data *Data) biz.ICartRepo {
	return &cartRepo{data: data}
}

func (r *cartRepo) CreateCartItem(ctx context.Context, uid, productId, collectionId int64, quantity int32) (int64, error) {
	exists, err := r.IsItemExists(ctx, uid, productId)
	if err != nil {
		return -1, errors.Wrapf(err, "CreateCartItem invoke IsItemExists failed, uid: %v, productId: %v", uid, productId)
	}
	if exists {
		return -1, errors.Wrapf(errors.ErrAlreadyExists, "CreateCartItem failed, uid: %v, productId: %v", uid, productId)
	}
	id := snowflake.GenInt64()
	item := &entity.CartItem{
		Id:           id,
		Uid:          uid,
		CollectionId: collectionId,
		ProductId:    productId,
		Quantity:     quantity,
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
	}
	_, err = r.data.cartMG.InsertOne(ctx, item)
	if err != nil {
		return -1, errors.Wrapf(errors.FromError(err), "CreateCartItem failed, item: %v", item)
	}
	return id, nil
}

func (r *cartRepo) IsItemExists(ctx context.Context, uid int64, productId int64) (bool, error) {
	sr := r.data.cartMG.FindOne(ctx, bson.M{"uid": uid, "product_id": productId})
	if sr.Err() != nil {
		if sr.Err() == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, errors.Wrapf(errors.FromError(sr.Err()), "data#IsItemExists failed, uid: %v, productId: %v", uid, productId)
	}
	return true, nil
}

func (r *cartRepo) UpdateCartItemQuantity(ctx context.Context, id int64, quantity int32) (err error) {
	if _, err = r.data.cartMG.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": bson.M{"quantity": quantity, "updated_at": time.Now().Unix()}}); err != nil {
		return errors.Wrapf(errors.FromError(err), "UpdateCartItemQuantity failed, id: %v, quantity: %v", id, quantity)
	}
	return nil
}

func (r *cartRepo) DeleteCartItem(ctx context.Context, id int64) (err error) {
	if _, err = r.data.cartMG.DeleteOne(ctx, bson.M{"id": id}); err != nil {
		return errors.Wrapf(errors.FromError(err), "DeleteCartItem failed, id: %v", id)
	}
	return nil
}

func (r *cartRepo) ListCartItemByUid(ctx context.Context, uid int64) ([]*entity.CartItem, error) {
	cur, err := r.data.cartMG.Find(ctx, bson.M{"uid": uid})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrapf(errors.ErrNoDocuments, "ListCartItemByUser failed, uid: %v", uid)
		}
		return nil, errors.Wrapf(errors.FromError(err), "ListCartItemByUser failed, uid: %v", uid)
	}
	// defer cur.Close(ctx)
	res := make([]*entity.CartItem, 0)
	for cur.Next(ctx) {
		var item entity.CartItem
		if err := cur.Decode(&item); err != nil {
			fmt.Println("ListCartItemByUid cursor Decode err: ", err)
			return nil, errors.Wrapf(errors.FromError(err), "ListCartItemByUser Decode failed, uid: %v", uid)
		}
		fmt.Println("cursor &item: ", item)
		res = append(res, &item)
	}
	fmt.Println("ListCartItemByUid res: ", res)
	return res, nil
}
