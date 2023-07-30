package entity

import (
	"redbook/common/basemodel"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id_       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Id        int64              `json:"id,omitempty" bson:"id,omitempty"`
	ShopId    int64              `json:"shop_id,omitempty" bson:"shop_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Price     int64              `json:"price,omitempty" bson:"price"`
	Inventory int32              `json:"inventory,omitempty" bson:"inventory,omitempty"`
	Thumb     *basemodel.Image   `json:"thumb,omitempty" bson:"thumb,omitempty"`
	Images    *basemodel.Images  `json:"images,omitempty" bson:"images,omitempty"`
	Video     *basemodel.Video   `json:"video,omitempty" bson:"video,omitempty"`
	Overview  *basemodel.Images  `json:"overview,omitempty" bson:"overview,omitempty"`
	Specs     string             `json:"specs,omitempty" bson:"specs,omitempty"`
	Tags      *basemodel.Tags    `json:"tags,omitempty" bson:"tags"`
	State     int32              `json:"state,omitempty" bson:"state"`
	CreatedAt int64              `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt int64              `json:"updated_at,omitempty" bson:"updated_at"`
	PublishAt int64              `json:"publish_at,omitempty" bson:"publish_at"`
	Version   int32              `json:"version,omitempty" bson:"version"`
}
