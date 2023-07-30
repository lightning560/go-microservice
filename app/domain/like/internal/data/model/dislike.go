package model

type Dislike struct {
	Oid      int64 `json:"oid,omitempty" bson:"oid,omitempty"`
	Mid      int64 `json:"mid,omitempty" bson:"mid,omitempty"`
	Bid      int64 `json:"bid,omitempty" bson:"bid,omitempty"`
	Sid      int64 `json:"sid,omitempty" bson:"sid,omitempty"`
	Dislike  int64 `json:"dislike,omitempty" bson:"dislike,omitempty"`
	State    int64 `json:"state,omitempty" bson:"state,omitempty"`
	CreateAt int64 `json:"create_at,omitempty" bson:"create_at,omitempty"`
	UpdateAt int64 `json:"update_at,omitempty" bson:"update_at,omitempty"`
}
