package data

import (
	"context"
	"miopkg/log"
	"redbook/app/domain/feed/internal/biz"
	"redbook/app/domain/feed/internal/biz/model"
	"time"

	"miopkg/errors"

	"go.mongodb.org/mongo-driver/mongo"
)

type favorRepo struct {
	data *Data
}

func NewFavorRepo(data *Data) biz.IFavorRepo {
	return &favorRepo{data: data}
}

func (r *favorRepo) AddFavor(ctx context.Context, uid, postId int64) error {
	exists, err := r.IsFavor(ctx, uid, postId)
	if err != nil {
		return errors.Wrapf(err, "AddFavor IsFavor uid: %d,postId:%d", uid, postId)
	}
	if exists {
		return nil
	}
	_, err = r.data.favorMG.InsertOne(ctx, model.Favor{Uid: uid, PostId: postId, CreatedAt: time.Now().Unix()})
	if err != nil {
		r.data.logger.Errorf("AddFavor InsertOne", log.Any("error", err))
		return err
	}
	return nil
}
func (r *favorRepo) CancelFavor(ctx context.Context, uid, postId int64) error {
	_, err := r.data.favorMG.DeleteOne(ctx, model.Favor{Uid: uid, PostId: postId})
	if err != nil {
		return errors.Wrapf(err, "CancelFavor DeleteOne uid: %d,postId:%d", uid, postId)
	}
	return nil
}
func (r *favorRepo) IsFavor(ctx context.Context, uid, postId int64) (bool, error) {
	sr := r.data.favorMG.FindOne(ctx, model.Favor{Uid: uid, PostId: postId})
	if sr.Err() != nil {
		if sr.Err() == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, errors.Wrapf(sr.Err(), "IsFavor FindOne uid: %d,postId:%d", uid, postId)
	}
	return true, nil
}
func (r *favorRepo) ListFavorByPost(ctx context.Context, postId int64) (*[]model.Favor, error) {
	var res []model.Favor
	cursor, err := r.data.favorMG.Find(ctx, model.Favor{PostId: postId})
	if err != nil {
		return nil, errors.Wrapf(err, "ListFavorByPost Find postId:%d", postId)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var favor model.Favor
		err = cursor.Decode(&favor)
		if err != nil {
			return nil, errors.Wrapf(err, "ListFavorByPost Decode")
		}
		res = append(res, favor)
	}
	return &res, nil
}
func (r *favorRepo) ListFavorByUid(ctx context.Context, uid int64) (*[]model.Favor, error) {
	var res []model.Favor
	cursor, err := r.data.favorMG.Find(ctx, model.Favor{Uid: uid})
	if err != nil {
		return nil, errors.Wrapf(err, "ListFavorByUser Find uid:%d", uid)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var favor model.Favor
		err = cursor.Decode(&favor)
		if err != nil {
			return nil, errors.Wrapf(err, "ListFavorByUser Decode")
		}
		res = append(res, favor)
	}
	return &res, nil
}

func (r *favorRepo) GetFavorCountByPost(ctx context.Context, postId int64) (int32, error) {
	count, err := r.data.favorMG.CountDocuments(ctx, model.Favor{PostId: postId})
	if err != nil {
		r.data.logger.Errorf("GetFavorCount CountDocuments", log.Any("error", err))
		return 0, err
	}
	return int32(count), nil
}
