package biz

import "context"

type IHistoryRepo interface {
	AddHistory(ctx context.Context, uid, itemId int64, bizType string) error
	// GetHistory(ctx context.Context, req *pb.GetHistoryReq) (*pb.GetHistoryResp, error)
	// GetHistoryList(ctx context.Context, req *pb.GetHistoryListReq) (*pb.GetHistoryListResp, error)
	// CreateHistory(ctx context.Context, req *pb.CreateHistoryReq) (*pb.CreateHistoryResp, error)
	// DeleteHistory(ctx context.Context, req *pb.DeleteHistoryReq) (*pb.DeleteHistoryResp, error)
}

type HistoryBiz struct {
	repo IHistoryRepo
}

func NewHistoryBiz(repo IHistoryRepo) *HistoryBiz {
	return &HistoryBiz{
		repo: repo,
	}
}
