package biz

import "context"

type ILikeRepo interface {
	AddLike(ctx context.Context, replyId, uid int64) error
	IsLike(ctx context.Context, replyId, uid int64) (bool, error)
	CancelLike(ctx context.Context, replyId, uid int64) error
}
type LikeBiz struct {
	repo ILikeRepo
}

func NewLikeBiz(repo ILikeRepo) *LikeBiz {
	return &LikeBiz{repo: repo}
}
func (r *LikeBiz) AddLike(ctx context.Context, replyId, uid int64) error {
	return r.repo.AddLike(ctx, replyId, uid)
}
func (r *LikeBiz) IsLike(ctx context.Context, replyId, uid int64) (bool, error) {
	return r.repo.IsLike(ctx, replyId, uid)
}

func (r *LikeBiz) CancelLike(ctx context.Context, replyId, uid int64) error {
	return r.repo.CancelLike(ctx, replyId, uid)
}
