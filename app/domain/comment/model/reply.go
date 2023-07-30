package model

import (
	commentv1 "redbook/api/domain/comment/v1"
	"redbook/app/domain/comment/model/entity"
	"redbook/common/basemodel"
)

// Reply Reply
type Reply struct {
	Id_          string `json:"_id,omitempty" bson:"_id,omitempty"`
	Id           int64  `json:"id" bson:"id,omitempty"`
	SubjectId    int64  `json:"subject_id,omitempty" bson:"subject_id,omitempty"`
	BizType      string `json:"Biz_type,omitempty" bson:"biz_type,omitempty"`
	OwnerUid     int64  `json:"owner_uid,omitempty" bson:"owner_uid,omitempty"`
	FloorId      int64  `json:"floor_id,omitempty" bson:"floor_id,omitempty"`
	Content      string `json:"content,omitempty" bson:"content,omitempty"`
	AtUid        int64  `json:"at_uid,omitempty" bson:"at_uid,omitempty"`
	LikeCount    int32  `json:"like_count,omitempty" bson:"like_count,omitempty"`
	DislikeCount int32  `json:"dislike_count,omitempty" bson:"dislike_count,omitempty"`

	Deleted int32 `json:"deleted,omitempty" bson:"deleted,omitempty"`
	State   int32 `json:"state,omitempty" bson:"state,omitempty"`

	FanGrade int32  `json:"fan_grade,omitempty" bson:"fan_grade,omitempty"`
	Platform int32  `json:"platform,omitempty" bson:"platform,omitempty"`
	Device   string `json:"device,omitempty" bson:"device,omitempty"`

	CreatedAt int64 `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt int64 `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Attr      int64 `json:"attr,omitempty" bson:"attr,omitempty"`

	Dialog int64 `json:"dialog,omitempty" bson:"dialog,omitempty"`
	Assist int   `json:"assist,omitempty" bson:"assist,omitempty"`

	FloorAttr *FloorAttr `json:"floor_attr,omitempty" bson:"floor_attr,omitempty"`
}
type FloorAttr struct {
	ReplyCount int32 `json:"reply_count,omitempty" bson:"reply_count,omitempty"`
	PinAdmin   int32 `json:"pin_admin,omitempty" bson:"pin_admin, omitempty"`
	PinOwner   int32 `json:"pin_owner,omitempty" bson:"pin_owner, omitempty"`
	Fold       int32 `json:"fold,omitempty" bson:"fold,omitempty"`
	Hot        bool  `json:"hot,omitempty" bson:"hot,omitempty"`
}

type Replies []*Reply

type Floor struct {
	RootReply Reply `json:"floor_root,omitempty" bson:"floor_root,omitempty"`

	Replies Replies          `json:"replies,omitempty" bson:"replies,omitempty"`
	Total   int32            `json:"total,omitempty" bson:"total,omitempty"`
	Cursor  basemodel.Cursor `json:"cursor,omitempty" bson:"cursor,omitempty"`
}

type Floors []*Floor

func (m *FloorAttr) FromEntity(floorAttr *entity.FloorAttr) {
	m.Fold = floorAttr.Fold
	m.PinAdmin = floorAttr.PinAdmin
	m.PinOwner = floorAttr.PinOwner
	m.Hot = floorAttr.Hot
	m.ReplyCount = floorAttr.ReplyCount
}

func (m *FloorAttr) ToProto() *commentv1.FloorAttr {
	return &commentv1.FloorAttr{
		Fold:       m.Fold,
		PinAdmin:   m.PinAdmin,
		PinOwner:   m.PinOwner,
		Hot:        m.Hot,
		ReplyCount: m.ReplyCount,
	}
}

func (m *Reply) FromEntity(reply *entity.Reply) {
	m.Id_ = reply.Id_.Hex()
	m.Id = reply.Id
	m.SubjectId = reply.SubjectId
	m.BizType = reply.BizType
	m.OwnerUid = reply.OwnerUid
	m.FloorId = reply.FloorId
	m.Content = reply.Content
	m.AtUid = reply.AtUid
	m.LikeCount = reply.LikeCount
	m.DislikeCount = reply.DislikeCount
	m.Deleted = reply.Deleted
	m.State = reply.State
	m.FanGrade = reply.FanGrade
	m.Platform = reply.Platform
	m.Device = reply.Device
	m.CreatedAt = reply.CreatedAt
	m.UpdatedAt = reply.UpdatedAt
	m.Attr = reply.Attr
	m.Dialog = reply.Dialog

	m.FloorAttr = &FloorAttr{}
	if reply.FloorId == 0 {
		m.FloorAttr.FromEntity(reply.FloorAttr)
	}
}

func (ms *Replies) ListFromEntity(replies []*entity.Reply) {
	for _, reply := range replies {
		m := &Reply{}
		m.FromEntity(reply)
		*ms = append(*ms, m)
	}
}

func (m *Reply) ToProto() *commentv1.Reply {
	return &commentv1.Reply{
		Id:           m.Id,
		SubjectId:    m.SubjectId,
		BizType:      m.BizType,
		OwnerUid:     m.OwnerUid,
		FloorId:      m.FloorId,
		Content:      m.Content,
		AtUid:        m.AtUid,
		LikeCount:    m.LikeCount,
		DislikeCount: m.DislikeCount,
		Deleted:      m.Deleted,
		State:        m.State,
		FanGrade:     m.FanGrade,
		Platform:     m.Platform,
		Device:       m.Device,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
		Attr:         m.Attr,
		Dialog:       m.Dialog,
		FloorAttr:    m.FloorAttr.ToProto(),
	}
}

func (ms *Replies) ListToProto() []*commentv1.Reply {
	var replies []*commentv1.Reply
	for _, m := range *ms {
		replies = append(replies, m.ToProto())
	}
	return replies
}

func (m *Floor) ToProto() *commentv1.Floor {
	return &commentv1.Floor{
		RootReply: m.RootReply.ToProto(),
		Replies:   m.Replies.ListToProto(),
		Total:     m.Total,
		Cursor:    m.Cursor.ToProto(),
	}
}

func (ms *Floors) ListToProto() []*commentv1.Floor {
	var floors []*commentv1.Floor
	for _, m := range *ms {
		floors = append(floors, m.ToProto())
	}
	return floors
}
