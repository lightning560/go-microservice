package model

import (
	domaincommentv1 "redbook/api/domain/comment/v1"
	interfacecommentv1 "redbook/api/interface/comment/v1"
	"redbook/common/basemodel"
	"strconv"
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

	FloorAttr FloorAttr `json:"floor_attr,omitempty" bson:"floor_attr,omitempty"`
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

func (m *FloorAttr) FromDomainProto(floorAttr *domaincommentv1.FloorAttr) {
	m.ReplyCount = floorAttr.ReplyCount
	m.PinAdmin = floorAttr.PinAdmin
	m.PinOwner = floorAttr.PinOwner
	m.Fold = floorAttr.Fold
	m.Hot = floorAttr.Hot
}

func (m *FloorAttr) ToInterfaceProto() *interfacecommentv1.FloorAttr {
	return &interfacecommentv1.FloorAttr{
		ReplyCount: m.ReplyCount,
		PinAdmin:   m.PinAdmin,
		PinOwner:   m.PinOwner,
		Fold:       m.Fold,
		Hot:        m.Hot,
	}
}

func (m *Reply) FromDomainProto(reply *domaincommentv1.Reply) {
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
	m.FloorAttr.FromDomainProto(reply.FloorAttr)
}
func (ms *Replies) ListFromDomainProto(replies []*domaincommentv1.Reply) {
	for _, reply := range replies {
		m := &Reply{}
		m.FromDomainProto(reply)
		*ms = append(*ms, m)
	}
}

func (m *Reply) ToInterfaceProto() *interfacecommentv1.Reply {
	id := strconv.FormatInt(m.Id, 10)
	subjectId := strconv.FormatInt(m.SubjectId, 10)
	ownerUid := strconv.FormatInt(m.OwnerUid, 10)
	floorId := strconv.FormatInt(m.FloorId, 10)
	atUid := strconv.FormatInt(m.AtUid, 10)
	return &interfacecommentv1.Reply{
		Id:           id,
		SubjectId:    subjectId,
		BizType:      m.BizType,
		OwnerUid:     ownerUid,
		FloorId:      floorId,
		Content:      m.Content,
		AtUid:        atUid,
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
		FloorAttr:    m.FloorAttr.ToInterfaceProto(),
	}
}

func (ms *Replies) ListToInterfaceProto() []*interfacecommentv1.Reply {
	var replies []*interfacecommentv1.Reply
	for _, m := range *ms {
		replies = append(replies, m.ToInterfaceProto())
	}
	return replies
}

func (m *Floor) FromDomainProto(floor *domaincommentv1.Floor) {
	m.RootReply.FromDomainProto(floor.RootReply)
	m.Replies.ListFromDomainProto(floor.Replies)
	m.Total = floor.Total
	m.Cursor.FromProto(floor.Cursor)
}

func (ms *Floors) ListFromDomainProto(floors []*domaincommentv1.Floor) {
	for _, floor := range floors {
		m := &Floor{}
		m.FromDomainProto(floor)
		*ms = append(*ms, m)
	}
}

func (m *Floor) ToInterfaceProto() *interfacecommentv1.Floor {
	return &interfacecommentv1.Floor{
		RootReply: m.RootReply.ToInterfaceProto(),
		Replies:   m.Replies.ListToInterfaceProto(),
		Total:     m.Total,
		Cursor:    m.Cursor.ToProto(),
	}
}

func (ms *Floors) ListToInterfaceProto() []*interfacecommentv1.Floor {
	var floors []*interfacecommentv1.Floor
	for _, m := range *ms {
		floors = append(floors, m.ToInterfaceProto())
	}
	return floors
}
