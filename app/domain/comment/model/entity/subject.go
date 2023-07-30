package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Subject struct {
	Id_      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Id       int64              `json:"id,omitempty" bson:"id,omitempty"`
	OwnerUid int64              `json:"owner_uid,omitempty" bson:"owner_uid,omitempty"`
	BelongId int64              `json:"belong_id,omitempty" bson:"belong_id,omitempty"`

	BizType    string `json:"type,omitempty" bson:"type,omitempty"`
	FloorCount int32  `json:"floor_count,omitempty" bson:"floor_count"`
	ReplyCount int32  `json:"reply_count,omitempty" bson:"reply_count"`
	State      int32  `json:"state,omitempty" bson:"state"`

	CreatedAt int64 `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt int64 `json:"updated_at,omitempty" bson:"updated_at"`

	Attr int64  `json:"attr,omitempty" bson:"attr"`
	Meta string `json:"meta,omitempty" bson:"meta,omitempty"`
}
