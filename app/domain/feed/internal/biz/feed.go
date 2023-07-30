package biz

import (
	"context"
	"fmt"

	"redbook/app/domain/feed/internal/biz/model"

	"miopkg/errors"
)

type IFeedRepo interface {
	CreatePost(ctx context.Context, m *model.Post) (int64, error)
	UpdatePost(ctx context.Context, m *model.Post) (*model.Post, error)
	GetPostById(ctx context.Context, id int64) (*model.Post, error)
	ListPostCard(ctx context.Context, offset, limit int32, by, order string) (model.Posts, error)
	ListPostCardByIds(ctx context.Context, ids []int64) (model.Posts, error)
	ListPostCardByUid(ctx context.Context, uid int64, offset, limit int32, by, order string) (model.Posts, error)
	ListPostCardByUidWithTimestamp(ctx context.Context, uid int64, offset, limit int32, timestamp int64, by, order string) (model.Posts, error)
	ListVideoPost(ctx context.Context, offset, limit int32, by, order string) (model.Posts, error)
}

type FeedBiz struct {
	repo IFeedRepo
}

func NewFeedBiz(repo IFeedRepo) *FeedBiz {
	return &FeedBiz{repo: repo}
}
func (b *FeedBiz) CreatePost(ctx context.Context, m *model.Post) (int64, error) {
	return b.repo.CreatePost(ctx, m)
}
func (b *FeedBiz) UpdatePost(ctx context.Context, m *model.Post) (*model.Post, error) {
	return b.repo.UpdatePost(ctx, m)
}
func (b *FeedBiz) GetPostById(ctx context.Context, id int64) (*model.Post, error) {
	return b.repo.GetPostById(ctx, id)
}

func (b *FeedBiz) ListPostCard(ctx context.Context, offset, limit int32, sortBy, sortOrder string) (model.PostCards, error) {
	rv, err := b.repo.ListPostCard(ctx, offset, limit, sortBy, sortOrder)
	if err != nil {
		return nil, errors.Wrap(err, "biz.ListPostCard")
	}
	return *rv.ListToPostCard(), nil
}

func (b *FeedBiz) ListPostCardByIds(ctx context.Context, ids []int64) (model.PostCards, error) {
	rv, err := b.repo.ListPostCardByIds(ctx, ids)
	if err != nil {
		return nil, errors.Wrap(err, "biz.ListPostCardByIds")
	}
	fmt.Println("biz---len(rv):", len(rv), "+++len(ids):", len(*rv.ListToPostCard()))
	return *rv.ListToPostCard(), nil
}

func (b *FeedBiz) ListPostCardByUid(ctx context.Context, uid int64, offset, limit int32, sortBy, sortOrder string) (model.PostCards, error) {
	rv, err := b.repo.ListPostCardByUid(ctx, uid, offset, limit, sortBy, sortOrder)
	if err != nil {
		return nil, errors.Wrap(err, "biz.ListPostCardByUid")
	}
	return *rv.ListToPostCard(), nil
}
func (b *FeedBiz) ListPostCardByUidWithTimestamp(ctx context.Context, uid int64, offset, limit int32, timestamp int64, sortBy, sortOrder string) (model.PostCards, error) {
	rv, err := b.repo.ListPostCardByUidWithTimestamp(ctx, uid, offset, limit, timestamp, sortBy, sortOrder)
	if err != nil {
		return nil, errors.Wrap(err, "biz.ListTimelinePostCardWithTimestamp")
	}
	return *rv.ListToPostCard(), nil
}
func (b *FeedBiz) ListVideoPost(ctx context.Context, offset, limit int32, sortBy, sortOrder string) (model.Posts, error) {
	return b.repo.ListVideoPost(ctx, offset, limit, sortBy, sortOrder)
}
