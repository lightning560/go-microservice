package model

import (
	mallv1 "redbook/api/domain/mall/v1"
	"redbook/app/domain/mall/model/entity"
	"redbook/common/basemodel"
)

type Sku struct {
	Index     int32    `json:"index,omitempty" bson:"index,omitempty"`
	Name      string   `json:"name,omitempty" bson:"name,omitempty"`
	ProductId int64    `json:"product_id,omitempty" bson:"product_id,omitempty"`
	Product   *Product `json:"product,omitempty" bson:"product,omitempty"`
}
type Skus []*Sku

type SkuOnlyId struct {
	Index     int32  `json:"index,omitempty" bson:"index,omitempty"`
	Name      string `json:"name,omitempty" bson:"name,omitempty"`
	ProductId int64  `json:"product_id,omitempty" bson:"product_id,omitempty"`
}
type SkusOnlyId []*SkuOnlyId

type Collection struct {
	Id_       string           `json:"_id,omitempty" bson:"_id,omitempty"`
	Id        int64            `json:"id,omitempty" bson:"id,omitempty"`
	ShopId    int64            `json:"shop_id,omitempty" bson:"shop_id,omitempty"`
	Name      string           `json:"name,omitempty" bson:"name,omitempty"`
	Title     string           `json:"title,omitempty" bson:"title,omitempty"`
	Cover     *basemodel.Image `json:"cover,omitempty" bson:"cover,omitempty"`
	Video     *basemodel.Video `json:"video,omitempty" bson:"video,omitempty"`
	Tags      *basemodel.Tags  `json:"tags,omitempty" bson:"tags,omitempty"`
	State     int32            `json:"state,omitempty" bson:"state,omitempty"`
	CreatedAt int64            `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt int64            `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	PublishAt int64            `json:"publish_at,omitempty" bson:"publish_at,omitempty"`
	Skus      *Skus            `json:"skus,omitempty" bson:"skus,omitempty"`
	Version   int32            `json:"version,omitempty" bson:"version,omitempty"`
}
type Collections []*Collection

type CollectionCard struct {
	Id_     string           `json:"_id,omitempty" bson:"_id,omitempty"`
	Id      int64            `json:"id,omitempty"`
	ShopId  int64            `json:"shop_id,omitempty"`
	Name    string           `json:"name,omitempty" bson:"name,omitempty"`
	Title   string           `json:"title,omitempty" bson:"title,omitempty"`
	Cover   *basemodel.Image `json:"cover,omitempty" bson:"cover,omitempty"`
	Video   *basemodel.Video `json:"video,omitempty" bson:"video,omitempty"`
	Tags    *basemodel.Tags  `json:"tags,omitempty"`
	State   int32            `json:"state,omitempty" bson:"state,omitempty"`
	Version int32            `json:"version,omitempty"`
	// Product is Skus[0]
	ProductId   int64        `json:"product_id,omitempty"`
	ProductCard *ProductCard `json:"product_card,omitempty"`
}

type CollectionCards []*CollectionCard

type CreateCollection struct {
	ShopId     int64            `json:"shop_id,omitempty" bson:"shop_id,omitempty"`
	Name       string           `json:"name,omitempty" bson:"name,omitempty"`
	Title      string           `json:"title,omitempty" bson:"title,omitempty"`
	Cover      *basemodel.Image `json:"cover,omitempty" bson:"cover,omitempty"`
	Video      *basemodel.Video `json:"video,omitempty" bson:"video,omitempty"`
	Tags       *basemodel.Tags  `json:"tags,omitempty" bson:"tags,omitempty"`
	PublishAt  int64            `json:"publish_at,omitempty" bson:"publish_at,omitempty"`
	SkusOnlyId *SkusOnlyId      `json:"skus_only_id,omitempty" bson:"skus_only_id,omitempty"`
}

