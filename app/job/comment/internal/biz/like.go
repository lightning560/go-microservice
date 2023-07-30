package biz

import "context"

type ILikeRepo interface {
	AddLike(ctx context.Context, replyId, uid int64) error
}

type LikeBiz struct {
	repo ILikeRepo
}

func NewLikeBiz(repo ILikeRepo) *LikeBiz {
	return &LikeBiz{repo: repo}
}

func (b *LikeBiz) AddLike(ctx context.Context, replyId, uid int64) error {
	return b.repo.AddLike(ctx, replyId, uid)
}
