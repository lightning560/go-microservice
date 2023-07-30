package data

import (
	"context"
	"fmt"
	"miopkg/errors"
	"redbook/app/domain/like/internal/biz"
	"redbook/app/domain/like/internal/data/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type dislikeRepo struct {
	data *Data
}

func NewDislikeRepo(data *Data) biz.IDislikeRepo {
	return &dislikeRepo{data: data}
}
func (r *dislikeRepo) AddDislike(ctx context.Context, m *model.Dislike) error {
	_, err := r.Dislike(ctx, m)
	if errors.Is(err, errors.ErrNoDocuments) {
		rv, err := r.data.likeMG.InsertOne(ctx, bson.M{"mid": m.Mid, "oid": m.Oid, "bid": m.Bid, "sid": m.Sid, "dislike": 1, "create_at": time.Now().Unix(), "updated_at": time.Now().Unix()})
		if err != nil {
			return err
		}
		fmt.Println("AddLike,InsertOne succ:", rv)
	} else {
		rv := r.data.likeMG.FindOneAndUpdate(ctx, bson.M{"mid": m.Mid, "oid": m.Oid}, bson.M{"$set": bson.M{"dislike": 1, "updated_at": time.Now().Unix()}})
		if err != nil {
			return err
		}
		fmt.Println("AddLike,FindOneAndUpdate succ:", rv)
	}
	return nil
}

func (r *dislikeRepo) CancelDislike(ctx context.Context, m *model.Dislike) error {
	rv := r.data.likeMG.FindOneAndUpdate(ctx, bson.M{"mid": m.Mid, "oid": m.Oid}, bson.M{"$set": bson.M{"dislike": 0, "updated_at": time.Now().Unix()}})
	if rv.Err() != nil {
		return rv.Err()
	}
	fmt.Println("CancelLike,FindOneAndUpdate succ:", rv)
	return nil
}
func (r *dislikeRepo) Dislike(ctx context.Context, m *model.Dislike) (res *model.Dislike, err error) {
	err = r.data.likeMG.FindOne(ctx, bson.M{"mid": m.Mid, "oid": m.Oid, "bid": m.Bid}).Decode(&res)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if err == mongo.ErrNoDocuments {
			return res, errors.ErrNoDocuments
		}
		// log.Fatal(err)
	}
	return
}
