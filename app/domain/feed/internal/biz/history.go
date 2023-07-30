package biz

import (
	"context"
	"redbook/app/domain/feed/internal/biz/model"
)

type IHistoryRepo interface {
	// AddHistoryItem add a history item.
	AddHistoryItem(ctx context.Context, item *model.HistoryItem) error
	// GetHistoryItems get history items.
	ListHistoryItem(ctx context.Context, uid int64, offset, limit int32, sortBy, sortOrder string) (*model.HistoryItems, error)
}
type HistoryBiz struct {
	repo IHistoryRepo
}

func NewHistoryBiz(repo IHistoryRepo) *HistoryBiz {
	return &HistoryBiz{repo: repo}
}
