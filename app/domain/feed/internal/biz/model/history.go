package model

import "redbook/app/domain/feed/internal/data/entity"

type HistoryItem struct {
	Id        int64    `json:"id,omitempty" bson:"id,omitempty"`
	PostID    int64    `json:"post_id,omitempty" bson:"post_id,omitempty"`
	Uid       int64    `json:"uid,omitempty" bson:"uid,omitempty"`
	CreatedAt int64    `json:"created_at,omitempty" bson:"created_at,omitempty"`
	PostCard  *PostCard `json:"post_card,omitempty" bson:"post_card,omitempty"`
}

type HistoryItems []*HistoryItem

func (m *HistoryItem) ToEntity() *entity.HistoryItem {
	return &entity.HistoryItem{
		PostID: m.PostID,
		Uid:    m.Uid,
	}
}

func (m *HistoryItem) FromEntity(entity *entity.HistoryItem) {
	m.Id = entity.Id
	m.PostID = entity.PostID
	m.Uid = entity.Uid
	m.CreatedAt = entity.CreatedAt
}

func (ms HistoryItems) ListToEntity() (entities []*entity.HistoryItem) {
	for _, m := range ms {
		entities = append(entities, m.ToEntity())
	}
	return
}
func (ms HistoryItems) ListFromEntity(entities []*entity.HistoryItem) {
	for _, entity := range entities {
		m := &HistoryItem{}
		m.FromEntity(entity)
		ms = append(ms, m)
	}
}
