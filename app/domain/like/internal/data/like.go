package data

import (
	"context"
	"fmt"
	"miopkg/errors"
	"miopkg/log"
	"time"

	"redbook/app/domain/like/internal/biz"

	"redbook/app/domain/like/internal/data/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type likeRepo struct {
	data *Data
}

func NewLikeRepo(data *Data) biz.ILikeRepo {
	return &likeRepo{data: data}
}

func (r *likeRepo) AddLike(ctx context.Context, m *model.Like) error {
	_, err := r.Like(ctx, m)
	if errors.Is(err, errors.ErrNoDocuments) {
		rv, err := r.data.likeMG.InsertOne(ctx, bson.M{"mid": m.Mid, "oid": m.Oid, "bid": m.Bid, "like": 1, "create_at": time.Now().Unix(), "updated_at": time.Now().Unix()})
		if err != nil {
			r.data.logger.Errorf("AddLike InsertOne", log.Any("error", err))
			return err
		}
		fmt.Println("AddLike,InsertOne succ:", rv)
	} else {
		rv := r.data.likeMG.FindOneAndUpdate(ctx, bson.M{"mid": m.Mid, "oid": m.Oid}, bson.M{"$set": bson.M{"like": 1, "create_at": time.Now().Unix(), "updated_at": time.Now().Unix()}})
		if rv.Err() != nil {
			r.data.logger.Errorf("AddLike InsertOne", log.Any("error", rv.Err()))
			return rv.Err()
		}
		fmt.Println("AddLike,FindOneAndUpdate succ:", rv)
	}
	return nil
}

func (r *likeRepo) CancelLike(ctx context.Context, m *model.Like) error {
	rv := r.data.likeMG.FindOneAndUpdate(ctx, bson.M{"mid": m.Mid, "oid": m.Oid}, bson.M{"$set": bson.M{"like": 0, "updated_at": time.Now().Unix()}})
	if rv.Err() != nil {
		r.data.logger.Errorf("CancelLike InsertOne", log.Any("error", rv.Err()))
		return rv.Err()
	}
	fmt.Println("CancelLike,FindOneAndUpdate succ:", rv)
	return nil
}

// Like 查询是否点赞 query like by mid and oid
func (r *likeRepo) Like(ctx context.Context, m *model.Like) (res *model.Like, err error) {
	err = r.data.likeMG.FindOne(ctx, bson.M{"mid": m.Mid, "oid": m.Oid}).Decode(&res)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if err == mongo.ErrNoDocuments {
			r.data.logger.Errorf("Like InsertOne", log.Any("error", err))
			return res, errors.ErrNoDocuments
		}
		r.data.logger.Errorf("Like InsertOne", log.Any("error", err))
	}
	return
}
