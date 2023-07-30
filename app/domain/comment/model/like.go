package model

import "redbook/app/domain/comment/model/entity"

type Like struct {
	Id_       string `json:"_id,omitempty" bson:"_id,omitempty"`
	ReplyId   int64  `json:"reply_id,omitempty" bson:"reply_id,omitempty"`
	Uid       int64  `json:"uid,omitempty" bson:"uid,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
type Dislike struct {
	Id_       string `json:"_id,omitempty" bson:"_id,omitempty"`
	ReplyId   int64  `json:"reply_id,omitempty" bson:"reply_id,omitempty"`
	Uid       int64  `json:"uid,omitempty" bson:"uid,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

func (m *Like) FromEntity(et *entity.Like) {
	m.Id_ = et.Id_.Hex()
	m.ReplyId = et.ReplyId
	m.Uid = et.Uid
	m.CreatedAt = et.CreatedAt
}

func (m *Dislike) FromEntity(et *entity.Dislike) {
	m.Id_ = et.Id_.Hex()
	m.ReplyId = et.ReplyId
	m.Uid = et.Uid
	m.CreatedAt = et.CreatedAt
}
