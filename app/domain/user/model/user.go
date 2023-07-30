package model

import (
	"encoding/hex"
)

// User .
type User struct {
	Id    int64  `json:"id"`
	Uid   int64  `json:"uid"`
	Phone string `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
	// Pwd      []byte `json:"pwd"`
	Password  string `json:"password"`
	Salt      string `json:"salt"`
	Status    int8   `json:"status"`
	CreatedAt int64  `json:"created_at"`
}

// DecodeUser .
type DecodeUser struct {
	Id        int64  `json:"id"`
	Uid       int64  `json:"uid"`
	Phone     string `json:"phone,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password"`
	Salt      string `json:"salt"`
	Status    int8   `json:"status"`
	CreatedAt int64  `json:"created_at"`
}

// Decode decode user
func (d *User) Decode() *DecodeUser {
	return &DecodeUser{
		Id:        d.Id,
		Uid:       d.Uid,
		Phone:     hex.EncodeToString([]byte(d.Phone)),
		Email:     hex.EncodeToString([]byte(d.Email)),
		Password:  hex.EncodeToString([]byte(d.Password)),
		Salt:      d.Salt,
		Status:    d.Status,
		CreatedAt: d.CreatedAt,
	}
}
