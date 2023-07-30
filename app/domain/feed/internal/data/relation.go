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

type relationRepo struct {
	data *Data
}

func NewRelationRepo(data *Data) biz.IRelationRepo {
	return &relationRepo{data: data}
}

func (r *relationRepo) AddFollow(ctx context.Context, followerUid, followeeUid int64) error {
	exists, err := r.IsFollow(ctx, followerUid, followeeUid)
	if err != nil {
		return errors.Wrapf(err, "AddFollow IsFollow followerUid: %d, followeeUid %d", followerUid, followeeUid)
	}
	if exists {
		return nil
	}
	_, err = r.data.relationMG.InsertOne(ctx, model.Relation{FollowerUid: followerUid, FolloweeUid: followeeUid, CreatedAt: time.Now().Unix()})
	if err != nil {
		return errors.Wrapf(err, "AddFollow InsertOne followerUid: %d, followeeUid %d", followerUid, followeeUid)
	}
	return nil
}

func (r *relationRepo) CancelFollow(ctx context.Context, followerUid, followeeUid int64) error {
	_, err := r.data.relationMG.DeleteOne(ctx, model.Relation{FollowerUid: followerUid, FolloweeUid: followeeUid})
	if err != nil {
		r.data.logger.Errorf("CancelFollow DeleteOne", log.Any("error", err))
		return err
	}
	return nil
}

func (r *relationRepo) IsFollow(ctx context.Context, followerUid, followeeUid int64) (bool, error) {
	if sr := r.data.relationMG.FindOne(ctx, model.Relation{FollowerUid: followerUid, FolloweeUid: followeeUid}); sr.Err() != nil {
		if sr.Err() == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, errors.Wrapf(sr.Err(), "IsFollow FindOne followerUid: %d, followeeUid %d", followerUid, followeeUid)
	}
	return true, nil
}

func (r *relationRepo) ListFollower(ctx context.Context, followeeUid int64) (*[]model.Relation, error) {
	var res []model.Relation
	cur, err := r.data.relationMG.Find(ctx, model.Relation{FolloweeUid: followeeUid})
	if err != nil {
		r.data.logger.Errorf("ListFollower Find", log.Any("error", err))
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var m model.Relation
		err = cur.Decode(&m)
		if err != nil {
			r.data.logger.Errorf("ListFollower Decode", log.Any("error", err))
			return nil, err
		}
		res = append(res, m)
	}
	return &res, nil
}

func (r *relationRepo) ListFollowee(ctx context.Context, followerUid int64) (*[]model.Relation, error) {
	var res []model.Relation
	cur, err := r.data.relationMG.Find(ctx, model.Relation{FollowerUid: followerUid})
	if err != nil {
		r.data.logger.Errorf("ListFollowee Find", log.Any("error", err))
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var m model.Relation
		err = cur.Decode(&m)
		if err != nil {
			r.data.logger.Errorf("ListFollowee Decode", log.Any("error", err))
			return nil, err
		}
		res = append(res, m)
	}
	return &res, nil
}
