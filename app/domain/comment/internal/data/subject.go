package data

import (
	"context"
	"fmt"
	"miopkg/errors"
	"miopkg/util/snowflake"
	"redbook/app/domain/comment/internal/biz"
	"redbook/app/domain/comment/model/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type subjectRepo struct {
	data *Data
}

func NewSubjectRepo(data *Data) biz.ISubjectRepo {
	return &subjectRepo{data: data}
}

func (s *subjectRepo) isSubjectExists(ctx context.Context, belongId int64, bizType string) (bool, error) {
	if rv := s.data.subjectMG.FindOne(ctx, entity.Subject{BelongId: belongId, BizType: bizType}); rv.Err() != nil {
		if rv.Err() == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, errors.Wrapf(rv.Err(), "isSubjectExisted FindOne belongId:%d, bizType:%d", belongId, bizType)
	}
	return true, nil
}

func (s *subjectRepo) GetSubjectByBelong(ctx context.Context, belongId int64, bizType string) (*entity.Subject, error) {
	var rv *entity.Subject
	if err := s.data.subjectMG.FindOne(ctx, entity.Subject{BelongId: belongId, BizType: bizType}).Decode(&rv); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrapf(errors.ErrNoDocuments, "GetSubjectByBlong FindOne belongId:%d, bizType:%d", belongId, bizType)
		}
		return nil, errors.Wrapf(err, "GetSubjectByBlong FindOne belongId:%d, bizType:%d", belongId, bizType)
	}
	return rv, nil
}

// CreateSubject creates a subject, returns the subject id.
// one belongId can only have one subject.
func (s *subjectRepo) CreateSubject(ctx context.Context, belongId int64, bizType string, ownerUid int64) (res int64, err error) {
	rv, err := s.GetSubjectByBelong(ctx, belongId, bizType)
	if errors.Is(err, errors.ErrNoDocuments) {
		id := snowflake.GenInt64()
		et := entity.Subject{
			Id: id,

			OwnerUid: ownerUid,
			BelongId: belongId,
			BizType:  bizType,

			State:      1,
			FloorCount: 0,
			ReplyCount: 0,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),

			Attr: 0,
			Meta: "",
		}
		if _, err := s.data.subjectMG.InsertOne(ctx, et); err != nil {
			return -1, errors.Wrapf(err, "AddSubject InsertOne BelongId:%d,OwnerUid:%d", belongId, ownerUid)
		}
		return id, nil
	}
	if err != nil {
		return -1, errors.Wrapf(err, "AddSubject GetSubjectByBelong BelongId:%d,OwnerUid:%d", belongId, ownerUid)
	}
	fmt.Println("rv.Id", rv.Id)
	return rv.Id, nil
}

func (s *subjectRepo) GetSubjectById(ctx context.Context, id int64) (*entity.Subject, error) {
	subject := new(entity.Subject)
	sr := s.data.subjectMG.FindOne(ctx, bson.M{"id": id})
	if sr.Err() != nil {
		return nil, errors.Wrapf(sr.Err(), "subjectRepo.GetSubject FindOne id:%d", id)
	}
	if err := sr.Decode(subject); err != nil {
		return nil, errors.Wrapf(err, "subjectRepo.GetSubject Decode id:%d", id)
	}
	return subject, nil
}
