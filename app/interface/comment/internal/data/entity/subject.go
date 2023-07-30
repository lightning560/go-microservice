package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Subject struct {
	Id         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	OwnerUid   int64              `json:"owner_uid,omitempty" bson:"owner_uid,omitempty"`
	BelongId   int64              `json:"belong_id,omitempty" bson:"belong_id,omitempty"`
	BizType    string             `json:"biz_type,omitempty" bson:"biz_type,omitempty"`
	FloorCount int64              `json:"floor_count,omitempty" bson:"floor_count,omitempty"`
	ReplyCount int64              `json:"reply_count,omitempty" bson:"reply_count,omitempty"`
	State      int64              `json:"state,omitempty" bson:"state,omitempty"`
	Attr       int64              `json:"attr,omitempty" bson:"attr,omitempty"`
	Meta       string             `json:"meta,omitempty" bson:"meta,omitempty"`
	CreatedAt  int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt  int64              `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