func (m *CollectionCard) ToProto() *mallv1.CollectionCard {
	if m.ProductCard == nil {
		return &mallv1.CollectionCard{
			Id:        m.Id,
			ShopId:    m.ShopId,
			Name:      m.Name,
			Title:     m.Title,
			Cover:     m.Cover.ToProto(),
			Video:     m.Video.ToProto(),
			Tags:      m.Tags.ListToProto(),
			State:     m.State,
			Version:   m.Version,
			ProductId: m.ProductId,
		}
	}
	return &mallv1.CollectionCard{
		Id:          m.Id,
		ShopId:      m.ShopId,
		Name:        m.Name,
		Title:       m.Title,
		Cover:       m.Cover.ToProto(),
		Video:       m.Video.ToProto(),
		Tags:        m.Tags.ListToProto(),
		State:       m.State,
		Version:     m.Version,
		ProductId:   m.ProductId,
		ProductCard: m.ProductCard.ToProto(),
	}
}

func (m *Sku) FromSkuOnlyId(sku *entity.SkuOnlyId) {
	m.Index = sku.Index
	m.Name = sku.Name
	m.ProductId = sku.ProductId
}

func (ms *Skus) ListFromSkuOnlyId(skus *entity.SkusOnlyId) {
	for _, sku := range *skus {
		m := &Sku{}
		m.FromSkuOnlyId(sku)
		*ms = append(*ms, m)
	}
}

func (m *Sku) ToEntitySkuOnlyId() *entity.SkuOnlyId {
	return &entity.SkuOnlyId{
		Index:     m.Index,
		Name:      m.Name,
		ProductId: m.ProductId,
	}
}

func (ms *Skus) ListToEntitySkuOnlyId() *entity.SkusOnlyId {
	skus := make(entity.SkusOnlyId, 0, len(*ms))
	for _, m := range *ms {
		skus = append(skus, m.ToEntitySkuOnlyId())
	}
	return &skus
}

func (m *SkuOnlyId) ToEntity() *entity.SkuOnlyId {
	return &entity.SkuOnlyId{
		Index:     m.Index,
		Name:      m.Name,
		ProductId: m.ProductId,
	}
}
func (m *SkuOnlyId) FromProto(pb *mallv1.SkuOnlyId) {
	m.Index = pb.Index
	m.Name = pb.Name
	m.ProductId = pb.ProductId
}

func (ms *SkusOnlyId) ListToEntity() *entity.SkusOnlyId {
	skus := make(entity.SkusOnlyId, 0, len(*ms))
	for _, m := range *ms {
		m.ToEntity()
		skus = append(skus, m.ToEntity())
	}
	return &skus
}

func (ms *SkusOnlyId) ListFromProto(pbs []*mallv1.SkuOnlyId) {
	for _, pb := range pbs {
		m := &SkuOnlyId{}
		m.FromProto(pb)
		*ms = append(*ms, m)
	}
}

func (m *Sku) FromProto(pb *mallv1.Sku) {
	m.Index = pb.Index
	m.Name = pb.Name
	m.ProductId = pb.ProductId
}

func (ms *Skus) ListFromProto(pbs []*mallv1.Sku) {
	for _, pb := range pbs {
		m := &Sku{}
		m.FromProto(pb)
		*ms = append(*ms, m)
	}
}

// TODO:use skusOnlyId
func (m *Sku) ToProto() *mallv1.Sku {
	if m.Product == nil {
		return &mallv1.Sku{
			Index:     m.Index,
			Name:      m.Name,
			ProductId: m.ProductId,
		}
	}
	return &mallv1.Sku{
		Index:     m.Index,
		Name:      m.Name,
		ProductId: m.ProductId,
		Product:   m.Product.ToProto(),
	}
}

func (ms *Skus) ListToProto() []*mallv1.Sku {
	pbs := make([]*mallv1.Sku, 0, len(*ms))
	for _, m := range *ms {
		pbs = append(pbs, m.ToProto())
	}
	return pbs
}

// ToCard convert Collection to CollectionCard
// but without ProductCard
func (m *Collection) ToCard() *CollectionCard {
	skus := *m.Skus
	productId := skus[0].ProductId
	return &CollectionCard{
		Id:        m.Id,
		ShopId:    m.ShopId,
		Name:      m.Name,
		Title:     m.Title,
		Cover:     m.Cover,
		Video:     m.Video,
		Tags:      m.Tags,
		State:     m.State,
		Version:   m.Version,
		ProductId: productId,
	}
}

