package model

type Reply struct {
	Id        int64 `json:"id" bson:"_id,omitempty"`
	SubjectId int64 `json:"subject_id,omitempty" bson:"subject_id,omitempty"`

	OwnerUid     int64  `json:"owner_uid,omitempty" bson:"owner_uid,omitempty"`
	BizType      string `json:"Biz_type,omitempty" bson:"biz_type,omitempty"`
	FloorId      int64  `json:"floor_id,omitempty" bson:"floor_id,omitempty"`
	Dialog       int64  `json:"dialog,omitempty" bson:"dialog,omitempty"`
	Content      string `json:"content,omitempty" bson:"content,omitempty"`
	AtUid        int64  `json:"at_uid,omitempty" bson:"at_uid,omitempty"`
	LikeCount    int64  `json:"like_count,omitempty" bson:"like_count,omitempty"`
	DislikeCount int64  `json:"dislike_count,omitempty" bson:"dislike_count,omitempty"`
	State        int32  `json:"state,omitempty" bson:"state,omitempty"`

	Attr     int64 `json:"attr,omitempty" bson:"attr,omitempty"`
	FanGrade int32 `json:"fans_grade,omitempty" bson:"fans_grade,omitempty"`

	Platform int32  `json:"platform,omitempty" bson:"platform,omitempty"`
	Device   string `json:"device,omitempty" bson:"device,omitempty"`

	CreatedAt int64 `json:"create_at,omitempty" bson:"create_at,omitempty"`
	UpdatedAt int64 `json:"update_at,omitempty" bson:"update_at,omitempty"`
}
