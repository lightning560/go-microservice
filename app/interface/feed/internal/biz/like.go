package biz

import "context"

type ILikeRepo interface {
	AddLike(ctx context.Context, uid, postId int64) error
	CancelLike(ctx context.Context, uid, postId int64) error
}
type LikeBiz struct {
	repo ILikeRepo
}

func NewLikeBiz(repo ILikeRepo) *LikeBiz {
	return &LikeBiz{repo: repo}
}
func (l *LikeBiz) AddLike(ctx context.Context, uid, postId int64) error {
	return l.repo.AddLike(ctx, uid, postId)
}
func (l *LikeBiz) CancelLike(ctx context.Context, uid, postId int64) error {
	return l.repo.CancelLike(ctx, uid, postId)
}
