package basemodel

import (
	basev1 "redbook/api/redbookpb/v1"
)

type Image struct {
	Id     int64  `json:"id,omitempty" bson:"id,omitempty"`
	Url    string `json:"url,omitempty" bson:"url,omitempty"`
	Hash   string `json:"hash,omitempty" bson:"hash,omitempty"`
	Name   string `json:"name,omitempty" bson:"name,omitempty"`
	SizeKb int32  `json:"size_kb,omitempty" bson:"size_kb,omitempty"`
	Width  int32  `json:"width,omitempty" bson:"width,omitempty"`
	Height int32  `json:"height,omitempty" bson:"height,omitempty"`
}

type Images []*Image

func (m *Image) ToProto() *basev1.Image {
	return &basev1.Image{
		Id:     m.Id,
		Url:    m.Url,
		Hash:   m.Hash,
		Name:   m.Name,
		SizeKb: m.SizeKb,
		Width:  m.Width,
		Height: m.Height,
	}
}


func (ms *Images) ListToProto() []*basev1.Image {
	pbs := make([]*basev1.Image, 0)
	for _, m := range *ms {
		pbs = append(pbs, m.ToProto())
	}
	return pbs
}

func (m *Image) FromProto(pb *basev1.Image) {
	m.Id = pb.Id
	m.Url = pb.Url
	m.Hash = pb.Hash
	m.Name = pb.Name
	m.SizeKb = pb.SizeKb
	m.Width = pb.Width
	m.Height = pb.Height
}

func (ms *Images) ListFromProto(pbs []*basev1.Image) {
	for _, pb := range pbs {
		m := &Image{}
		m.FromProto(pb)
		*ms = append(*ms, m)
	}
}
