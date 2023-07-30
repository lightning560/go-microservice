package entity

import (
	"redbook/common/basemodel"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reply struct {
	Id_       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Id        int64              `json:"id,omitempty" bson:"id,omitempty"`
	SubjectId int64              `json:"subject_id,omitempty" bson:"subject_id,omitempty"`
	OwnerUid  int64              `json:"owner_uid,omitempty" bson:"owner_uid,omitempty"`
	// 0 is FloorRoot, >0 is Reply
	FloorId      int64  `json:"floor_id,omitempty" bson:"floor_id,omitempty"`
	BizType      string `json:"Biz_type,omitempty" bson:"biz_type,omitempty"`
	Content      string `json:"content,omitempty" bson:"content,omitempty"`
	AtUid        int64  `json:"at_uid,omitempty" bson:"at_uid,omitempty"`
	LikeCount    int32  `json:"like_count,omitempty" bson:"like_count"`
	DislikeCount int32  `json:"dislike_count,omitempty" bson:"dislike_count"`

	Deleted   int32 `json:"deleted,omitempty" bson:"deleted"`
	State     int32 `json:"state,omitempty" bson:"state"`
	CreatedAt int64 `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt int64 `json:"updated_at,omitempty" bson:"updated_at"`

	FanGrade int32 `json:"fan_grade,omitempty" bson:"fan_grade,omitempty"`

	Platform int32  `json:"platform,omitempty" bson:"platform,omitempty"`
	Device   string `json:"device,omitempty" bson:"device,omitempty"`

	Attr   int64 `json:"attr,omitempty" bson:"attr"`
	Dialog int64 `json:"dialog,omitempty" bson:"dialog,omitempty"`

	// blow part is for floorRoot only
	FloorAttr *FloorAttr `json:"floor_attr,omitempty" bson:"floor_attr,omitempty"`
}

type FloorAttr struct {
	ReplyCount int32 `json:"reply_count,omitempty" bson:"reply_count"`
	PinAdmin   int32 `json:"pin_admin,omitempty" bson:"pin_admin"`
	PinOwner   int32 `json:"pin_owner,omitempty" bson:"pin_owner"`
	Fold       int32 `json:"fold,omitempty" bson:"fold"`
	Hot        bool  `json:"hot,omitempty" bson:"hot"`
}
type Replies []*Reply

type Floor struct {
	RootReply Reply `json:"floor_root,omitempty" bson:"floor_root,omitempty"`

	Replies Replies          `json:"replies,omitempty" bson:"replies,omitempty"`
	Cursor  basemodel.Cursor `json:"cursor,omitempty" bson:"cursor,omitempty"`
}

type Floors []*Floor
