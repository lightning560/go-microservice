package model

import (
	domainmallv1 "redbook/api/domain/mall/v1"
	interfacemallv1 "redbook/api/interface/mall/v1"
	"redbook/common/basemodel"
	"strconv"
)

type SkuOnlyId struct {
	Index     int32  `json:"index,omitempty" bson:"index,omitempty"`
	Name      string `json:"name,omitempty" bson:"name,omitempty"`
	ProductId int64  `json:"product_id,omitempty" bson:"product_id,omitempty"`
}

type SkusOnlyId []*SkuOnlyId

type Sku struct {
	Index     int32   `json:"index,omitempty" bson:"index,omitempty"`
	Name      string  `json:"name,omitempty" bson:"name,omitempty"`
	ProductId int64   `json:"product_id,omitempty" bson:"product_id,omitempty"`
	Product   Product `json:"product,omitempty" bson:"product,omitempty"`
}

type Skus []*Sku
type Collection struct {
	Id        int64           `json:"id,omitempty" bson:"id,omitempty"`
	ShopId    int64           `json:"shop_id,omitempty" bson:"shop_id,omitempty"`
	Name      string          `json:"name,omitempty" bson:"name,omitempty"`
	Title     string          `json:"title,omitempty" bson:"title,omitempty"`
	Cover     basemodel.Image `json:"cover,omitempty" bson:"cover,omitempty"`
	Video     basemodel.Video `json:"video,omitempty" bson:"video,omitempty"`
	Tags      basemodel.Tags  `json:"tags,omitempty" bson:"tags,omitempty"`
	State     int32           `json:"state,omitempty" bson:"state,omitempty"`
	CreatedAt int64           `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt int64           `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	PublishAt int64           `json:"publish_at,omitempty" bson:"publish_at,omitempty"`
	Skus      Skus            `json:"skus,omitempty" bson:"skus,omitempty"`
	Version   int32           `json:"version,omitempty" bson:"version,omitempty"`
}

type CollectionCard struct {
	Id_     string          `json:"_id,omitempty" bson:"_id,omitempty"`
	Id      int64           `json:"id,omitempty"`
	ShopId  int64           `json:"shop_id,omitempty"`
	Name    string          `json:"name,omitempty" bson:"name,omitempty"`
	Title   string          `json:"title,omitempty" bson:"title,omitempty"`
	Cover   basemodel.Image `json:"cover,omitempty" bson:"cover,omitempty"`
	Video   basemodel.Video `json:"video,omitempty" bson:"video,omitempty"`
	Tags    basemodel.Tags  `json:"tags,omitempty"`
	State   int32           `json:"state,omitempty" bson:"state,omitempty"`
	Version int32           `json:"version,omitempty"`
	// Product is Skus[0]
	ProductId   int64       `json:"product_id,omitempty"`
	ProductCard ProductCard `json:"product_card,omitempty"`
}

type CollectionCards []*CollectionCard

type CreateCollection struct {
	ShopId     int64           `json:"shop_id,omitempty" bson:"shop_id,omitempty"`
	Name       string          `json:"name,omitempty" bson:"name,omitempty"`
	Title      string          `json:"title,omitempty" bson:"title,omitempty"`
	Cover      basemodel.Image `json:"cover,omitempty" bson:"cover,omitempty"`
	Video      basemodel.Video `json:"video,omitempty" bson:"video,omitempty"`
	Tags       basemodel.Tags  `json:"tags,omitempty" bson:"tags,omitempty"`
	PublishAt  int64           `json:"publish_at,omitempty" bson:"publish_at,omitempty"`
	SkusOnlyId SkusOnlyId      `json:"skus_only_id,omitempty" bson:"skus_only_id,omitempty"`
}

func (m *Sku) ToDomainProto() *domainmallv1.Sku {
	return &domainmallv1.Sku{
		Index:     m.Index,
		Name:      m.Name,
		ProductId: m.ProductId,
		Product:   m.Product.ToDomainProto(),
	}
}

