package model

import (
	"fmt"
	mallv1 "redbook/api/domain/mall/v1"
	"redbook/app/domain/mall/model/entity"
	"redbook/common/basemodel"
)

type Shop struct {
	Id_       string           `json:"_id,omitempty" bson:"_id,omitempty"`
	Id        int64            `json:"id,omitempty" bson:"id,omitempty"`
	Uid       int64            `json:"uid,omitempty" bson:"uid,omitempty"`
	Name      string           `json:"name,omitempty" bson:"name,omitempty"`
	Sign      string           `json:"sign,omitempty" bson:"sign,omitempty"`
	Logo      *basemodel.Image `json:"logo,omitempty" bson:"logo,omitempty"`
	State     int32            `json:"state,omitempty" bson:"state,omitempty"`
	CreatedAt int64            `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

func (m *Shop) ToEntity() *entity.Shop {

	return &entity.Shop{
		Id:        m.Id,
		Uid:       m.Uid,
		Name:      m.Name,
		Sign:      m.Sign,
		Logo:      m.Logo,
		State:     m.State,
		CreatedAt: m.CreatedAt,
	}
}

func (m *Shop) FromEntity(entity *entity.Shop) {
	m.Id = entity.Id
	fmt.Println("entity.Shop.id_", entity.Id_)
	fmt.Println("entity.Shop.id_ hex()", entity.Id_.Hex())
	fmt.Println("Id", entity.Id)
	// m.Id_ = entity.Id_.Hex()
	m.Id = entity.Id
	fmt.Println("entity.Uid", entity.Uid)
	m.Uid = entity.Uid
	fmt.Println("entity.Name", entity.Name)
	m.Name = entity.Name
	fmt.Println("entity.Sign", entity.Sign)
	m.Sign = entity.Sign
	fmt.Println("entity.Logo", entity.Logo)
	m.Logo = entity.Logo
	fmt.Println("entity.State", entity.State)
	m.State = entity.State
	fmt.Println("entity.CreatedAt", entity.CreatedAt)
	m.CreatedAt = entity.CreatedAt
}

func (m *Shop) ToProto() *mallv1.Shop {
	return &mallv1.Shop{
		Id:        m.Id,
		Uid:       m.Uid,
		Name:      m.Name,
		Sign:      m.Sign,
		Logo:      m.Logo.ToProto(),
		State:     m.State,
		CreatedAt: m.CreatedAt,
	}
}

func (m *Shop) FromProto(pb *mallv1.Shop) {
	m.Id = pb.Id
	m.Uid = pb.Uid
	m.Name = pb.Name
	m.Sign = pb.Sign
	m.Logo = &basemodel.Image{}
	m.Logo.FromProto(pb.Logo)
	m.State = pb.State
	m.CreatedAt = pb.CreatedAt
}
