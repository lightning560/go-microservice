package biz

import (
	"context"
	"miopkg/errors"
	commentv1 "redbook/api/domain/comment/v1"
	"redbook/app/interface/comment/internal/biz/model"
)

type IReplyRepo interface {
	CreateReply(c context.Context, isFloor bool, subjectId int64, bizType string, owner_uid, floorId int64, content string, atUid int64) error

	ListFloorBySubjectSortReply(c context.Context, subjectId int64, offset, limit int32, by, order string) ([]*commentv1.Floor, error)
	ListFloorBySubjectSortTime(c context.Context, subjectId int64, offset, limit int32, by, order string) ([]*commentv1.Floor, error)

	ListReplyByFloorSortLike(c context.Context, subjectId, floorId int64, offset, limit int32, by, order string) ([]*commentv1.Reply, error)
	ListReplyByFloorSortTime(c context.Context, subjectId, floorId int64, offset, limit int32, by, order string) ([]*commentv1.Reply, error)
}

type ReplyBiz struct {
	repo IReplyRepo
}

func NewReplyBiz(repo IReplyRepo) *ReplyBiz {
	return &ReplyBiz{repo: repo}
}

func (b *ReplyBiz) CreateReply(c context.Context, isFloor bool, subjectId int64, bizType string, owner_uid, floorId int64, content string, atUid int64) error {
	return b.repo.CreateReply(c, isFloor, subjectId, bizType, owner_uid, floorId, content, atUid)
}

func (b *ReplyBiz) ListFloorBySubjectSortReply(c context.Context, subjectId int64, offset, limit int32, by, order string) (res *model.Floors, err error) {
	var rv []*commentv1.Floor
	if rv, err = b.repo.ListFloorBySubjectSortReply(c, subjectId, offset, limit, by, order); err != nil {
		return nil, errors.Wrapf(err, "ListFloorBySubjectSortReply subjectId:%d", subjectId)
	}
	res = &model.Floors{}
	res.ListFromDomainProto(rv)
	return
}

func (b *ReplyBiz) ListFloorBySubjectSortTime(c context.Context, subjectId int64, offset, limit int32, by, order string) (res *model.Floors, err error) {
	var rv []*commentv1.Floor
	if rv, err = b.repo.ListFloorBySubjectSortTime(c, subjectId, offset, limit, by, order); err != nil {
		return nil, errors.Wrapf(err, "ListFloorBySubjectSortTime subjectId:%d", subjectId)
	}
	res = &model.Floors{}
	res.ListFromDomainProto(rv)
	return
}

func (b *ReplyBiz) ListReplyByFloorSortLike(c context.Context, subjectId, floorId int64, offset, limit int32, by, order string) (res *model.Replies, err error) {
	var rv []*commentv1.Reply
	if rv, err = b.repo.ListReplyByFloorSortLike(c, subjectId, floorId, offset, limit, by, order); err != nil {
		return nil, errors.Wrapf(err, "ListReplyByFloorSortLike subjectId:%d floorId:%d", subjectId, floorId)
	}
	res = &model.Replies{}
	res.ListFromDomainProto(rv)
	return
}

func (b *ReplyBiz) ListReplyByFloorSortTime(c context.Context, subjectId, floorId int64, offset, limit int32, by, order string) (res *model.Replies, err error) {
	var rv []*commentv1.Reply
	if rv, err = b.repo.ListReplyByFloorSortTime(c, subjectId, floorId, offset, limit, by, order); err != nil {
		return nil, errors.Wrapf(err, "ListReplyByFloorSortTime subjectId:%d floorId:%d", subjectId, floorId)
	}
	res = &model.Replies{}
	res.ListFromDomainProto(rv)
	return
}
