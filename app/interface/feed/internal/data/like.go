package data

import (
	"context"
	"miopkg/log"
	domainfeedv1 "redbook/api/domain/feed/v1"
	"redbook/app/interface/feed/internal/biz"
)

type likeRepo struct {
	data *Data
}

func NewLikeRepo(data *Data) biz.ILikeRepo {
	return &likeRepo{data: data}
}

func (r *likeRepo) AddLike(ctx context.Context, uid, postId int64) error {
	_, err := r.data.feedRpc.AddLike(ctx, &domainfeedv1.AddLikeReq{
		Uid:    uid,
		PostId: postId,
	})
	if err != nil {
		log.Errorf("AddLike error: %v", err)
		return err
	}
	return nil
}

func (r *likeRepo) CancelLike(ctx context.Context, uid, postId int64) error {
	_, err := r.data.feedRpc.CancelLike(ctx, &domainfeedv1.CancelLikeReq{
		Uid:    uid,
		PostId: postId,
	})
	if err != nil {
		log.Errorf("CancelLike error: %v", err)
		return err
	}
	return nil
}
