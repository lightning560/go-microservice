package data

import (
	"context"
	"miopkg/errors"
	"miopkg/log"
	"redbook/app/domain/feed/internal/biz"
	"redbook/app/domain/feed/internal/biz/model"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type likeRepo struct {
	data *Data
}

func NewLikeRepo(data *Data) biz.ILikeRepo {
	return &likeRepo{data: data}
}

func (r *likeRepo) AddLike(ctx context.Context, uid, postId int64) error {
	exists, err := r.IsLike(ctx, uid, postId)
	if err != nil {
		return errors.Wrapf(err, "AddLike IsLike uid: %d,postId:%d", uid, postId)
	}
	if exists {
		return nil
	}
	_, err = r.data.likeMG.InsertOne(ctx, model.Like{Uid: uid, PostId: postId, CreatedAt: time.Now().Unix()})
	if err != nil {
		r.data.logger.Errorf("AddLike InsertOne", log.Any("error", err))
		return err
	}
	return nil
}
func (r *likeRepo) IsLike(ctx context.Context, uid, postId int64) (bool, error) {
	rv := r.data.likeMG.FindOne(ctx, model.Like{Uid: uid, PostId: postId})
	if rv.Err() != nil {
		if rv.Err() == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, errors.Wrapf(rv.Err(), "IsLike FindOne uid: %d,postId:%d", uid, postId)
	}
	return true, nil
}

func (r *likeRepo) CancelLike(ctx context.Context, uid, postId int64) error {
	_, err := r.data.likeMG.DeleteOne(ctx, model.Like{Uid: uid, PostId: postId})
	if err != nil {
		r.data.logger.Errorf("CancelLike DeleteOne", log.Any("error", err))
		return err
	}
	return nil
}

func (r *likeRepo) ListLikeByPost(ctx context.Context, postId int64) (*[]model.Like, error) {
	var res []model.Like
	cursor, err := r.data.likeMG.Find(ctx, model.Like{PostId: postId})
	if err != nil {
		r.data.logger.Errorf("ListLikeByPost Find", log.Any("error", err))
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var like model.Like
		err = cursor.Decode(&like)
		if err != nil {
			r.data.logger.Errorf("ListLikeByPost Decode", log.Any("error", err))
			return nil, err
		}
		res = append(res, like)
	}
	return &res, nil
}

func (r *likeRepo) ListLikeByUid(ctx context.Context, uid int64) (*[]model.Like, error) {
	var res []model.Like
	cursor, err := r.data.likeMG.Find(ctx, model.Like{Uid: uid})
	if err != nil {
		r.data.logger.Errorf("ListLikeByUser Find", log.Any("error", err))
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var like model.Like
		err = cursor.Decode(&like)
		if err != nil {
			r.data.logger.Errorf("ListLikeByUser Decode", log.Any("error", err))
			return nil, err
		}
		res = append(res, like)
	}
	return &res, nil
}

func (r *likeRepo) GetLikeCountByPost(ctx context.Context, postId int64) (int32, error) {
	count, err := r.data.likeMG.CountDocuments(ctx, model.Like{PostId: postId})
	if err != nil {
		r.data.logger.Errorf("GetLikeCount CountDocuments", log.Any("error", err))
		return 0, err
	}
	return int32(count), nil
}
