package biz

import (
	"context"
	"redbook/app/domain/feed/internal/biz/model"
)

type ILikeRepo interface {
	AddLike(ctx context.Context, uid, postId int64) error
	IsLike(ctx context.Context, uid, postId int64) (bool, error)
	CancelLike(ctx context.Context, uid, postId int64) error
	ListLikeByPost(ctx context.Context, postId int64) (*[]model.Like, error)
	ListLikeByUid(ctx context.Context, uid int64) (*[]model.Like, error)
	GetLikeCountByPost(ctx context.Context, postId int64) (int32, error)
}

type LikeBiz struct {
	repo ILikeRepo
}

func NewLikeBiz(repo ILikeRepo) *LikeBiz {
	return &LikeBiz{repo: repo}
}

func (b *LikeBiz) AddLike(ctx context.Context, uid, postId int64) error {
	return b.repo.AddLike(ctx, uid, postId)
}

func (b *LikeBiz) IsLike(ctx context.Context, uid, postId int64) (bool, error) {
	return b.repo.IsLike(ctx, uid, postId)
}

func (b *LikeBiz) CancelLike(ctx context.Context, uid, postId int64) error {
	return b.repo.CancelLike(ctx, uid, postId)
}

func (b *LikeBiz) ListLikeByPost(ctx context.Context, postId int64) (*[]model.Like, error) {
	return b.repo.ListLikeByPost(ctx, postId)
}

func (b *LikeBiz) ListLikeByUid(ctx context.Context, uid int64) (*[]model.Like, error) {
	return b.repo.ListLikeByUid(ctx, uid)
}
func (b *LikeBiz) GetLikeCountByPost(ctx context.Context, postId int64) (count int32, err error) {
	return b.repo.GetLikeCountByPost(ctx, postId)
}
