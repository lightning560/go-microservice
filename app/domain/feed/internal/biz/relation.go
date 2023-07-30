package biz

import (
	"context"
	"redbook/app/domain/feed/internal/biz/model"
)

type IRelationRepo interface {
	AddFollow(ctx context.Context, followerUid, followeeUid int64) error
	CancelFollow(ctx context.Context, followerUid, followeeUid int64) error
	IsFollow(ctx context.Context, followerUid, followeeUid int64) (bool, error)
	ListFollower(ctx context.Context, followeeUid int64) (*[]model.Relation, error)
	ListFollowee(ctx context.Context, followerUid int64) (*[]model.Relation, error)
}
type RelationBiz struct {
	repo IRelationRepo
}

func NewRelationBiz(repo IRelationRepo) *RelationBiz {
	return &RelationBiz{repo: repo}
}

func (b *RelationBiz) AddFollow(ctx context.Context, followerUid, followeeUid int64) error {
	return b.repo.AddFollow(ctx, followerUid, followeeUid)
}

func (b *RelationBiz) CancelFollow(ctx context.Context, followerUid, followeeUid int64) error {
	return b.repo.CancelFollow(ctx, followerUid, followeeUid)
}
func (b *RelationBiz) IsFollow(ctx context.Context, followerUid, followeeUid int64) (bool, error) {
	return b.repo.IsFollow(ctx, followerUid, followeeUid)
}

func (b *RelationBiz) ListFollower(ctx context.Context, followeeUid int64) (*[]model.Relation, error) {
	return b.repo.ListFollower(ctx, followeeUid)
}
func (b *RelationBiz) ListFollowee(ctx context.Context, followerUid int64) (*[]model.Relation, error) {
	return b.repo.ListFollowee(ctx, followerUid)
}
