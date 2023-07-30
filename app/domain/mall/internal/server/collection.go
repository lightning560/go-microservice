package server

import (
	"context"
	"fmt"
	"miopkg/errors"

	mallv1 "redbook/api/domain/mall/v1"
	"redbook/app/domain/mall/model"
)

func (s *MallServer) CreateCollection(c context.Context, req *mallv1.CreateCollectionReq) (*mallv1.CreateCollectionResponse, error) {
	// m := &model.Collection{
	// 	Id:     req.Collection.Id,
	// 	ShopId: req.Collection.ShopId,
	// 	Cover: basemodel.Image{
	// 		Url:    req.Collection.Cover.Url,
	// 		Name:   req.Collection.Cover.Name,
	// 		Hash:   req.Collection.Cover.Hash,
	// 		Width:  req.Collection.Cover.Width,
	// 		Height: req.Collection.Cover.Height,
	// 		SizeKb: req.Collection.Cover.SizeKb,
	// 	},
	// }
	// for _, v := range req.Collection.Skus {
	// 	m.Skus = append(m.Skus, &entity.Sku{
	// 		Index:     v.Index,
	// 		ProductId: v.ProductId,
	// 	})
	// }
	// for _, v := range req.Collection.Tags {
	// 	m.Tags = append(m.Tags, &basemodel.Tag{
	// 		TagId: v.TagId,
	// 		Name:  v.Name,
	// 	})
	// }
	m := &model.CreateCollection{}
	m.FromProto(req.CreateCollection)
	rv, err := s.collectionBiz.CreateCollection(c, m)
	if err != nil {
		return nil, err
	}
	return &mallv1.CreateCollectionResponse{
		Id: rv,
	}, nil
}

func (s *MallServer) GetCollectionById(c context.Context, req *mallv1.GetCollectionByIdReq) (*mallv1.GetCollectionByIdResponse, error) {
	m, err := s.collectionBiz.GetCollectionById(c, req.Id)
	if err != nil {
		return nil, err
	}
	fmt.Printf("m.Collection: %+v\n", m.Skus)
	pids := make([]int64, 0)
	for _, v := range *m.Skus {
		pids = append(pids, v.ProductId)
	}
	productList, err := s.productBiz.ListProductByIds(c, pids)
	if err != nil {
		return nil, errors.Wrapf(err, "ListProductByIds failed, ids: %v", pids)
	}
	productMap := make(map[int64]*model.Product)
	for _, v := range productList {
		productMap[v.Id] = v
	}
	res := &mallv1.GetCollectionByIdResponse{
		Collection: m.ToProto(),
	}
	for k, v := range res.Collection.Skus {
		p := productMap[v.ProductId]
		res.Collection.Skus[k].Product = p.ToProto()
	}

	// for _, v := range m.Skus {
	// 	mp := products[v.ProductId]
	// 	p := &mallv1.Product{
	// 		Id:        mp.Id,
	// 		ShopId:    mp.ShopId,
	// 		Name:      mp.Name,
	// 		Price:     mp.Price,
	// 		Inventory: mp.Inventory,
	// 		Title:     mp.Title,
	// 		Specs:     mp.Specs,
	// 		Thumb: &basev1.Image{
	// 			Url:    mp.Thumb.Url,
	// 			Name:   mp.Thumb.Name,
	// 			Hash:   mp.Thumb.Hash,
	// 			Width:  mp.Thumb.Width,
	// 			Height: mp.Thumb.Height,
	// 			SizeKb: mp.Thumb.SizeKb,
	// 		},
	// 		Video: &basev1.Video{
	// 			Url:    mp.Video.Url,
	// 			Name:   mp.Video.Name,
	// 			Hash:   mp.Video.Hash,
	// 			Width:  mp.Video.Width,
	// 			Height: mp.Video.Height,
	// 			SizeKb: mp.Video.SizeKb,
	// 		},
	// 	}
	// 	for _, v := range mp.Images {
	// 		p.Images = append(p.Images, &basev1.Image{
	// 			Url:    v.Url,
	// 			Name:   v.Name,
	// 			Hash:   v.Hash,
	// 			Width:  v.Width,
	// 			Height: v.Height,
	// 			SizeKb: v.SizeKb,
	// 		})
	// 	}
	// 	for _, v := range mp.Tags {
	// 		p.Tags = append(p.Tags, &basev1.Tag{
	// 			TagId: v.TagId,
	// 			Name:  v.Name,
	// 		})
	// 	}
	// 	for _, v := range mp.Overview {
	// 		p.Overview = append(p.Overview, &basev1.Image{
	// 			Url:    v.Url,
	// 			Name:   v.Name,
	// 			Hash:   v.Hash,
	// 			Width:  v.Width,
	// 			Height: v.Height,
	// 			SizeKb: v.SizeKb,
	// 		})
	// 	}
	// 	res.Collection.Skus = append(res.Collection.Skus, &mallv1.Sku{
	// 		Index:     v.Index,
	// 		ProductId: v.ProductId,
	// 		Product:   p,
	// 	})
	// }

	// for _, v := range m.Tags {
	// 	res.Collection.Tags = append(res.Collection.Tags, &basev1.Tag{
	// 		TagId: v.TagId,
	// 		Name:  v.Name,
	// 	})
	// }
	return res, nil
}

func (s *MallServer) MapCollectionCardByIds(c context.Context, req *mallv1.MapCollectionCardByIdsReq) (*mallv1.MapCollectionCardByIdsResponse, error) {
	ms, err := s.collectionBiz.ListCollectionByIds(c, req.Ids)
	if err != nil {
		return nil, err
	}
	res := &mallv1.MapCollectionCardByIdsResponse{}
	res.CollectionCards = make(map[int64]*mallv1.CollectionCard)
	for _, v := range ms {
		// TODO: use ids to get products
		skus := *v.Skus
		productId := skus[0].ProductId
		product, err := s.productBiz.GetProductById(c, productId)
		if err != nil {
			return nil, err
		}
		res.CollectionCards[v.Id] = v.ToCard().ToProto()
		res.CollectionCards[v.Id].ProductCard = product.ToCard().ToProto()
	}
	return res, nil
}

func (s *MallServer) MapCollectionCardByShopId(c context.Context, req *mallv1.MapCollectionCardByShopIdReq) (*mallv1.MapCollectionCardByShopIdResponse, error) {
	ms, err := s.collectionBiz.ListCollectionByShopId(c, req.Id, req.State, req.Cursor.Offset, req.Cursor.Limit, req.Sort.By, req.Sort.Order)
	if err != nil {
		return nil, err
	}

	res := &mallv1.MapCollectionCardByShopIdResponse{
		CollectionCards: make(map[int64]*mallv1.CollectionCard),
	}
	for _, v := range ms {
		// TODO: use ids to get products
		skus := *v.Skus
		productId := skus[0].ProductId
		product, err := s.productBiz.GetProductById(c, productId)
		if err != nil {
			return nil, err
		}
		res.CollectionCards[v.Id] = v.ToCard().ToProto()
		res.CollectionCards[v.Id].ProductCard = product.ToCard().ToProto()
	}
	return res, nil
}
