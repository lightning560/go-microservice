package data

import (
	"context"
	"miopkg/log"
	domainfeedv1 "redbook/api/domain/feed/v1"
	"redbook/app/interface/feed/internal/biz"
)

type favorRepo struct {
	data *Data
}

func NewFavorRepo(data *Data) biz.IFavorRepo {
	return &favorRepo{data: data}
}
func (r *favorRepo) AddFavor(ctx context.Context, uid, postId int64) error {
	_, err := r.data.feedRpc.AddFavor(ctx, &domainfeedv1.AddFavorReq{
		Uid:    uid,
		PostId: postId,
	})
	if err != nil {
		log.Errorf("AddFavor error: %v", err)
		return err
	}
	return nil
}
func (r *favorRepo) CancelFavor(ctx context.Context, uid, postId int64) error {
	_, err := r.data.feedRpc.CancelFavor(ctx, &domainfeedv1.CancelFavorReq{
		Uid:    uid,
		PostId: postId,
	})
	if err != nil {
		log.Errorf("CancelFavor error: %v", err)
		return err
	}
	return nil
}
