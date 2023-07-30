package model

import (
	"encoding/hex"
	"redbook/common/basemodel"
)

// User .
type Profile struct {
	Id     int64            `json:"id,omitempty" bson:"_id,omitempty"`
	Uid    int64            `json:"uid,omitempty" bson:"uid,omitempty"`
	Name   string           `json:"name,omitempty" bson:"name,omitempty"`
	Sex    int32            `json:"sex,omitempty" bson:"sex"`
	Avatar *basemodel.Image `json:"avatar,omitempty" bson:"avatar,omitempty"`
	Sign   string           `json:"sign,omitempty" bson:"sign,omitempty"`
	State  int32            `json:"state,omitempty" bson:"state"`
	Level  int32            `json:"level,omitempty" bson:"level"`

	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

// DecodeUser .
type DecodeProfile struct {
	Id     int64           `json:"id"`
	Uid    int64           `json:"uid"`
	Name   string          `json:"name"`
	Sex    int32           `json:"sex"`
	Sign   string          `json:"sign,omitempty"`
	State  int32           `json:"state"`
	Phone  string          `json:"Phone"`
	Avatar *basemodel.Image `json:"avatar"`
	Level  int32           `json:"level"`
	Email  string          `json:"email"`
}

// Decode decode user
func (d *Profile) Decode() *DecodeProfile {
	return &DecodeProfile{
		Id:     d.Id,
		Uid:    d.Uid,
		Name:   d.Name,
		Sex:    d.Sex,
		Sign:   d.Sign,
		State:  d.State,
		Phone:  d.Phone,
		Avatar: d.Avatar,
		Level:  d.Level,
		Email:  hex.EncodeToString([]byte(d.Email)),
	}
}

func (m *Profile) ToInfo() *basemodel.UserInfo {
	return &basemodel.UserInfo{
		Id:     m.Id,
		Uid:    m.Uid,
		Name:   m.Name,
		Sex:    m.Sex,
		Avatar: m.Avatar,
		Sign:   m.Sign,
		State:  m.State,
		Level:  m.Level,
	}
}