func (m *Sku) FromDomainProto(sku *domainmallv1.Sku) {
	m.Index = sku.Index
	m.Name = sku.Name
	m.ProductId = sku.ProductId
	m.Product.FromDomainProto(sku.Product)
}

func (ms *Skus) ListToDomainProto() []*domainmallv1.Sku {
	var res []*domainmallv1.Sku
	for _, sku := range *ms {
		res = append(res, sku.ToDomainProto())
	}
	return res
}
func (ms *Skus) ListFromDomainProto(skus []*domainmallv1.Sku) {
	for _, sku := range skus {
		m := &Sku{}
		m.FromDomainProto(sku)
		*ms = append(*ms, m)
	}
}

func (m *Sku) ToInterfaceProto() *interfacemallv1.Sku {
	productIdStr := strconv.FormatInt(m.ProductId, 10)
	return &interfacemallv1.Sku{
		Index:     m.Index,
		Name:      m.Name,
		ProductId: productIdStr,
		Product:   m.Product.ToInterfaceProto(),
	}
}

func (ms *Skus) ListToInterfaceProto() []*interfacemallv1.Sku {
	var res []*interfacemallv1.Sku
	for _, sku := range *ms {
		res = append(res, sku.ToInterfaceProto())
	}
	return res
}

func (m *SkuOnlyId) FromDomainProto(sku *domainmallv1.SkuOnlyId) {
	m.Index = sku.Index
	m.ProductId = sku.ProductId
}
func (ms *SkusOnlyId) ListFromDomainProto(skus []*domainmallv1.SkuOnlyId) {
	for _, sku := range skus {
		m := &SkuOnlyId{}
		m.FromDomainProto(sku)
		*ms = append(*ms, m)
	}
}
func (m *SkuOnlyId) ToDomainProto() *domainmallv1.SkuOnlyId {
	return &domainmallv1.SkuOnlyId{
		Index:     m.Index,
		Name:      m.Name,
		ProductId: m.ProductId,
	}
}
func (ms *SkusOnlyId) ListToDomainProto() []*domainmallv1.SkuOnlyId {
	var res []*domainmallv1.SkuOnlyId
	for _, sku := range *ms {
		res = append(res, sku.ToDomainProto())
	}
	return res
}

func (m *SkuOnlyId) ToInterfaceProto() *interfacemallv1.SkuOnlyId {
	return &interfacemallv1.SkuOnlyId{
		Index:     m.Index,
		Name:      m.Name,
		ProductId: m.ProductId,
	}
}
func (ms *SkusOnlyId) ListToInterfaceProto() []*interfacemallv1.SkuOnlyId {
	var res []*interfacemallv1.SkuOnlyId
	for _, sku := range *ms {
		res = append(res, sku.ToInterfaceProto())
	}
	return res
}

func (m *SkuOnlyId) FromInterfaceProto(pb *interfacemallv1.SkuOnlyId) {
	m.Index = pb.Index
	m.Name = pb.Name
	m.ProductId = pb.ProductId
}

func (ms *SkusOnlyId) ListFromInterfaceProto(pbs []*interfacemallv1.SkuOnlyId) {
	for _, pb := range pbs {
		m := &SkuOnlyId{}
		m.FromInterfaceProto(pb)
		*ms = append(*ms, m)
	}
}

func (m *Collection) ToDomainProto() *domainmallv1.Collection {
	return &domainmallv1.Collection{
		Id:        m.Id,
		ShopId:    m.ShopId,
		Name:      m.Name,
		Title:     m.Title,
		Cover:     m.Cover.ToProto(),
		Video:     m.Video.ToProto(),
		Tags:      m.Tags.ListToProto(),
		State:     m.State,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		PublishAt: m.PublishAt,
		Version:   m.Version,
		Skus:      m.Skus.ListToDomainProto(),
	}
}

