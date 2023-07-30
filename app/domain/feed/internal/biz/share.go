package biz

import (
	"context"
	"redbook/app/domain/feed/internal/biz/model"
)

type IShareRepo interface {
	AddShare(ctx context.Context, uid, postId int64) error
	CancelShare(ctx context.Context, uid, postId int64) error
	IsShare(ctx context.Context, uid, postId int64) (bool, error)
	ListShareByPost(ctx context.Context, postId int64) (*[]model.Share, error)
	ListShareByUid(ctx context.Context, uid int64) (*[]model.Share, error)
	GetShareCountByPost(ctx context.Context, postId int64) (int32, error)
}

type ShareBiz struct {
	repo IShareRepo
}

func NewShareBiz(repo IShareRepo) *ShareBiz {
	return &ShareBiz{repo: repo}
}

func (b *ShareBiz) AddShare(ctx context.Context, uid, postId int64) error {
	err := b.repo.AddShare(ctx, uid, postId)
	if err != nil {
		return err
	}
	return nil
}

func (b *ShareBiz) CancelShare(ctx context.Context, uid, postId int64) error {
	err := b.repo.CancelShare(ctx, uid, postId)
	if err != nil {
		return err
	}
	return nil
}
func (b *ShareBiz) IsShare(ctx context.Context, uid, postId int64) (bool, error) {
	rv, err := b.repo.IsShare(ctx, uid, postId)
	if err != nil {
		return false, err
	}
	return rv, nil
}
func (b *ShareBiz) ListShareByUid(ctx context.Context, uid int64) (*[]model.Share, error) {
	rv, err := b.repo.ListShareByUid(ctx, uid)
	if err != nil {
		return nil, err
	}
	return rv, nil
}
func (b *ShareBiz) ListShareByPost(ctx context.Context, postId int64) (*[]model.Share, error) {
	rv, err := b.repo.ListShareByPost(ctx, postId)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

func (b *ShareBiz) GetShareCountByPost(ctx context.Context, uid int64) (int32, error) {
	rv, err := b.repo.GetShareCountByPost(ctx, uid)
	if err != nil {
		return 0, err
	}
	return rv, nil
}
