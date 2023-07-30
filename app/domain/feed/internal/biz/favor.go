package biz

import (
	"context"
	"redbook/app/domain/feed/internal/biz/model"
)

type IFavorRepo interface {
	AddFavor(ctx context.Context, uid, postId int64) error
	CancelFavor(ctx context.Context, uid, postId int64) error
	IsFavor(ctx context.Context, uid, postId int64) (bool, error)
	ListFavorByPost(ctx context.Context, postId int64) (*[]model.Favor, error)
	ListFavorByUid(ctx context.Context, uid int64) (*[]model.Favor, error)
	GetFavorCountByPost(ctx context.Context, postId int64) (int32, error)
}
type FavorBiz struct {
	repo IFavorRepo
}

func NewFavorBiz(repo IFavorRepo) *FavorBiz {
	return &FavorBiz{repo: repo}
}

func (b *FavorBiz) AddFavor(ctx context.Context, uid, postId int64) error {
	return b.repo.AddFavor(ctx, uid, postId)
}

func (b *FavorBiz) CancelFavor(ctx context.Context, uid, postId int64) error {
	return b.repo.CancelFavor(ctx, uid, postId)
}
func (b *FavorBiz) IsFavor(ctx context.Context, uid, postId int64) (bool, error) {
	return b.repo.IsFavor(ctx, uid, postId)
}

func (b *FavorBiz) ListFavorByPost(ctx context.Context, postId int64) (*[]model.Favor, error) {
	return b.repo.ListFavorByPost(ctx, postId)
}

func (b *FavorBiz) ListFavorByUid(ctx context.Context, uid int64) (*[]model.Favor, error) {
	return b.repo.ListFavorByUid(ctx, uid)
}
func (b *FavorBiz) GetFavorCountByPost(ctx context.Context, postId int64) (int32, error) {
	return b.repo.GetFavorCountByPost(ctx, postId)
}
