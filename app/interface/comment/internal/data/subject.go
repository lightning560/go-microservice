package data

import (
	"context"
	commentv1 "redbook/api/domain/comment/v1"
	"redbook/app/interface/comment/internal/biz"
)

type subjectRepo struct {
	data *Data
}

func NewSubjectRepo(data *Data) biz.ISubjectRepo {
	return &subjectRepo{data: data}
}

func (r *subjectRepo) CreateSubject(c context.Context, belongId int64, bizType string, ownerUid int64) (int64, error) {
	rv, err := r.data.commentRpcClient.CreateSubject(c, &commentv1.CreateSubjectReq{
		BelongId: belongId,
		BizType:  bizType,
		OwnerUid: ownerUid,
	})
	if err != nil {
		return -1, err
	}
	return rv.SubjectId, nil
}

func (r *subjectRepo) GetSubjectById(c context.Context, id int64) (s *commentv1.Subject, err error) {
	rv, err := r.data.commentRpcClient.GetSubjectById(c, &commentv1.GetSubjectByIdReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	return rv.Subject, nil
}
