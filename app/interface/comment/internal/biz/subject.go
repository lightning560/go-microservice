package biz

import (
	"context"
	"miopkg/errors"
	commentv1 "redbook/api/domain/comment/v1"
	"redbook/app/interface/comment/internal/biz/model"
)

type ISubjectRepo interface {
	CreateSubject(c context.Context, belongId int64, bizType string, ownerUid int64) (int64, error)
	GetSubjectById(c context.Context, id int64) (s *commentv1.Subject, err error)
}
type SubjectBiz struct {
	repo ISubjectRepo
}

func NewSubjectBiz(repo ISubjectRepo) *SubjectBiz {
	return &SubjectBiz{repo: repo}
}

func (b *SubjectBiz) CreateSubject(c context.Context, belongId int64, bizType string, ownerUid int64) (int64, error) {
	return b.repo.CreateSubject(c, belongId, bizType, ownerUid)
}
func (b *SubjectBiz) GetSubjectById(c context.Context, id int64) (res *model.Subject, err error) {
	var rv *commentv1.Subject
	if rv, err = b.repo.GetSubjectById(c, id); err != nil {
		return nil, errors.Wrapf(err, "biz.GetSubjectById id %d", id)
	}
	res = &model.Subject{}
	res.FromDomainProto(rv)
	return
}
