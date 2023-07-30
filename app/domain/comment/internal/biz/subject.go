package biz

import (
	"context"
	"miopkg/errors"
	"redbook/app/domain/comment/model"
	"redbook/app/domain/comment/model/entity"
)

type ISubjectRepo interface {
	CreateSubject(ctx context.Context, belongId int64, bizType string, ownerUid int64) (int64, error)
	GetSubjectById(ctx context.Context, subjectId int64) (*entity.Subject, error)
	GetSubjectByBelong(ctx context.Context, belongId int64, bizType string) (*entity.Subject, error)
}

type SubjectBiz struct {
	repo ISubjectRepo
}

func NewSubjectBiz(repo ISubjectRepo) *SubjectBiz {
	return &SubjectBiz{repo: repo}
}

func (s *SubjectBiz) CreateSubject(ctx context.Context, belongId int64, bizType string, ownerUid int64) (int64, error) {
	return s.repo.CreateSubject(ctx, belongId, bizType, ownerUid)
}
func (s *SubjectBiz) GetSubjectById(ctx context.Context, subjectId int64) (res *model.Subject, err error) {
	var rv *entity.Subject
	if rv, err = s.repo.GetSubjectById(ctx, subjectId); err != nil {
		return nil, errors.Wrapf(err, "GetSubject subjectId:%v", subjectId)
	}
	res = &model.Subject{}
	res.FromEntity(rv)
	return
}

func (s *SubjectBiz) GetSubjectByBelong(ctx context.Context, belongId int64, bizType string) (res *model.Subject, err error) {
	var rv *entity.Subject
	if rv, err = s.repo.GetSubjectByBelong(ctx, belongId, bizType); err != nil {
		return nil, errors.Wrapf(err, "GetSubjectByBlongId belongId:%v, bizType:%v", belongId, bizType)
	}
	res = &model.Subject{}
	res.FromEntity(rv)
	return
}
