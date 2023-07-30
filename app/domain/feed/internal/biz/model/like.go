package model

import (
	"redbook/app/domain/feed/internal/data/entity"
)

type Like struct {
	Id_       string `json:"_id,omitempty" bson:"_id,omitempty"`
	Id        int64  `json:"id,omitempty" bson:"id,omitempty"`
	Uid       int64  `json:"uid,omitempty" bson:"uid,omitempty"`
	PostId    int64  `json:"post_id,omitempty" bson:"post_id,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

func (m *Like) ToEntity() (*entity.Like, error) {
	return &entity.Like{
		Id:        m.Id,
		Uid:       m.Uid,
		PostId:    m.PostId,
		CreatedAt: m.CreatedAt,
	}, nil
}

func (m *Like) FromEntity(e *entity.Like) error {
	m.Id = e.Id
	m.Uid = e.Uid
	m.PostId = e.PostId
	m.CreatedAt = e.CreatedAt
	return nil
}
