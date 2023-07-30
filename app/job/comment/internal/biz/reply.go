package biz

import (
	"context"
)

type IReplyRepo interface {
	AddReply(ctx context.Context, isFloor bool, subjectId, floorId, ownerUid, atUid int64, bizType, content string) error
}

type ReplyBiz struct {
	repo IReplyRepo
}

func NewReplyBiz(repo IReplyRepo) *ReplyBiz {
	return &ReplyBiz{
		repo: repo,
	}
}

func (biz *ReplyBiz) AddReply(ctx context.Context, isFloor bool, subjectId, floorId, ownerUid, atUid int64, bizType, content string) error {
	return biz.repo.AddReply(ctx, isFloor, subjectId, floorId, ownerUid, atUid, bizType, content)
}
