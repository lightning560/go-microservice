package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Relation struct {
	Id_         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Id          int64              `json:"id,omitempty" bson:"id,omitempty"`
	FollowerUid int64              `json:"follower_uid,omitempty" bson:"follower_uid,omitempty"`
	FolloweeUid int64              `json:"followee_uid,omitempty" bson:"followee_uid,omitempty"`
	CreatedAt   int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