func (m *Collection) ToEntity() (*entity.Collection, error) {
	skus := m.Skus
	return &entity.Collection{
		Id:         m.Id,
		ShopId:     m.ShopId,
		Name:       m.Name,
		Title:      m.Title,
		Cover:      m.Cover,
		Video:      m.Video,
		Tags:       m.Tags,
		State:      m.State,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
		PublishAt:  m.PublishAt,
		Version:    m.Version,
		SkusOnlyId: skus.ListToEntitySkuOnlyId(),
	}, nil
}

func (ms *Collections) ListToEntity() ([]*entity.Collection, error) {
	ets := make([]*entity.Collection, 0, len(*ms))
	for _, m := range *ms {
		et, err := m.ToEntity()
		if err != nil {
			return nil, err
		}
		ets = append(ets, et)
	}
	return ets, nil
}

func (m *Collection) FromEntity(et *entity.Collection) error {
	m.Id_ = et.Id_.Hex()
	m.Id = et.Id
	m.ShopId = et.ShopId
	m.Name = et.Name
	m.Title = et.Title
	// m.Cover = &basemodel.Image{}
	m.Cover = et.Cover
	// m.Video = &basemodel.Video{}
	m.Video = et.Video
	// m.Tags = &basemodel.Tags{}
	m.Tags = et.Tags
	m.State = et.State
	m.CreatedAt = et.CreatedAt
	m.UpdatedAt = et.UpdatedAt
	m.PublishAt = et.PublishAt
	m.Version = et.Version
	m.Skus = &Skus{}
	m.Skus.ListFromSkuOnlyId(et.SkusOnlyId)
	return nil
}
func (ms *Collections) ListFromEntity(ets []*entity.Collection) error {
	for _, et := range ets {
		m := &Collection{}
		if err := m.FromEntity(et); err != nil {
			return err
		}
		*ms = append(*ms, m)
	}
	return nil
}

func (ms *Collections) MapFromEntity(ets map[int64]*entity.Collection) error {
	for _, et := range ets {
		m := &Collection{}
		if err := m.FromEntity(et); err != nil {
			return err
		}
		*ms = append(*ms, m)
	}
	return nil
}

func (m *Collection) ToProto() *mallv1.Collection {
	return &mallv1.Collection{
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
		Skus:      m.Skus.ListToProto(),
	}
}

func (ms *Collections) ListToProto() []*mallv1.Collection {
	pbs := make([]*mallv1.Collection, 0, len(*ms))
	for _, m := range *ms {
		pbs = append(pbs, m.ToProto())
	}
	return pbs
}

func (m *Collection) FromProto(pb *mallv1.Collection) {
	m.Id = pb.Id
	m.ShopId = pb.ShopId
	m.Name = pb.Name
	m.Title = pb.Title
	m.Cover = &basemodel.Image{}
	m.Cover.FromProto(pb.Cover)
	m.Video = &basemodel.Video{}
	m.Video.FromProto(pb.Video)
	m.Tags = &basemodel.Tags{}
	m.Tags.ListFromProto(pb.Tags)
	m.Skus = &Skus{}
	m.Skus.ListFromProto(pb.Skus)
}

func (m *CreateCollection) FromProto(pb *mallv1.CreateCollection) {
	m.ShopId = pb.ShopId
	m.Name = pb.Name
	m.Title = pb.Title
	m.Cover = &basemodel.Image{}
	m.Cover.FromProto(pb.Cover)
	m.Video = &basemodel.Video{}
	m.Video.FromProto(pb.Video)
	m.Tags = &basemodel.Tags{}
	m.Tags.ListFromProto(pb.Tags)
	m.PublishAt = pb.PublishAt
	m.SkusOnlyId = &SkusOnlyId{}
	m.SkusOnlyId.ListFromProto(pb.SkusOnlyId)
}
