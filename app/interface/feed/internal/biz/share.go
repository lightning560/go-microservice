package biz

import "context"

type IShareRepo interface {
	AddShare(ctx context.Context, uid, postId int64) error
	CancelShare(ctx context.Context, uid, postId int64) error
}

type ShareBiz struct {
	repo IShareRepo
}

func NewShareBiz(repo IShareRepo) *ShareBiz {
	return &ShareBiz{repo: repo}
}
func (s *ShareBiz) AddShare(ctx context.Context, uid, postId int64) error {
	return s.repo.AddShare(ctx, uid, postId)
}

func (s *ShareBiz) CancelShare(ctx context.Context, uid, postId int64) error {
	return s.repo.CancelShare(ctx, uid, postId)
}
