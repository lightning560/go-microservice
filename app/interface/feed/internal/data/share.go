package data

import (
	"context"
	"miopkg/log"
	domainfeedv1 "redbook/api/domain/feed/v1"
	"redbook/app/interface/feed/internal/biz"
)

type shareRepo struct {
	data *Data
}

func NewShareRepo(data *Data) biz.IShareRepo {
	return &shareRepo{data: data}
}

func (r *shareRepo) AddShare(ctx context.Context, uid, postId int64) error {
	_, err := r.data.feedRpc.AddShare(ctx, &domainfeedv1.AddShareReq{
		Uid:    uid,
		PostId: postId,
	})
	if err != nil {
		log.Errorf("AddShare error: %v", err)
		return err
	}
	return nil
}

func (r *shareRepo) CancelShare(ctx context.Context, uid, postId int64) error {
	_, err := r.data.feedRpc.CancelShare(ctx, &domainfeedv1.CancelShareReq{
		Uid:    uid,
		PostId: postId,
	})
	if err != nil {
		log.Errorf("CancelShare error: %v", err)
		return err
	}
	return nil
}
