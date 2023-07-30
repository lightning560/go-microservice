package model

import (
	domainmallv1 "redbook/api/domain/mall/v1"
	interfacemallv1 "redbook/api/interface/mall/v1"
	"redbook/common/basemodel"
	"strconv"
)

type Shop struct {
	Id        int64
	Uid       int64
	Name      string
	Sign      string
	Logo      basemodel.Image
	State     int32
	CreatedAt int64
}

func (m *Shop) FromDomainProto(entity *domainmallv1.Shop) {
	m.Id = entity.Id
	m.Uid = entity.Uid
	m.Name = entity.Name
	m.Sign = entity.Sign
	m.Logo.FromProto(entity.Logo)
	m.State = entity.State
	m.CreatedAt = entity.CreatedAt
}

func (m *Shop) ToDomainProto() *domainmallv1.Shop {
	pb := &domainmallv1.Shop{}
	pb.Id = m.Id
	pb.Uid = m.Uid
	pb.Name = m.Name
	pb.Sign = m.Sign
	pb.Logo = m.Logo.ToProto()
	pb.State = m.State
	pb.CreatedAt = m.CreatedAt
	return pb
}

func (m *Shop) FromInterfaceProto(pb *interfacemallv1.Shop) {
	id, _ := strconv.ParseInt(pb.Id, 10, 64)
	uid, _ := strconv.ParseInt(pb.Uid, 10, 64)
	m.Id = id
	m.Uid = uid
	m.Name = pb.Name
	m.Sign = pb.Sign
	m.Logo.FromProto(pb.Logo)
	m.State = pb.State
	m.CreatedAt = pb.CreatedAt
}

func (m *Shop) ToInterfaceProto() *interfacemallv1.Shop {
	idStr := strconv.FormatInt(m.Id, 10)
	uidStr := strconv.FormatInt(m.Uid, 10)
	pb := &interfacemallv1.Shop{}
	pb.Id = idStr
	pb.Uid = uidStr
	pb.Name = m.Name
	pb.Sign = m.Sign
	pb.Logo = m.Logo.ToProto()
	pb.State = m.State
	pb.CreatedAt = m.CreatedAt
	return pb
}
