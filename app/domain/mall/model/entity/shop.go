package entity

import (
	"redbook/common/basemodel"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Shop struct {
	Id_       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Id        int64              `json:"id,omitempty" bson:"id,omitempty"`
	Uid       int64              `json:"uid,omitempty" bson:"uid,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Sign      string             `json:"sign,omitempty" bson:"sign,omitempty"`
	Logo      *basemodel.Image   `json:"logo,omitempty" bson:"logo,omitempty"`
	State     int32              `json:"state,omitempty" bson:"state,omitempty"`
	CreatedAt int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
