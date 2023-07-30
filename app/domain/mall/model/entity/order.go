package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type OrderItem struct {
	Id_          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CollectionId int64              `json:"collection_id,omitempty" bson:"collection_id,omitempty"`
	ProductId    int64              `json:"product_id,omitempty" bson:"product_id,omitempty"`
	Quantity     int32              `json:"quantity,omitempty" bson:"quantity,omitempty"`
}

type Order struct {
	Id_        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Id         int64              `json:"id,omitempty" bson:"id,omitempty"`
	Uid        int64              `json:"Uid,omitempty" bson:"Uid,omitempty"`
	OrderItems []OrderItem        `json:"order_items,omitempty" bson:"order_items,omitempty"`
	State      int32              `json:"state,omitempty" bson:"state,omitempty"`
	CreatedAt  int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt  int64              `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
