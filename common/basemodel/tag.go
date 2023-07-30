package basemodel

import basev1 "redbook/api/redbookpb/v1"

type Tag struct {
	TagId   int64  `json:"tag_id,omitempty" bson:"tag_id,omitempty"`
	Name    string `json:"name,omitempty" bson:"name,omitempty"`
	BizType string `json:"biz_type,omitempty" bson:"biz_type,omitempty"`
	// State     int32  `json:"state,omitempty" bson:"state,omitempty"`
	// Hot       int32  `json:"hot,omitempty" bson:"hot,omitempty"`
	// Count     int64  `json:"count,omitempty" bson:"count,omitempty"`
	// CreatedAt int64  `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type Tags []*Tag

func (m *Tag) ToProto() *basev1.Tag {
	return &basev1.Tag{
		TagId:   m.TagId,
		Name:    m.Name,
		BizType: m.BizType,
	}
}

func (ms *Tags) ListToProto() []*basev1.Tag {
	pbs := make([]*basev1.Tag, 0, len(*ms))
	for _, m := range *ms {
		pbs = append(pbs, m.ToProto())
	}
	return pbs
}

func (m *Tag) FromProto(pb *basev1.Tag) {
	m.TagId = pb.TagId
	m.Name = pb.Name
	m.BizType = pb.BizType
}

func (ms *Tags) ListFromProto(pbs []*basev1.Tag) {
	for _, pb := range pbs {
		m := &Tag{}
		m.FromProto(pb)
		*ms = append(*ms, m)
	}
}
