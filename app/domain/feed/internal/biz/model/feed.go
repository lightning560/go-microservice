package model

import (
	feedv1 "redbook/api/domain/feed/v1"
	"redbook/app/domain/feed/internal/data/entity"
	"redbook/common/basemodel"
)

type Posts []*Post
type PostCards []*PostCard

type Post struct {
	Id_        string            `json:"_id,omitempty" bson:"_id,omitempty"`
	Id         int64             `json:"id,omitempty" bson:"id,omitempty"`
	Uid        int64             `json:"uid,omitempty" bson:"uid,omitempty"`
	Title      string            `json:"title,omitempty" bson:"title,omitempty"`
	Content    string            `json:"content,omitempty" bson:"content,omitempty"`
	CreatedAt  int64             `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt  int64             `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Cover      *basemodel.Image  `json:"cover,omitempty" bson:"cover,omitempty"`
	Type       string            `json:"type,omitempty" bson:"type,omitempty"`
	Video      *basemodel.Video  `json:"video,omitempty" bson:"video,omitempty"`
	Images     *basemodel.Images `json:"images,omitempty" bson:"images,omitempty"`
	Tags       *basemodel.Tags   `json:"tags,omitempty" bson:"tags,omitempty"`
	Deleted    int8              `json:"deleted,omitempty" bson:"deleted,omitempty"`
	State      int32             `json:"state,omitempty" bson:"state,omitempty"`
	LikeCount  int32             `json:"like_count,omitempty" bson:"like_count,omitempty"`
	ShareCount int32             `json:"share_count,omitempty" bson:"share_count,omitempty"`
	FavorCount int32             `json:"favor_count,omitempty" bson:"favor_count,omitempty"`
	ViewCount  int32             `json:"view_count,omitempty" bson:"view_count,omitempty"`
	CommentId  int64             `json:"comment_id,omitempty" bson:"comment_id,omitempty"`
}

type PostCard struct {
	Id_       string           `json:"_id,omitempty" bson:"_id,omitempty"`
	Id        int64            `json:"id,omitempty" bson:"id,omitempty"`
	Uid       int64            `json:"uid,omitempty" bson:"uid,omitempty"`
	Type      string           `json:"type,omitempty" bson:"type,omitempty"`
	Title     string           `json:"title,omitempty" bson:"title,omitempty"`
	Cover     *basemodel.Image `json:"cover,omitempty" bson:"cover,omitempty"`
	State     int32            `json:"state,omitempty" bson:"state,omitempty"`
	LikeCount int32            `json:"like_count,omitempty" bson:"like_count,omitempty"`
	Tags      *basemodel.Tags  `json:"tags,omitempty" bson:"tags,omitempty"`
}

func (m *Post) ToPostCard() *PostCard {
	return &PostCard{
		Id:        m.Id,
		Uid:       m.Uid,
		Type:      m.Type,
		Title:     m.Title,
		Cover:     m.Cover,
		State:     m.State,
		LikeCount: m.LikeCount,
		Tags:      m.Tags,
	}
}

func (ms *Posts) ListToPostCard() *PostCards {
	var pcs PostCards
	for _, m := range *ms {
		pcs = append(pcs, m.ToPostCard())
	}
	return &pcs
}

func (m *Post) FromEntity(e *entity.Post) {
	m.Id_ = e.Id_.Hex()
	m.Id = e.Id
	m.Uid = e.Uid
	m.Title = e.Title
	m.Content = e.Content
	m.CreatedAt = e.CreatedAt
	m.UpdatedAt = e.UpdatedAt
	m.Cover = e.Cover
	m.Type = e.Type
	m.Video = e.Video
	m.Images = e.Images
	m.Tags = e.Tags
	m.State = e.State
	m.LikeCount = e.LikeCount
	m.ShareCount = e.ShareCount
	m.FavorCount = e.FavorCount
	m.ViewCount = e.ViewCount
	m.CommentId = e.CommentId
}

func (m *Post) ToEntity() *entity.Post {
	e := &entity.Post{}
	e.Id = m.Id
	e.Uid = m.Uid
	e.Title = m.Title
	e.Content = m.Content
	e.CreatedAt = m.CreatedAt
	e.UpdatedAt = m.UpdatedAt
	e.Cover = m.Cover
	e.Type = m.Type
	e.Video = m.Video
	e.Images = m.Images
	e.Tags = m.Tags
	e.State = m.State
	e.LikeCount = m.LikeCount
	e.ShareCount = m.ShareCount
	e.FavorCount = m.FavorCount
	e.ViewCount = m.ViewCount
	e.CommentId = m.CommentId
	return e
}

func (ms *Posts) ListFromEntity(es []*entity.Post) error {
	for _, e := range es {
		var m Post
		m.FromEntity(e)
		*ms = append(*ms, &m)
	}
	return nil
}

func (ms *Posts) ListToEntity() ([]*entity.Post, error) {
	var es []*entity.Post
	for _, m := range *ms {
		e := m.ToEntity()
		es = append(es, e)
	}
	return es, nil
}

func (m *Post) ToProto() *feedv1.Post {
	p := &feedv1.Post{}
	p.Id = m.Id
	p.Uid = m.Uid
	p.Title = m.Title
	p.Content = m.Content
	p.CreatedAt = m.CreatedAt
	p.UpdatedAt = m.UpdatedAt
	p.Cover = m.Cover.ToProto()
	p.Type = m.Type
	if m.Video != nil {
		p.Video = m.Video.ToProto()
	}
	if m.Images != nil {
		p.Images = m.Images.ListToProto()
	}
	p.Tags = m.Tags.ListToProto()
	p.State = m.State
	p.LikeCount = m.LikeCount
	p.ShareCount = m.ShareCount
	p.FavorCount = m.FavorCount
	p.ViewCount = m.ViewCount
	p.CommentId = m.CommentId
	return p
}

func (ms *Posts) ListToProto() []*feedv1.Post {
	var ps []*feedv1.Post
	for _, m := range *ms {
		p := m.ToProto()
		ps = append(ps, p)
	}
	return ps
}

func (m *PostCard) FromEntity(e *entity.Post) error {
	m.Id_ = e.Id_.Hex()
	m.Id = e.Id
	m.Uid = e.Uid
	m.Type = e.Type
	m.Title = e.Title
	m.Cover = e.Cover
	m.State = e.State
	m.LikeCount = e.LikeCount
	m.Tags = e.Tags
	return nil
}

func (m *PostCard) ToProto() (*feedv1.PostCard, error) {
	p := &feedv1.PostCard{}
	p.Id = m.Id
	p.Uid = m.Uid
	p.Type = m.Type
	p.Title = m.Title
	p.Cover = m.Cover.ToProto()
	p.State = m.State
	p.LikeCount = m.LikeCount
	p.Tags = m.Tags.ListToProto()
	return p, nil
}

func (ms *PostCards) ListFromEntity(es []*entity.Post) error {
	for _, e := range es {
		var m PostCard
		if err := m.FromEntity(e); err != nil {
			return err
		}
		*ms = append(*ms, &m)
	}
	return nil
}

func (ms *PostCards) ListToProto() []*feedv1.PostCard {
	var ps []*feedv1.PostCard
	for _, m := range *ms {
		p, _ := m.ToProto()
		ps = append(ps, p)
	}
	return ps
}
