package basemodel

import basev1 "redbook/api/redbookpb/v1"

type UserInfo struct {
	Id     int64  `json:"id,omitempty" bson:"_id,omitempty"`
	Uid    int64  `json:"uid,omitempty" bson:"uid,omitempty"`
	Name   string `json:"name,omitempty" bson:"name,omitempty"`
	Sex    int32  `json:"sex,omitempty" bson:"sex,omitempty"`
	Avatar *Image `json:"avatar,omitempty" bson:"avatar,omitempty"`
	Sign   string `json:"sign,omitempty" bson:"sign,omitempty"`
	State  int32  `json:"state,omitempty" bson:"state,omitempty"`
	Level  int32  `json:"level,omitempty" bson:"level,omitempty"`
}
type UserInfos []*UserInfo

func (m *UserInfo) ToProto() *basev1.UserInfo {
	return &basev1.UserInfo{
		Id:     m.Id,
		Uid:    m.Uid,
		Name:   m.Name,
		Sex:    m.Sex,
		Avatar: m.Avatar.ToProto(),
		Sign:   m.Sign,
		State:  m.State,
		Level:  m.Level,
	}
}

func (ms *UserInfos) ListToProto() []*basev1.UserInfo {
	var res []*basev1.UserInfo
	for _, m := range *ms {
		res = append(res, m.ToProto())
	}
	return res
}

func (m *UserInfo) FromProto(pb *basev1.UserInfo) {
	m.Id = pb.Id
	m.Uid = pb.Uid
	m.Name = pb.Name
	m.Sex = pb.Sex
	m.Avatar = &Image{}
	m.Avatar.FromProto(pb.Avatar)
	m.Sign = pb.Sign
	m.State = pb.State
	m.Level = pb.Level
}

func (ms *UserInfos) ListFromProto(pb []*basev1.UserInfo) {
	for _, m := range pb {
		var model UserInfo
		model.FromProto(m)
		*ms = append(*ms, &model)
	}
}
