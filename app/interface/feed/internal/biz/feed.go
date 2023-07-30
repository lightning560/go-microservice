package biz

import (
	"context"
	"redbook/app/interface/feed/internal/biz/model"
)

type IFeedRepo interface {
	CreatePost(ctx context.Context, p *model.Post) (int64, error)
	GetPost(ctx context.Context, id int64) (*model.Post, error)
	ListPostCard(ctx context.Context, offset, limit int32, by, order string) (*model.PostCards, error)
	ListPostCardByIds(ctx context.Context, ids []int64) (model.PostCards, error)
	UpdatePost(ctx context.Context, p *model.Post) error
	DeletePost(ctx context.Context, id int64) error
	ListVideoPost(ctx context.Context, offset, limit int32, by, order string) (*model.Posts, error)
}

type FeedBiz struct {
	repo IFeedRepo
}

func NewFeedBiz(repo IFeedRepo) *FeedBiz {
	return &FeedBiz{repo: repo}
}

func (f *FeedBiz) CreatePost(ctx context.Context, m *model.Post) (int64, error) {
	return f.repo.CreatePost(ctx, m)
}

func (f *FeedBiz) GetPost(ctx context.Context, id int64) (*model.Post, error) {
	return f.repo.GetPost(ctx, id)
}

func (f *FeedBiz) ListPostCard(ctx context.Context, offset, limit int32, by, order string) (*model.PostCards, error) {
	return f.repo.ListPostCard(ctx, offset, limit, by, order)
}

func (f *FeedBiz) ListPostCardByIds(ctx context.Context, ids []int64) (model.PostCards, error) {
	return f.repo.ListPostCardByIds(ctx, ids)
}

func (f *FeedBiz) UpdatePost(ctx context.Context, m *model.Post) error {
	return f.repo.UpdatePost(ctx, m)
}
func (f *FeedBiz) DeletePost(ctx context.Context, id int64) error {
	return f.repo.DeletePost(ctx, id)
}

func (f *FeedBiz) ListVideoPost(ctx context.Context, offset, limit int32, by, order string) (*model.Posts, error) {
	return f.repo.ListVideoPost(ctx, offset, limit, by, order)
}
