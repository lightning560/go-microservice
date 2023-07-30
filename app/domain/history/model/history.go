package model

type History struct {
	Id        int64  `json:"id,omitempty"`
	Uid       int64  `json:"uid,omitempty"`
	ItemId    int64  `json:"item_id,omitempty"`
	BizType   string `json:"biz_type,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
}
