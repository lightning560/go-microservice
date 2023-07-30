package basemodel

import (
	basev1 "redbook/api/redbookpb/v1"
)

type Video struct {
	Id        int64  `json:"id,omitempty" bson:"id,omitempty"`
	Url       string `json:"url,omitempty" bson:"url,omitempty"`
	Type      string `json:"type,omitempty" bson:"type,omitempty"`
	Cover     *Image `json:"cover,omitempty" bson:"cover,omitempty"`
	Hash      string `json:"hash,omitempty" bson:"hash,omitempty"`
	Name      string `json:"name,omitempty" bson:"name,omitempty"`
	SizeKb    int32  `json:"size_kb,omitempty" bson:"size_kb,omitempty"`
	Width     int32  `json:"width,omitempty" bson:"width,omitempty"`
	Height    int32  `json:"height,omitempty" bson:"height,omitempty"`
	Length    int32  `json:"length,omitempty" bson:"length,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

func (m *Video) ToProto() *basev1.Video {
	return &basev1.Video{
		Id:        m.Id,
		Url:       m.Url,
		Type:      m.Type,
		Cover:     m.Cover.ToProto(),
		Hash:      m.Hash,
		Name:      m.Name,
		SizeKb:    m.SizeKb,
		Width:     m.Width,
		Height:    m.Height,
		Length:    m.Length,
		CreatedAt: m.CreatedAt,
	}
}

func (m *Video) FromProto(pb *basev1.Video) {
	m.Id = pb.Id
	m.Url = pb.Url
	m.Type = pb.Type
	cover := &Image{}
	cover.FromProto(pb.Cover)
	m.Cover = cover
	m.Hash = pb.Hash
	m.Name = pb.Name
	m.SizeKb = pb.SizeKb
	m.Width = pb.Width
	m.Height = pb.Height
	m.Length = pb.Length
	m.CreatedAt = pb.CreatedAt
}
