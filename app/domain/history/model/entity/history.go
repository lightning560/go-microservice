package entity

type History struct {
	Id        int64  `json:"id,omitempty" bson:"id,omitempty"`
	Uid       int64  `json:"uid,omitempty" bson:"uid,omitempty"`
	ItemId    int64  `json:"item_id,omitempty" bson:"item_id,omitempty"`
	BizType   string `json:"biz_type,omitempty" bson:"biz_type,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
