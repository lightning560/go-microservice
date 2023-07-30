package entity

import "redbook/common/basemodel"

// AddReplyReq is a request struct for AddReply
type AddReply struct {
	IsFloor   bool   `json:"is_floor,omitempty" bson:"is_floor,omitempty"`
	SubjectId int64  `json:"subject_id,omitempty" bson:"subject_id,omitempty"`
	BizType   string `json:"Biz_type,omitempty" bson:"biz_type,omitempty"`
	OwnerUid  int64  `json:"owner_uid,omitempty" bson:"owner_uid,omitempty"`
	FloorId   int64  `json:"floor_id,omitempty" bson:"floor_id,omitempty"`
	Content   string `json:"content,omitempty" bson:"content,omitempty"`
	AtUid     int64  `json:"at_uid,omitempty" bson:"at_uid,omitempty"`
}

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
