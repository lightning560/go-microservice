package model

import (
	domainfeedv1 "redbook/api/domain/feed/v1"
	interfacefeedv1 "redbook/api/interface/feed/v1"
	"redbook/common/basemodel"
	"strconv"
)

type Post struct {
	Id         int64               `json:"id,omitempty" bson:"_id,omitempty"`
	Uid        int64               `json:"uid,omitempty" bson:"id,omitempty"`
	AuthorInfo *basemodel.UserInfo `json:"author_info,omitempty" bson:"author_info,omitempty"`
	Title      string              `json:"title,omitempty" bson:"title,omitempty"`
	Content    string              `json:"content,omitempty" bson:"content,omitempty"`
	CreatedAt  int64               `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt  int64               `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Cover      *basemodel.Image    `json:"cover,omitempty" bson:"cover,omitempty"`
	Type       string              `json:"type,omitempty" bson:"type,omitempty"`
	Video      *basemodel.Video    `json:"video,omitempty" bson:"video,omitempty"`
	Images     *basemodel.Images   `json:"images,omitempty" bson:"images,omitempty"`
	Tags       *basemodel.Tags     `json:"tags,omitempty" bson:"tags,omitempty"`
	State      int32               `json:"state,omitempty" bson:"state,omitempty"`
	LikeCount  int32               `json:"like_count,omitempty" bson:"like_count,omitempty"`
	ShareCount int32               `json:"share_count,omitempty" bson:"share_count,omitempty"`
	FavorCount int32               `json:"favor_count,omitempty" bson:"favor_count,omitempty"`
	ViewCount  int32               `json:"view_count,omitempty" bson:"view_count,omitempty"`
	CommentId  int64               `json:"comment_id,omitempty" bson:"comment_id,omitempty"`
}

type PostCard struct {
	Id        int64            `json:"id,omitempty" bson:"_id,omitempty"`
	Uid       int64            `json:"uid,omitempty" bson:"uid,omitempty"`
	Type      string           `json:"type,omitempty" bson:"type,omitempty"`
	Title     string           `json:"title,omitempty" bson:"title,omitempty"`
	Cover     *basemodel.Image `json:"cover,omitempty" bson:"cover,omitempty"`
	LikeCount int32            `json:"like_count,omitempty" bson:"like_count,omitempty"`
	Tags      *basemodel.Tags  `json:"tags,omitempty" bson:"tags,omitempty"`
	State     int32            `json:"state,omitempty" bson:"state,omitempty"`
}

type Posts []*Post
type PostCards []*PostCard

func (m *Post) ToDomainProto() *domainfeedv1.Post {
	pb := &domainfeedv1.Post{
		Id:         m.Id,
		Uid:        m.Uid,
		Title:      m.Title,
		Content:    m.Content,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
		Cover:      m.Cover.ToProto(),
		Type:       m.Type,
		Tags:       m.Tags.ListToProto(),
		State:      m.State,
		LikeCount:  m.LikeCount,
		ShareCount: m.ShareCount,
		FavorCount: m.FavorCount,
		ViewCount:  m.ViewCount,
		CommentId:  m.CommentId,
	}

	if m.Video != nil {
		pb.Video = m.Video.ToProto()
	}
	if m.Images != nil {
		pb.Images = m.Images.ListToProto()
	}
	return pb
}

func (m *Post) FromDomainProto(pb *domainfeedv1.Post) {
	m.Id = pb.Id
	m.Uid = pb.Uid
	m.Title = pb.Title
	m.Content = pb.Content
	m.CreatedAt = pb.CreatedAt
	m.UpdatedAt = pb.UpdatedAt

	cover := &basemodel.Image{}
	cover.FromProto(pb.Cover)
	m.Cover = cover
	m.Type = pb.Type
	if pb.Video != nil {
		video := &basemodel.Video{}
		video.FromProto(pb.Video)
		m.Video = video
	}
	if pb.Images != nil {
		images := &basemodel.Images{}
		images.ListFromProto(pb.Images)
		m.Images = images
	}
	if pb.Tags != nil {
		tags := &basemodel.Tags{}
		tags.ListFromProto(pb.Tags)
		m.Tags = tags
	}
	m.State = pb.State
	m.LikeCount = pb.LikeCount
	m.ShareCount = pb.ShareCount
	m.FavorCount = pb.FavorCount
	m.ViewCount = pb.ViewCount
	m.CommentId = pb.CommentId
}

func (ms *Posts) ListFromDomainProto(pbs []*domainfeedv1.Post) {
	for _, pb := range pbs {
		m := &Post{}
		m.FromDomainProto(pb)
		*ms = append(*ms, m)
	}
}

func (ms *Posts) ListToInterfaceProto() []*interfacefeedv1.Post {
	var pbs []*interfacefeedv1.Post
	for _, m := range *ms {
		pbs = append(pbs, m.ToInterfaceProto())
	}
	return pbs
}

func (m *Post) ToInterfaceProto() *interfacefeedv1.Post {
	pb := &interfacefeedv1.Post{
		Id:         strconv.FormatInt(m.Id, 10),
		Uid:        m.Uid,
		Title:      m.Title,
		Content:    m.Content,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
		Cover:      m.Cover.ToProto(),
		Type:       m.Type,
		Tags:       m.Tags.ListToProto(),
		State:      m.State,
		LikeCount:  m.LikeCount,
		ShareCount: m.ShareCount,
		FavorCount: m.FavorCount,
		ViewCount:  m.ViewCount,
		CommentId:  strconv.FormatInt(m.CommentId, 10),
	}
	if m.Video != nil {
		pb.Video = m.Video.ToProto()
	}
	if m.Images != nil {
		pb.Images = m.Images.ListToProto()
	}
	return pb
}

func (m *PostCard) ToDomainProto() *domainfeedv1.PostCard {
	return &domainfeedv1.PostCard{
		Id:        m.Id,
		Uid:       m.Uid,
		Type:      m.Type,
		Title:     m.Title,
		Cover:     m.Cover.ToProto(),
		LikeCount: m.LikeCount,
		Tags:      m.Tags.ListToProto(),
		State:     m.State,
	}
}

func (m *PostCard) FromDomainProto(pb *domainfeedv1.PostCard) {
	m.Id = pb.Id
	m.Uid = pb.Uid
	m.Type = pb.Type
	m.Title = pb.Title
	cover := &basemodel.Image{}
	cover.FromProto(pb.Cover)
	m.Cover = cover
	m.LikeCount = pb.LikeCount
	if pb.Tags != nil {
		m.Tags = &basemodel.Tags{}
		m.Tags.ListFromProto(pb.Tags)
	}
	m.State = pb.State
}
func (m *PostCard) ToInterfaceProto() *interfacefeedv1.PostCard {
	return &interfacefeedv1.PostCard{
		Id:        strconv.FormatInt(m.Id, 10),
		Uid:       m.Uid,
		Type:      m.Type,
		Title:     m.Title,
		Cover:     m.Cover.ToProto(),
		LikeCount: m.LikeCount,
		Tags:      m.Tags.ListToProto(),
		State:     m.State,
	}
}

func (ms *PostCards) ListToInterfaceProto() []*interfacefeedv1.PostCard {
	list := make([]*interfacefeedv1.PostCard, len(*ms))
	for i, item := range *ms {
		list[i] = item.ToInterfaceProto()
	}
	return list
}

func (ms *PostCards) ListFromDomainProto(pb []*domainfeedv1.PostCard) {
	list := make([]*PostCard, len(pb))
	for i, item := range pb {
		list[i] = &PostCard{}
		list[i].FromDomainProto(item)
	}
	*ms = list
}
