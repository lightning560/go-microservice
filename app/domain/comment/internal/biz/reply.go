package biz

import (
	"context"
	"fmt"
	"miopkg/errors"
	"redbook/app/domain/comment/model"
	"redbook/app/domain/comment/model/entity"
	"redbook/common/basemodel"
)

type IReplyRepo interface {
	AddReply(ctx context.Context, isFloor bool, subjectId int64, bizType string, floorId, ownerUid int64, content string, atUid int64) (int64, error)
	DeleteReply(context.Context, int64) error

	ListFloorRootBySubjectSortReply(ctx context.Context, subjectId int64, offset, limit int32, by, order string) (entity.Replies, error)
	ListFloorRootBySubjectSortTime(ctx context.Context, subjectId int64, offset, limit int32, by, order string) (entity.Replies, error)

	ListReplyByIds(ctx context.Context, ids []int64) (entity.Replies, error)
	ListReplyByFloorSortLike(ctx context.Context, subjectId, floorId int64, offset, limit int32, by, order string) (entity.Replies, error)
	ListReplyByFloorSortTime(ctx context.Context, subjectId, floorId int64, offset, limit int32, by, order string) (entity.Replies, error)
}
type ReplyBiz struct {
	repo IReplyRepo
}

func NewReplyBiz(repo IReplyRepo) *ReplyBiz {
	return &ReplyBiz{repo: repo}
}

func (r *ReplyBiz) AddReply(ctx context.Context, isFloor bool, subjectId int64, bizType string, floorId, ownerUid int64, content string, atUid int64) (int64, error) {
	return r.repo.AddReply(ctx, isFloor, subjectId, bizType, floorId, ownerUid, content, atUid)
}

func (r *ReplyBiz) DeleteReply(ctx context.Context, id int64) error {
	return r.repo.DeleteReply(ctx, id)
}

func (r *ReplyBiz) ListReplyByIds(ctx context.Context, ids []int64) (res *model.Replies, err error) {
	var rv entity.Replies
	if rv, err = r.repo.ListReplyByIds(ctx, ids); err != nil {
		return nil, errors.Wrapf(err, "ListReplyByIds ids:%v", ids)
	}
	res = &model.Replies{}
	res.ListFromEntity(rv)
	return res, nil
}
func (r *ReplyBiz) ListFloorBySubjectSortReply(ctx context.Context, subjectId int64, offset, limit int32, by, order string) (res *model.Floors, err error) {
	var (
		rv entity.Replies
		// replies        *model.Replies
		repliesOfFloor entity.Replies
	)
	res = &model.Floors{}
	rootReply := &model.Reply{}

	if rv, err = r.repo.ListFloorRootBySubjectSortReply(ctx, subjectId, offset, limit, by, order); err != nil {
		return nil, errors.Wrapf(err, "ListFloorBySubjectSortTime subjectId:%v", subjectId)
	}
	fmt.Printf("biz#ListFloorRootBySubjectSortReply floorRoot %+v\n:", rv)
	for _, v := range rv {
		if repliesOfFloor, err = r.repo.ListReplyByFloorSortLike(ctx, subjectId, v.Id, 0, 2, by, order); err != nil {
			return nil, errors.Wrapf(err, "ListFloorBySubjectSortTime invoke ListReplyByFloorSortLike subjectId:%v , floorId:%v", subjectId, v.Id)
		}
		replies := &model.Replies{}
		rootReply.FromEntity(v)
		replies.ListFromEntity(repliesOfFloor)
		*res = append(*res, &model.Floor{
			RootReply: *rootReply,
			Replies:   *replies,
			Total:     int32(len(*replies)),
			Cursor: basemodel.Cursor{
				Offset: 1,
				Limit:  2,
			},
		})
	}
	return res, nil
}
func (r *ReplyBiz) ListFloorBySubjectSortTime(ctx context.Context, subjectId int64, offset, limit int32, by, order string) (res model.Floors, err error) {
	var rv entity.Replies
	var replies model.Replies
	res = model.Floors{}
	rootReply := &model.Reply{}

	if rv, err = r.repo.ListFloorRootBySubjectSortTime(ctx, subjectId, offset, limit, by, order); err != nil {
		return nil, errors.Wrapf(err, "ListFloorBySubjectSortTime subjectId:%v", subjectId)
	}
	for _, v := range rv {
		if rv, err = r.repo.ListReplyByFloorSortLike(ctx, subjectId, v.Id, 0, 2, by, order); err != nil {
			return nil, errors.Wrapf(err, "ListFloorBySubjectSortTime invoke ListReplyByFloorSortLike subjectId:%v , floorId:%v", subjectId, v.Id)
		}
		rootReply.FromEntity(v)
		replies.ListFromEntity(rv)
		res = append(res, &model.Floor{
			RootReply: *rootReply,
			Replies:   replies,
			Total:     int32(len(replies)),
			Cursor: basemodel.Cursor{
				Offset: 0,
				Limit:  2,
			},
		})
	}
	return res, nil
}

func (r *ReplyBiz) ListReplyByFloorSortLike(ctx context.Context, subjectId, floorId int64, offset, limit int32, by, order string) (res *model.Replies, err error) {
	var rv entity.Replies
	if rv, err = r.repo.ListReplyByFloorSortLike(ctx, subjectId, floorId, offset, limit, by, order); err != nil {
		return nil, errors.Wrapf(err, "ListReplyByFloorSortLike subjectId:%v floorId:%v", subjectId, floorId)
	}
	fmt.Println("biz#ListReplyByFloorSortLike len(rv):", len(rv))
	res = &model.Replies{}
	res.ListFromEntity(rv)
	fmt.Println("biz#ListReplyByFloorSortLike len(res):", len(*res))
	return res, nil
}
func (r *ReplyBiz) ListReplyByFloorSortTime(ctx context.Context, subjectId, floorId int64, offset, limit int32, by, order string) (res model.Replies, err error) {
	var rv entity.Replies
	if rv, err = r.repo.ListReplyByFloorSortTime(ctx, subjectId, floorId, offset, limit, by, order); err != nil {
		return nil, errors.Wrapf(err, "ListReplyByFloorSortTime subjectId:%v floorId:%v", subjectId, floorId)
	}
	res = model.Replies{}
	res.ListFromEntity(rv)
	return res, nil
}
