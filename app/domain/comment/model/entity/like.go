package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Like struct {
	Id_       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ReplyId   int64              `json:"reply_id,omitempty" bson:"reply_id,omitempty"`
	Uid       int64              `json:"uid,omitempty" bson:"uid,omitempty"`
	CreatedAt int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
type Dislike struct {
	Id_       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ReplyId   int64              `json:"reply_id,omitempty" bson:"reply_id,omitempty"`
	Uid       int64              `json:"uid,omitempty" bson:"uid,omitempty"`
	CreatedAt int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
