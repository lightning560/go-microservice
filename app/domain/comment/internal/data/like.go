package data

import (
	"context"
	"fmt"
	"miopkg/errors"
	"redbook/app/domain/comment/internal/biz"
	"redbook/app/domain/comment/model/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type likeRepo struct {
	data *Data
}

func NewLikeRepo(data *Data) biz.ILikeRepo {
	return &likeRepo{data: data}
}

func (r *likeRepo) AddLike(ctx context.Context, replyId, uid int64) error {
	exists, err := r.IsLike(ctx, replyId, uid)
	if err != nil {
		return errors.Wrapf(err, "AddLike IsLike replyId:%d uid:%d", replyId, uid)
	}
	if exists {
		return nil
	}
	if _, err = r.data.likeMG.InsertOne(ctx, entity.Like{ReplyId: replyId, Uid: uid, CreatedAt: time.Now().Unix()}); err != nil {
		return errors.Wrapf(err, "AddLike InsertOne replyId:%d uid:%d", replyId, uid)
	}
	return nil
}

func (r *likeRepo) IsLike(ctx context.Context, replyId, uid int64) (bool, error) {
	if rv := r.data.likeMG.FindOne(ctx, entity.Like{ReplyId: replyId, Uid: uid}); rv.Err() != nil {
		if rv.Err() == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, errors.Wrapf(rv.Err(), "IsLike FindOne replyId:%d uid:%d", replyId, uid)
	}
	return true, nil
}

func (r *likeRepo) CancelLike(ctx context.Context, replyId, uid int64) error {
	fmt.Println("CancelLike replyId:", replyId, "uid:", uid)
	if _, err := r.data.likeMG.DeleteOne(ctx, bson.M{"reply_id": replyId, "uid": uid}); err != nil {
		return errors.Wrapf(err, "CancelLike DeleteOne replyId:%d uid:%d", replyId, uid)
	}
	return nil
}
