package entity

import (
	"redbook/common/basemodel"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SkuOnlyId struct {
	Index     int32  `json:"index,omitempty" bson:"index"`
	Name      string `json:"name,omitempty" bson:"name"`
	ProductId int64  `json:"product_id,omitempty" bson:"product_id"`
}

type SkusOnlyId []*SkuOnlyId

type Collection struct {
	Id_        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Id         int64              `json:"id,omitempty" bson:"id,omitempty"`
	ShopId     int64              `json:"shop_id,omitempty" bson:"shop_id,omitempty"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty"`
	Title      string             `json:"title,omitempty" bson:"title,omitempty"`
	Cover      *basemodel.Image   `json:"cover,omitempty" bson:"cover"`
	Video      *basemodel.Video   `json:"video,omitempty" bson:"video"`
	Tags       *basemodel.Tags    `json:"tags,omitempty" bson:"tags"`
	State      int32              `json:"state,omitempty" bson:"state"`
	CreatedAt  int64              `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt  int64              `json:"updated_at,omitempty" bson:"updated_at"`
	PublishAt  int64              `json:"publish_at,omitempty" bson:"publish_at"`
	Version    int32              `json:"version,omitempty" bson:"version"`
	SkusOnlyId *SkusOnlyId        `json:"skus,omitempty" bson:"skus"`
}

type Collections []*Collection
