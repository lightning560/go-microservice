package biz

import "context"

type ILikeRepo interface {
	AddLike(c context.Context, replyId, uid int64) error
	CancelLike(c context.Context, replyId, uid int64) error
	IsLike(c context.Context, replyId, uid int64) (bool, error)
}
type LikeBiz struct {
	repo ILikeRepo
}

func NewLikeBiz(repo ILikeRepo) *LikeBiz {
	return &LikeBiz{repo: repo}
}
func (b *LikeBiz) AddLike(c context.Context, replyId, uid int64) error {
	return b.repo.AddLike(c, replyId, uid)
}
func (b *LikeBiz) CancelLike(c context.Context, replyId, uid int64) error {
	return b.repo.CancelLike(c, replyId, uid)
}
func (b *LikeBiz) IsLike(c context.Context, replyId, uid int64) (bool, error) {
	return b.repo.IsLike(c, replyId, uid)
}
