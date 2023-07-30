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

type shareRepo struct {
	data *Data
}

func NewShareRepo(data *Data) biz.IShareRepo {
	return &shareRepo{data: data}
}

func (r *shareRepo) AddShare(ctx context.Context, uid, postId int64) error {
	exists, err := r.IsShare(ctx, uid, postId)
	if err != nil {
		return errors.Wrapf(err, "AddShare IsShare uid: %d,postId:%d", uid, postId)
	}
	if exists {
		return nil
	}
	_, err = r.data.shareMG.InsertOne(ctx, model.Share{Uid: uid, PostId: postId, CreatedAt: time.Now().Unix()})
	if err != nil {
		r.data.logger.Errorf("AddShare InsertOne", log.Any("error", err))
		return err
	}
	return nil
}
func (r *shareRepo) CancelShare(ctx context.Context, uid, postId int64) error {
	_, err := r.data.shareMG.DeleteOne(ctx, model.Share{Uid: uid, PostId: postId})
	if err != nil {
		r.data.logger.Errorf("CancelShare DeleteOne", log.Any("error", err))
		return err
	}
	return nil
}
func (r *shareRepo) IsShare(ctx context.Context, uid, postId int64) (bool, error) {
	sr := r.data.shareMG.FindOne(ctx, model.Share{Uid: uid, PostId: postId})
	if sr.Err() != nil {
		if sr.Err() == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, errors.Wrapf(sr.Err(), "IsShare FindOne uid: %d,postId:%d", uid, postId)
	}
	return true, nil
}

func (r *shareRepo) ListShareByPost(ctx context.Context, postId int64) (*[]model.Share, error) {
	var res []model.Share
	cursor, err := r.data.shareMG.Find(ctx, model.Share{PostId: postId})
	if err != nil {
		return nil, errors.Wrapf(err, "ListShareByPost Find postId: %d", postId)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var share model.Share
		err = cursor.Decode(&share)
		if err != nil {
			return nil, errors.Wrapf(err, "ListShareByPost Decode")
		}
		res = append(res, share)
	}
	return &res, nil
}

func (r *shareRepo) ListShareByUid(ctx context.Context, uid int64) (*[]model.Share, error) {
	var res []model.Share
	cursor, err := r.data.shareMG.Find(ctx, model.Share{Uid: uid})
	if err != nil {
		r.data.logger.Errorf("ListShareByUser Find", log.Any("error", err))
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var share model.Share
		err = cursor.Decode(&share)
		if err != nil {
			r.data.logger.Errorf("ListShareByUser Decode", log.Any("error", err))
			return nil, err
		}
		res = append(res, share)
	}
	return &res, nil
}

func (r *shareRepo) GetShareCountByPost(ctx context.Context, postId int64) (int32, error) {
	count, err := r.data.shareMG.CountDocuments(ctx, model.Share{PostId: postId})
	if err != nil {
		r.data.logger.Errorf("CountShareByPost CountDocuments", log.Any("error", err))
		return 0, err
	}
	return int32(count), nil
}
