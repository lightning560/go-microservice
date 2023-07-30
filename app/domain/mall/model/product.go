package model

import (
	mallv1 "redbook/api/domain/mall/v1"
	"redbook/app/domain/mall/model/entity"
	"redbook/common/basemodel"
)

type Product struct {
	Id_       string            `json:"_id,omitempty" bson:"_id,omitempty"`
	Id        int64             `json:"id,omitempty" bson:"id,omitempty"`
	ShopId    int64             `json:"shop_id,omitempty" bson:"shop_id,omitempty"`
	Name      string            `json:"name,omitempty" bson:"name,omitempty"`
	Title     string            `json:"title,omitempty" bson:"title,omitempty"`
	Price     int64             `json:"price,omitempty" bson:"price,omitempty"`
	Inventory int32             `json:"inventory,omitempty" bson:"inventory,omitempty"`
	Thumb     *basemodel.Image  `json:"thumb,omitempty" bson:"thumb,omitempty"`
	Images    *basemodel.Images `json:"images,omitempty" bson:"images,omitempty"`
	Video     *basemodel.Video  `json:"video,omitempty" bson:"video,omitempty"`
	Overview  *basemodel.Images `json:"overview,omitempty" bson:"overview,omitempty"`
	Specs     string            `json:"specs,omitempty" bson:"specs,omitempty"`
	Tags      *basemodel.Tags   `json:"tags,omitempty" bson:"tags,omitempty"`
	State     int32             `json:"state,omitempty" bson:"state,omitempty"`
	CreatedAt int64             `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt int64             `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	PublishAt int64             `json:"publish_at,omitempty" bson:"publish_at,omitempty"`
	Version   int32             `json:"version,omitempty" bson:"version,omitempty"`
}

type Products []*Product

type ProductCard struct {
	Id_       string           `json:"_id,omitempty" bson:"_id,omitempty"`
	Id        int64            `json:"id,omitempty"`
	ShopId    int64            `json:"shop_id,omitempty"`
	Name      string           `json:"name,omitempty"`
	Title     string           `json:"title,omitempty"`
	Price     int64            `json:"price,omitempty"`
	Thumb     *basemodel.Image `json:"thumb,omitempty"`
	Video     *basemodel.Video `json:"video,omitempty" bson:"video,omitempty"`
	Tags      *basemodel.Tags  `json:"tags,omitempty"`
	State     int32            `json:"state,omitempty"`
	PublishAt int64            `json:"publish_at,omitempty" bson:"publish_at,omitempty"`
	Version   int32            `json:"version,omitempty" bson:"version,omitempty"`
}

type ProductCards []*ProductCard

type CreateProduct struct {
	ShopId    int64             `json:"shop_id,omitempty" bson:"shop_id,omitempty"`
	Name      string            `json:"name,omitempty" bson:"name,omitempty"`
	Title     string            `json:"title,omitempty" bson:"title,omitempty"`
	Price     int64             `json:"price,omitempty" bson:"price,omitempty"`
	Inventory int32             `json:"inventory,omitempty" bson:"inventory,omitempty"`
	Thumb     *basemodel.Image  `json:"thumb,omitempty" bson:"thumb,omitempty"`
	Images    *basemodel.Images `json:"images,omitempty" bson:"images,omitempty"`
	Video     *basemodel.Video  `json:"video,omitempty" bson:"video,omitempty"`
	Overview  *basemodel.Images `json:"overview,omitempty" bson:"overview,omitempty"`
	Specs     string            `json:"specs,omitempty" bson:"specs,omitempty"`
	Tags      *basemodel.Tags   `json:"tags,omitempty" bson:"tags,omitempty"`
	PublishAt int64             `json:"publish_at,omitempty" bson:"publish_at,omitempty"`
}

func (m *Product) FromProto(pb *mallv1.Product) {
	m.Id = pb.Id
	m.ShopId = pb.ShopId
	m.Name = pb.Name
	m.Title = pb.Title
	m.Price = pb.Price
	m.Inventory = pb.Inventory
	m.Thumb = &basemodel.Image{}
	m.Thumb.FromProto(pb.Thumb)
	m.Images = &basemodel.Images{}
	m.Images.ListFromProto(pb.Images)
	m.Video = &basemodel.Video{}
	m.Video.FromProto(pb.Video)
	m.Overview = &basemodel.Images{}
	m.Overview.ListFromProto(pb.Overview)
	m.Specs = pb.Specs
	m.Tags = &basemodel.Tags{}
	m.Tags.ListFromProto(pb.Tags)
	m.State = pb.State
	m.CreatedAt = pb.CreatedAt
	m.UpdatedAt = pb.UpdatedAt
	m.PublishAt = pb.PublishAt
	m.Version = pb.Version
}

