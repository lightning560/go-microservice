package biz

import (
	"context"
	"redbook/app/domain/like/internal/data/model"
)

type ILikeRepo interface {
	AddLike(context.Context, *model.Like) error
	Like(context.Context, *model.Like) (*model.Like, error)
}

type LikeBiz struct {
	repo ILikeRepo
}

func NewLikeBiz(repo ILikeRepo) *LikeBiz {
	return &LikeBiz{repo: repo}
}
func (b *LikeBiz) AddLike(ctx context.Context, m *model.Like) error {
	return b.repo.AddLike(ctx, m)
}

func (b *LikeBiz) Like(ctx context.Context, m *model.Like) (*model.Like, error) {
	return b.repo.Like(ctx, m)
}
