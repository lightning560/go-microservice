package data

import (
	"context"
	"fmt"
	"miopkg/errors"
	commentv1 "redbook/api/domain/comment/v1"
	basev1 "redbook/api/redbookpb/v1"
	"redbook/app/interface/comment/internal/biz"
)

type replyRepo struct {
	data *Data
}

func NewReplyRepo(data *Data) biz.IReplyRepo {
	return &replyRepo{data: data}
}

func (r *replyRepo) ListFloorBySubjectSortReply(c context.Context, subjectId int64, offset, limit int32, by, order string) (res []*commentv1.Floor, err error) {
	rv, err := r.data.commentRpcClient.ListFloorBySubjectSortReply(c, &commentv1.ListFloorBySubjectSortReplyReq{
		SubjectId: subjectId,
		Cursor: &basev1.Cursor{
			Offset: offset,
			Limit:  limit,
		},
		Sort: &basev1.Sort{
			By:    by,
			Order: order,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "ListFloorBySubjectSortReply subjectId:%d", subjectId)
	}
	return rv.Floors, nil
}

func (r *replyRepo) ListFloorBySubjectSortTime(c context.Context, subjectId int64, offset, limit int32, by, order string) (res []*commentv1.Floor, err error) {
	rv, err := r.data.commentRpcClient.ListFloorBySubjectSortTime(c, &commentv1.ListFloorBySubjectSortTimeReq{
		SubjectId: subjectId,
		Cursor: &basev1.Cursor{
			Offset: offset,
			Limit:  limit,
		},
		Sort: &basev1.Sort{
			By:    by,
			Order: order,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "ListFloorBySubjectSortTime subjectId:%d", subjectId)
	}
	return rv.Floors, nil
}

func (r *replyRepo) ListReplyByFloorSortLike(c context.Context, subjectId, floorId int64, offset, limit int32, by, order string) ([]*commentv1.Reply, error) {
	rv, err := r.data.commentRpcClient.ListReplyByFloorSortLike(c, &commentv1.ListReplyByFloorSortLikeReq{
		SubjectId: subjectId,
		FloorId:   floorId,
		Cursor: &basev1.Cursor{
			Offset: offset,
			Limit:  limit,
		},
		Sort: &basev1.Sort{
			By:    by,
			Order: order,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "ListReplyByFloorSortLike subjectId:%d , floorId :%d", subjectId, floorId)
	}
	fmt.Println("data ListReplyByFloorSortLike rv.Replies", len(rv.Replies))
	return rv.Replies, nil
}

func (r *replyRepo) ListReplyByFloorSortTime(c context.Context, subjectId, floorId int64, offset, limit int32, by, order string) ([]*commentv1.Reply, error) {
	rv, err := r.data.commentRpcClient.ListReplyByFloorSortTime(c, &commentv1.ListReplyByFloorSortTimeReq{
		SubjectId: subjectId,
		FloorId:   floorId,
		Cursor: &basev1.Cursor{
			Offset: offset,
			Limit:  limit,
		},
		Sort: &basev1.Sort{
			By:    by,
			Order: order,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "ListReplyByFloorSortTime subjectId:%d , floorId : %d", subjectId, floorId)
	}
	return rv.Replies, nil
}
