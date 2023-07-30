package basemodel

import (
	basev1 "redbook/api/redbookpb/v1"
)

type Pagination struct {
	Page int32 `json:"page,omitempty"`
	Size int32 `json:"size,omitempty"`
}
type Cursor struct {
	Offset int32 `json:"offset,omitempty"`
	Limit  int32 `json:"limit,omitempty"`
}

func (m *Pagination) FromProto(pb *basev1.Pagination) {
	m.Page = pb.Page
	m.Size = pb.Size
}

func (m *Cursor) FromProto(pb *basev1.Cursor) {
	m.Offset = pb.Offset
	m.Limit = pb.Limit
}

func (m *Pagination) ToProto() *basev1.Pagination {
	return &basev1.Pagination{
		Page: m.Page,
		Size: m.Size,
	}
}
func (m *Cursor) ToProto() *basev1.Cursor {
	return &basev1.Cursor{
		Offset: m.Offset,
		Limit:  m.Limit,
	}
}
