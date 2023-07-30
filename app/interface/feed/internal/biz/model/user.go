package model

import "redbook/common/basemodel"

type UserProflie struct {
	Id     int64
	Uid    int64
	Name   string
	Avatar *basemodel.Image `json:"avatar,omitempty"`
}
