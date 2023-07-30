package server

import (
	"context"
	"miopkg/log"
	commentv1 "redbook/api/domain/comment/v1"
	basev1 "redbook/api/redbookpb/v1"
	"redbook/app/domain/comment/model"
)

func (s *CommentServer) AddReply(c context.Context, req *commentv1.CreateReplyReq) (res *commentv1.CreateReplyResponse, err error) {
	var rv int64
	if rv, err = s.replybiz.AddReply(c, req.IsFloor, req.SubjectId, req.BizType, req.FloorId, req.OwnerUid, req.Content, req.AtUid); err != nil {
		return nil, err
	}
	res = &commentv1.CreateReplyResponse{}
	res.ReplyId = rv
	return
}

func (s *CommentServer) ListReplyByIds(ctx context.Context, req *commentv1.ListReplyByIdsReq) (res *commentv1.ListReplyByIdsResponse, err error) {
	var rv *model.Replies
	if rv, err = s.replybiz.ListReplyByIds(ctx, req.Ids); err != nil {
		return nil, err
	}
	res = &commentv1.ListReplyByIdsResponse{}
	res.Total = int32(len(*rv))
	res.Replies = rv.ListToProto()
	return res, nil
}

func (s *CommentServer) ListFloorBySubjectSortReply(ctx context.Context, req *commentv1.ListFloorBySubjectSortReplyReq) (res *commentv1.ListFloorBySubjectSortReplyResponse, err error) {
	var rv *model.Floors
	if rv, err = s.replybiz.ListFloorBySubjectSortReply(ctx, req.SubjectId, req.Cursor.Offset, req.Cursor.Limit, req.Sort.By, req.Sort.Order); err != nil {
		return nil, err
	}
	res = &commentv1.ListFloorBySubjectSortReplyResponse{}
	res.Total = int32(len(*rv))
	res.Floors = rv.ListToProto()
	res.Cursor = req.Cursor
	return res, nil
}

func (s *CommentServer) ListFloorBySubjectSortTime(ctx context.Context, req *commentv1.ListFloorBySubjectSortTimeReq) (res *commentv1.ListFloorBySubjectSortTimeResponse, err error) {
	var rv model.Floors
	if rv, err = s.replybiz.ListFloorBySubjectSortTime(ctx, req.SubjectId, req.Cursor.Offset, req.Cursor.Limit, req.Sort.By, req.Sort.Order); err != nil {
		return nil, err
	}
	res = &commentv1.ListFloorBySubjectSortTimeResponse{}
	res.Total = int32(len(rv))
	res.Floors = rv.ListToProto()
	res.Cursor = req.Cursor
	return res, nil
}
func (s *CommentServer) ListReplyByFloorSortLike(ctx context.Context, req *commentv1.ListReplyByFloorSortLikeReq) (res *commentv1.ListReplyByFloorSortLikeResponse, err error) {
	var rv *model.Replies
	log.Infof("ListReplyByFloorSortLike req:%+v\n:", req)
	log.Infof("ListReplyByFloorSortLike req.Cursor.Offset:%d\n", req.Cursor.Offset)
	if rv, err = s.replybiz.ListReplyByFloorSortLike(ctx, req.SubjectId, req.FloorId, req.Cursor.Offset, req.Cursor.Limit, req.Sort.By, req.Sort.Order); err != nil {
		return nil, err
	}
	res = &commentv1.ListReplyByFloorSortLikeResponse{}
	res.Total = int32(len(*rv))
	res.Replies = rv.ListToProto()
	res.Cursor = &basev1.Cursor{
		Offset: req.Cursor.Offset,
		Limit:  req.Cursor.Limit,
	}
	return res, nil
}
func (s *CommentServer) ListReplyByFloorSortTime(ctx context.Context, req *commentv1.ListReplyByFloorSortTimeReq) (res *commentv1.ListReplyByFloorSortTimeResponse, err error) {
	var rv model.Replies
	if rv, err = s.replybiz.ListReplyByFloorSortTime(ctx, req.SubjectId, req.FloorId, req.Cursor.Offset, req.Cursor.Limit, req.Sort.By, req.Sort.Order); err != nil {
		return nil, err
	}
	res = &commentv1.ListReplyByFloorSortTimeResponse{}
	res.Total = int32(len(rv))
	res.Replies = rv.ListToProto()
	res.Cursor = req.Cursor
	return res, nil
}
