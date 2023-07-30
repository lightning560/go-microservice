package server

import (
	"context"
	mallv1 "redbook/api/domain/mall/v1"
	"redbook/app/domain/mall/model"
)

func (s *MallServer) CreateProduct(c context.Context, req *mallv1.CreateProductReq) (*mallv1.CreateProductResponse, error) {
	// m := &model.Product{
	// 	ShopId: req.Product.ShopId,
	// 	Name:   req.Product.Name,
	// 	Title:  req.Product.Title,
	// 	Price:  req.Product.Price,
	// 	Thumb: basemodel.Image{
	// 		Url:    req.Product.Thumb.Url,
	// 		Hash:   req.Product.Thumb.Hash,
	// 		Name:   req.Product.Thumb.Name,
	// 		Width:  req.Product.Thumb.Width,
	// 		Height: req.Product.Thumb.Height,
	// 		SizeKb: req.Product.Thumb.SizeKb,
	// 	},
	// 	Video: basemodel.Video{
	// 		Url:    req.Product.Video.Url,
	// 		Type:   req.Product.Video.Type,
	// 		Hash:   req.Product.Video.Hash,
	// 		Name:   req.Product.Video.Name,
	// 		Length: req.Product.Video.Length,
	// 	},
	// 	Inventory: req.Product.Inventory,
	// }

	// for _, v := range req.Product.Images {
	// 	m.Images = append(m.Images, &basemodel.Image{
	// 		Url:    v.Url,
	// 		Hash:   v.Hash,
	// 		Name:   v.Name,
	// 		Width:  v.Width,
	// 		Height: v.Height,
	// 		SizeKb: v.SizeKb,
	// 	})
	// }

	m := &model.CreateProduct{}
	m.FromProto(req.CreateProduct)
	rv, err := s.productBiz.CreateProduct(c, m)
	if err != nil {
		return nil, err
	}
	return &mallv1.CreateProductResponse{
		Id: rv,
	}, nil
}

func (s *MallServer) GetProductById(c context.Context, req *mallv1.GetProductByIdReq) (*mallv1.GetProductByIdResponse, error) {
	rv, err := s.productBiz.GetProductById(c, req.Id)
	if err != nil {
		return nil, err
	}
	// res := &mallv1.GetProductByIdResponse{
	// 	Product: &mallv1.Product{
	// 		Id:     rv.Id,
	// 		ShopId: rv.ShopId,
	// 		Name:   rv.Name,
	// 		Title:  rv.Title,
	// 		Price:  rv.Price,
	// 		Thumb: &basev1.Image{
	// 			Url:    rv.Thumb.Url,
	// 			Hash:   rv.Thumb.Hash,
	// 			Name:   rv.Thumb.Name,
	// 			Width:  rv.Thumb.Width,
	// 			Height: rv.Thumb.Height,
	// 			SizeKb: rv.Thumb.SizeKb,
	// 		},
	// 		Video: &basev1.Video{
	// 			Url:    rv.Video.Url,
	// 			Type:   rv.Video.Type,
	// 			Hash:   rv.Video.Hash,
	// 			Name:   rv.Video.Name,
	// 			Length: rv.Video.Length,
	// 		},
	// 		Inventory: rv.Inventory,
	// 	},
	// }
	// for _, v := range rv.Images {
	// 	res.Product.Images = append(res.Product.Images, &basev1.Image{
	// 		Url:    v.Url,
	// 		Hash:   v.Hash,
	// 		Name:   v.Name,
	// 		Width:  v.Width,
	// 		Height: v.Height,
	// 		SizeKb: v.SizeKb,
	// 	})
	// }
	// for _, v := range rv.Tags {
	// 	res.Product.Tags = append(res.Product.Tags, &basev1.Tag{
	// 		TagId: v.TagId,
	// 		Name:  v.Name,
	// 	})
	// }

	res := &mallv1.GetProductByIdResponse{}
	res.Product = rv.ToProto()
	return res, nil
}

func (s *MallServer) MapProductByIds(c context.Context, req *mallv1.MapProductByIdsReq) (*mallv1.MapProductByIdsResponse, error) {
	rv, err := s.productBiz.ListProductByIds(c, req.Ids)
	if err != nil {
		return nil, err
	}
	products := make(map[int64]*mallv1.Product)
	for _, v := range rv {
		products[v.Id] = v.ToProto()
	}
	// for k, v := range rv {
	// 	p := &mallv1.Product{
	// 		Id:     v.Id,
	// 		ShopId: v.ShopId,
	// 		Name:   v.Name,
	// 		Title:  v.Title,
	// 		Price:  v.Price,
	// 		Thumb: &basev1.Image{
	// 			Url:    v.Thumb.Url,
	// 			Hash:   v.Thumb.Hash,
	// 			Name:   v.Thumb.Name,
	// 			Width:  v.Thumb.Width,
	// 			Height: v.Thumb.Height,
	// 			SizeKb: v.Thumb.SizeKb,
	// 		},
	// 		Video: &basev1.Video{
	// 			Url:    v.Video.Url,
	// 			Type:   v.Video.Type,
	// 			Hash:   v.Video.Hash,
	// 			Name:   v.Video.Name,
	// 			Length: v.Video.Length,
	// 		},
	// 		Inventory: v.Inventory,
	// 	}

	// 	for _, v := range v.Images {
	// 		p.Images = append(p.Images, &basev1.Image{
	// 			Url:    v.Url,
	// 			Hash:   v.Hash,
	// 			Name:   v.Name,
	// 			Width:  v.Width,
	// 			Height: v.Height,
	// 			SizeKb: v.SizeKb,
	// 		})
	// 	}
	// 	for _, v := range v.Tags {
	// 		p.Tags = append(p.Tags, &basev1.Tag{
	// 			TagId: v.TagId,
	// 			Name:  v.Name,
	// 		})
	// 	}
	// 	products[k] = p
	// }

	return &mallv1.MapProductByIdsResponse{
		Products: products,
		Total:    int32(len(rv)),
	}, nil
}