func (ms *Products) ListFromProto(pbs []*mallv1.Product) {
	for _, pb := range pbs {
		m := &Product{}
		m.FromProto(pb)
		*ms = append(*ms, m)
	}
}

func (m *Product) ToProto() *mallv1.Product {
	return &mallv1.Product{
		Id:        m.Id,
		ShopId:    m.ShopId,
		Name:      m.Name,
		Title:     m.Title,
		Price:     m.Price,
		Inventory: m.Inventory,
		Thumb:     m.Thumb.ToProto(),
		Images:    m.Images.ListToProto(),
		Video:     m.Video.ToProto(),
		Overview:  m.Overview.ListToProto(),
		Specs:     m.Specs,
		Tags:      m.Tags.ListToProto(),
		State:     m.State,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		PublishAt: m.PublishAt,
		Version:   m.Version,
	}
}

func (m *Product) ToCard() *ProductCard {
	return &ProductCard{
		Id:        m.Id,
		ShopId:    m.ShopId,
		Name:      m.Name,
		Title:     m.Title,
		Price:     m.Price,
		Thumb:     m.Thumb,
		Video:     m.Video,
		Tags:      m.Tags,
		State:     m.State,
		PublishAt: m.PublishAt,
		Version:   m.Version,
	}
}

func (m *Product) ToEntity() *entity.Product {
	return &entity.Product{
		Id:        m.Id,
		ShopId:    m.ShopId,
		Name:      m.Name,
		Title:     m.Title,
		Price:     m.Price,
		Inventory: m.Inventory,
		Thumb:     m.Thumb,
		Images:    m.Images,
		Video:     m.Video,
		Overview:  m.Overview,
		Specs:     m.Specs,
		Tags:      m.Tags,
		State:     m.State,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		PublishAt: m.PublishAt,
		Version:   m.Version,
	}
}

func (ms *Products) ListToEntity() []*entity.Product {
	var res []*entity.Product
	for _, m := range *ms {
		res = append(res, m.ToEntity())
	}
	return res
}

func (m *Product) FromEntity(et *entity.Product) {
	m.Id_ = et.Id_.Hex()
	m.Id = et.Id
	m.ShopId = et.ShopId
	m.Name = et.Name
	m.Title = et.Title
	m.Price = et.Price
	m.Inventory = et.Inventory
	m.Thumb = et.Thumb
	m.Images = et.Images
	m.Video = et.Video
	m.Overview = et.Overview
	m.Specs = et.Specs
	m.Tags = et.Tags
	m.State = et.State
	m.CreatedAt = et.CreatedAt
	m.UpdatedAt = et.UpdatedAt
	m.PublishAt = et.PublishAt
	m.Version = et.Version
}

func (ms *Products) ListFromEntity(entities []*entity.Product) {
	for _, et := range entities {
		m := &Product{}
		m.FromEntity(et)
		*ms = append(*ms, m)
	}
}

func (m *ProductCard) FromProduct(p *Product) {
	m.Id_ = p.Id_
	m.Id = p.Id
	m.ShopId = p.ShopId
	m.Name = p.Name
	m.Title = p.Title
	m.Price = p.Price
	m.Thumb = p.Thumb
	m.Video = p.Video
	m.Tags = p.Tags
	m.State = p.State
	m.PublishAt = p.PublishAt
	m.Version = p.Version
}
func (m *ProductCard) ToProto() *mallv1.ProductCard {
	return &mallv1.ProductCard{
		Id:     m.Id,
		ShopId: m.ShopId,
		Name:   m.Name,
		Title:  m.Title,
		Price:  m.Price,
		Thumb:  m.Thumb.ToProto(),
		Video:  m.Video.ToProto(),
		Tags:   m.Tags.ListToProto(),
		State:  m.State,
	}
}

func (ms ProductCards) ListFromProducts(ps []*Product) {
	for _, p := range ps {
		m := &ProductCard{}
		m.FromProduct(p)
		ms = append(ms, m)
	}
}

func (m *CreateProduct) FromProto(pb *mallv1.CreateProduct) {
	m.ShopId = pb.ShopId
	m.Name = pb.Name
	m.Title = pb.Title
	m.Price = pb.Price
	m.Inventory = pb.Inventory
	m.Thumb = &basemodel.Image{}
	m.Thumb.FromProto(pb.Thumb)
	m.Images = &basemodel.Images{}
	m.Images.ListFromProto(pb.Images)
	m.Video = &basemodel.Video{}
	m.Video.FromProto(pb.Video)
	m.Overview = &basemodel.Images{}
	m.Overview.ListFromProto(pb.Overview)
	m.Specs = pb.Specs
	m.Tags = &basemodel.Tags{}
	m.Tags.ListFromProto(pb.Tags)
	m.PublishAt = pb.PublishAt
}