func (m *Collection) FromDomainProto(collection *domainmallv1.Collection) {
	m.Id = collection.Id
	m.ShopId = collection.ShopId
	m.Name = collection.Name
	m.Title = collection.Title
	m.Cover.FromProto(collection.Cover)
	m.Video.FromProto(collection.Video)
	m.Tags.ListFromProto(collection.Tags)
	m.State = collection.State
	m.CreatedAt = collection.CreatedAt
	m.UpdatedAt = collection.UpdatedAt
	m.PublishAt = collection.PublishAt
	m.Version = collection.Version
	m.Skus.ListFromDomainProto(collection.Skus)
}
func (m *Collection) ToInterfaceProto() *interfacemallv1.Collection {
	idStr := strconv.FormatInt(m.Id, 10)
	shopIdStr := strconv.FormatInt(m.ShopId, 10)
	return &interfacemallv1.Collection{
		Id:        idStr,
		ShopId:    shopIdStr,
		Name:      m.Name,
		Title:     m.Title,
		Cover:     m.Cover.ToProto(),
		Video:     m.Video.ToProto(),
		Tags:      m.Tags.ListToProto(),
		State:     m.State,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		PublishAt: m.PublishAt,
		Version:   m.Version,
		Skus:      m.Skus.ListToInterfaceProto(),
	}
}

func (m *CollectionCard) FromDomainProto(collection *domainmallv1.CollectionCard) {
	m.Id = collection.Id
	m.ShopId = collection.ShopId
	m.Name = collection.Name
	m.Title = collection.Title
	m.Cover.FromProto(collection.Cover)
	m.Video.FromProto(collection.Video)
	m.Tags.ListFromProto(collection.Tags)
	m.State = collection.State
	m.Version = collection.Version
	m.ProductId = collection.ProductId
	m.ProductCard.FromDomainProto(collection.ProductCard)
}

func (ms *CollectionCards) ListFromDomainProto(collectionCards []*domainmallv1.CollectionCard) {
	for _, collectionCard := range collectionCards {
		m := &CollectionCard{}
		m.FromDomainProto(collectionCard)
		*ms = append(*ms, m)
	}
}

func (m *CollectionCard) ToInterfaceProto() *interfacemallv1.CollectionCard {
	idStr := strconv.FormatInt(m.Id, 10)
	shopIdStr := strconv.FormatInt(m.ShopId, 10)
	productIdStr := strconv.FormatInt(m.ProductId, 10)
	return &interfacemallv1.CollectionCard{
		Id:          idStr,
		ShopId:      shopIdStr,
		Name:        m.Name,
		Title:       m.Title,
		Cover:       m.Cover.ToProto(),
		Video:       m.Video.ToProto(),
		Tags:        m.Tags.ListToProto(),
		State:       m.State,
		Version:     m.Version,
		ProductId:   productIdStr,
		ProductCard: m.ProductCard.ToInterfaceProto(),
	}
}

func (ms *CollectionCards) ListToInterfaceProto() []*interfacemallv1.CollectionCard {
	var res []*interfacemallv1.CollectionCard
	for _, collectionCard := range *ms {
		res = append(res, collectionCard.ToInterfaceProto())
	}
	return res
}

func (m *CreateCollection) FromInterfaceProto(pb *interfacemallv1.CreateCollection) {
	m.ShopId = pb.ShopId
	m.Name = pb.Name
	m.Title = pb.Title
	m.Cover.FromProto(pb.Cover)
	m.Video.FromProto(pb.Video)
	m.Tags.ListFromProto(pb.Tags)
	m.PublishAt = pb.PublishAt
	m.SkusOnlyId.ListFromInterfaceProto(pb.SkusOnlyId)
}

func (m *CreateCollection) ToDomainProto() (pb *domainmallv1.CreateCollection) {
	pb = &domainmallv1.CreateCollection{
		ShopId:     m.ShopId,
		Name:       m.Name,
		Title:      m.Title,
		Cover:      m.Cover.ToProto(),
		Video:      m.Video.ToProto(),
		Tags:       m.Tags.ListToProto(),
		PublishAt:  m.PublishAt,
		SkusOnlyId: m.SkusOnlyId.ListToDomainProto(),
	}
	return
}
