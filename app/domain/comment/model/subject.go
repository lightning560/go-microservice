package model

import (
	commentv1 "redbook/api/domain/comment/v1"
	"redbook/app/domain/comment/model/entity"
)

type Subject struct {
	Id_ string `json:"_id,omitempty" bson:"_id,omitempty"`
	Id  int64  `json:"id,omitempty" bson:"id,omitempty"`

	BelongId int64  `json:"belong_id,omitempty" bson:"belong_id,omitempty"`
	BizType  string `json:"biz_type,omitempty" bson:"biz_type,omitempty"`

	OwnerUid   int64 `json:"owner_uid,omitempty" bson:"owner_uid,omitempty"`
	FloorCount int32 `json:"floor_count,omitempty" bson:"floor_count,omitempty"`
	ReplyCount int32 `json:"reply_count,omitempty" bson:"reply_count,omitempty"`
	State      int32 `json:"state,omitempty" bson:"state,omitempty"`

	Attr int64  `json:"attr,omitempty" bson:"attr,omitempty"`
	Meta string `json:"meta,omitempty" bson:"meta,omitempty"`

	CreatedAt int64 `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt int64 `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

func (m *Subject) FromEntity(subject *entity.Subject) {
	m.Id_ = subject.Id_.Hex()
	m.Id = subject.Id
	m.BelongId = subject.BelongId
	m.BizType = subject.BizType
	m.OwnerUid = subject.OwnerUid
	m.FloorCount = subject.FloorCount
	m.ReplyCount = subject.ReplyCount
	m.State = subject.State
	m.Attr = subject.Attr
	m.Meta = subject.Meta
	m.CreatedAt = subject.CreatedAt
	m.UpdatedAt = subject.UpdatedAt
}

func (m *Subject) ToProto() *commentv1.Subject {
	return &commentv1.Subject{
		Id:         m.Id,
		OwnerUid:   m.OwnerUid,
		BelongId:   m.BelongId,
		BizType:    m.BizType,
		FloorCount: m.FloorCount,
		ReplyCount: m.ReplyCount,
		State:      m.State,
		Attr:       m.Attr,
		Meta:       m.Meta,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
}
