package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Share struct {
	Id_       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Id        int64              `json:"id,omitempty" bson:"id,omitempty"`
	Uid       int64              `json:"uid,omitempty" bson:"uid,omitempty"`
	PostId    int64              `json:"post_id,omitempty" bson:"post_id,omitempty"`
	CreatedAt int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
