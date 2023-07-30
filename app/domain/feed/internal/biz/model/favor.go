package model

type Favor struct {
	Uid       int64 `json:"uid,omitempty" bson:"uid,omitempty"`
	PostId    int64 `json:"post_id,omitempty" bson:"post_id,omitempty"`
	CreatedAt int64 `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
