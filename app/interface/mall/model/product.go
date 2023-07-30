package model

import (
	domainmallv1 "redbook/api/domain/mall/v1"
	interfacemallv1 "redbook/api/interface/mall/v1"
	"redbook/common/basemodel"
	"strconv"
)

type Product struct {
	Id        int64
	ShopId    int64
	Name      string
	Title     string
	Price     int64
	Inventory int32
	Thumb     basemodel.Image
	Images    basemodel.Images
	Video     basemodel.Video
	Overview  basemodel.Images
	Specs     string
	Tags      basemodel.Tags
	State     int32
	CreatedAt int64
	UpdatedAt int64
	PublishAt int64
	Version   int32
}

type Products []*Product

type CreateProduct struct {
	ShopId    int64            `json:"shop_id,omitempty" bson:"shop_id,omitempty"`
	Name      string           `json:"name,omitempty" bson:"name,omitempty"`
	Title     string           `json:"title,omitempty" bson:"title,omitempty"`
	Price     int64            `json:"price,omitempty" bson:"price,omitempty"`
	Inventory int32            `json:"inventory,omitempty" bson:"inventory,omitempty"`
	Thumb     basemodel.Image  `json:"thumb,omitempty" bson:"thumb,omitempty"`
	Images    basemodel.Images `json:"images,omitempty" bson:"images,omitempty"`
	Video     basemodel.Video  `json:"video,omitempty" bson:"video,omitempty"`
	Overview  basemodel.Images `json:"overview,omitempty" bson:"overview,omitempty"`
	Specs     string           `json:"specs,omitempty" bson:"specs,omitempty"`
	Tags      basemodel.Tags   `json:"tags,omitempty" bson:"tags,omitempty"`
	PublishAt int64            `json:"publish_at,omitempty" bson:"publish_at,omitempty"`
}
type ProductCard struct {
	Id        int64           `json:"id,omitempty"`
	ShopId    int64           `json:"shop_id,omitempty"`
	Name      string          `json:"name,omitempty"`
	Title     string          `json:"title,omitempty"`
	Price     int64           `json:"price,omitempty"`
	Thumb     basemodel.Image `json:"thumb,omitempty"`
	Video     basemodel.Video `json:"video,omitempty" bson:"video,omitempty"`
	Tags      basemodel.Tags  `json:"tags,omitempty"`
	State     int32           `json:"state,omitempty"`
	PublishAt int64           `json:"publish_at,omitempty" bson:"publish_at,omitempty"`
	Version   int32           `json:"version,omitempty" bson:"version,omitempty"`
}

type ProductCards []*ProductCard

func (m *Product) FromDomainProto(proto *domainmallv1.Product) {
	m.Id = proto.Id
	m.ShopId = proto.ShopId
	m.Name = proto.Name
	m.Title = proto.Title
	m.Price = proto.Price
	m.Inventory = proto.Inventory
	m.Thumb.FromProto(proto.Thumb)
	m.Images.ListFromProto(proto.Images)
	m.Video.FromProto(proto.Video)
	m.Overview.ListFromProto(proto.Overview)
	m.Specs = proto.Specs
	m.Tags.ListFromProto(proto.Tags)
	m.State = proto.State
	m.CreatedAt = proto.CreatedAt
	m.UpdatedAt = proto.UpdatedAt
}

func (m *Product) ToDomainProto() *domainmallv1.Product {
	return &domainmallv1.Product{
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

func (m *Product) ToInterfaceProto() *interfacemallv1.Product {
	idStr := strconv.FormatInt(m.Id, 10)
	shopIdStr := strconv.FormatInt(m.ShopId, 10)
	return &interfacemallv1.Product{
		Id:        idStr,
		ShopId:    shopIdStr,
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

func (ms Products) ListFromDomainProto(protos []*domainmallv1.Product) {
	for _, proto := range protos {
		m := &Product{}
		m.FromDomainProto(proto)
		ms = append(ms, m)
	}
}

func (m *CreateProduct) FromInterfaceProto(pb *interfacemallv1.CreateProduct) {
	m.ShopId = pb.ShopId
	m.Name = pb.Name
	m.Title = pb.Title
	m.Price = pb.Price
	m.Inventory = pb.Inventory
	m.Thumb.FromProto(pb.Thumb)
	m.Images.ListFromProto(pb.Images)
	m.Video.FromProto(pb.Video)
	m.Overview.ListFromProto(pb.Overview)
	m.Specs = pb.Specs
	m.Tags.ListFromProto(pb.Tags)
	m.PublishAt = pb.PublishAt
}
func (m *CreateProduct) ToDomainProto() *domainmallv1.CreateProduct {
	return &domainmallv1.CreateProduct{
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
		PublishAt: m.PublishAt,
	}
}

func (m *ProductCard) FromDomainProto(pb *domainmallv1.ProductCard) {
	m.Id = pb.Id
	m.ShopId = pb.ShopId
	m.Name = pb.Name
	m.Title = pb.Title
	m.Price = pb.Price
	m.Thumb.FromProto(pb.Thumb)
	m.Video.FromProto(pb.Video)
	m.Tags.ListFromProto(pb.Tags)
	m.State = pb.State
	m.PublishAt = pb.PublishAt
	m.Version = pb.Version
}

func (ms *ProductCards) ListFromDomainProto(protos []*domainmallv1.ProductCard) {
	for _, proto := range protos {
		m := &ProductCard{}
		m.FromDomainProto(proto)
		*ms = append(*ms, m)
	}
}

func (m *ProductCard) ToInterfaceProto() (pb *interfacemallv1.ProductCard) {
	idStr := strconv.FormatInt(m.Id, 10)
	shopIdStr := strconv.FormatInt(m.ShopId, 10)
	pb = &interfacemallv1.ProductCard{}
	pb.Id = idStr
	pb.ShopId = shopIdStr
	pb.Name = m.Name
	pb.Title = m.Title
	pb.Price = m.Price
	pb.Thumb = m.Thumb.ToProto()
	pb.Video = m.Video.ToProto()
	pb.Tags = m.Tags.ListToProto()
	pb.State = m.State
	pb.PublishAt = m.PublishAt
	pb.Version = m.Version
	return
}

func (ms *ProductCards) ListToInterfaceProto() (pbs []*interfacemallv1.ProductCard) {
	for _, m := range *ms {
		pbs = append(pbs, m.ToInterfaceProto())
	}
	return
}
