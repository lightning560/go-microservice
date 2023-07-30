package entity

import (
	"redbook/common/basemodel"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id_        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Id         int64              `json:"id,omitempty" bson:"id,omitempty"`
	Uid        int64              `json:"uid,omitempty" bson:"uid,omitempty"`
	Title      string             `json:"title,omitempty" bson:"title,omitempty"`
	Content    string             `json:"content,omitempty" bson:"content,omitempty"`
	CreatedAt  int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt  int64              `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Cover      *basemodel.Image   `json:"cover,omitempty" bson:"cover,omitempty"`
	Type       string             `json:"type,omitempty" bson:"type,omitempty"`
	Video      *basemodel.Video   `json:"video,omitempty" bson:"video,omitempty"`
	Images     *basemodel.Images  `json:"images,omitempty" bson:"images,omitempty"`
	Tags       *basemodel.Tags    `json:"tags,omitempty" bson:"tags,omitempty"`
	Deleted    int8               `json:"deleted,omitempty" bson:"deleted,"`
	State      int32              `json:"state,omitempty" bson:"state,"`
	LikeCount  int32              `json:"like_count,omitempty" bson:"like_count,"`
	ShareCount int32              `json:"share_count,omitempty" bson:"share_count,"`
	FavorCount int32              `json:"favor_count,omitempty" bson:"favor_count,"`
	ViewCount  int32              `json:"view_count,omitempty" bson:"view_count,"`
	CommentId  int64              `json:"comment_id,omitempty" bson:"comment_id,omitempty"`
}
type Posts []*Post

type PostCard struct {
	Id_       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Id        int64              `json:"id,omitempty" bson:"id,omitempty"`
	Uid       int64              `json:"uid,omitempty" bson:"uid,omitempty"`
	Type      string             `json:"type,omitempty" bson:"type,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	Cover     *basemodel.Image   `json:"cover,omitempty" bson:"cover,omitempty"`
	State     int32              `json:"state,omitempty" bson:"state,omitempty"`
	LikeCount int32              `json:"like_count,omitempty" bson:"like_count,omitempty"`
	Tags      *basemodel.Tags    `json:"tags,omitempty" bson:"tags,omitempty"`
}
type PostCards []*PostCard
