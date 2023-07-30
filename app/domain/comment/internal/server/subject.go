package server

import (
	"context"
	commentv1 "redbook/api/domain/comment/v1"
	"redbook/app/domain/comment/model"
)

func (s *CommentServer) GetSubjectById(c context.Context, req *commentv1.GetSubjectByIdReq) (res *commentv1.GetSubjectByIdResponse, err error) {
	var rv *model.Subject
	if rv, err = s.subjectbiz.GetSubjectById(c, req.Id); err != nil {
		return nil, err
	}
	res = &commentv1.GetSubjectByIdResponse{}
	res.Subject = rv.ToProto()
	return
}

func (s *CommentServer) GetSubjectByBelong(c context.Context, req *commentv1.GetSubjectByBelongReq) (res *commentv1.GetSubjectByBelongResponse, err error) {
	var rv *model.Subject
	if rv, err = s.subjectbiz.GetSubjectByBelong(c, req.BelongId, req.BizType); err != nil {
		return nil, err
	}
	res = &commentv1.GetSubjectByBelongResponse{}
	res.Subject = rv.ToProto()
	return
}

func (s *CommentServer) CreateSubject(c context.Context, req *commentv1.CreateSubjectReq) (res *commentv1.CreateSubjectResponse, err error) {
	var rv int64
	if rv, err = s.subjectbiz.CreateSubject(c, req.BelongId, req.BizType, req.OwnerUid); err != nil {
		return nil, err
	}
	res = &commentv1.CreateSubjectResponse{}
	res.SubjectId = rv
	return
}
