package biz

import "context"

type IFavorRepo interface {
	AddFavor(ctx context.Context, uid, postId int64) error
	CancelFavor(ctx context.Context, uid, postId int64) error
}

type FavorBiz struct {
	repo IFavorRepo
}

func NewFavorBiz(repo IFavorRepo) *FavorBiz {
	return &FavorBiz{repo: repo}
}

func (f *FavorBiz) AddFavor(ctx context.Context, uid, postId int64) error {
	return f.repo.AddFavor(ctx, uid, postId)
}

func (f *FavorBiz) CancelFavor(ctx context.Context, uid, postId int64) error {
	return f.repo.CancelFavor(ctx, uid, postId)
}
