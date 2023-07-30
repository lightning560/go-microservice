package server

import (
	"context"
	"fmt"
	feedv1 "redbook/api/domain/feed/v1"
	basev1 "redbook/api/redbookpb/v1"
	"redbook/app/domain/feed/internal/biz/model"
	"redbook/common/basemodel"
)

func (s *FeedServer) CreatePost(c context.Context, req *feedv1.CreatePostReq) (*feedv1.CreatePostResponse, error) {
	tags := &basemodel.Tags{}
	tags.ListFromProto(req.Post.Tags)
	imgs := &basemodel.Images{}
	imgs.ListFromProto(req.Post.Images)
	cover := &basemodel.Image{}
	cover.FromProto(req.Post.Cover)
	video := &basemodel.Video{}
	video.FromProto(req.Post.Video)
	m := &model.Post{
		Uid:     req.Post.Uid,
		Title:   req.Post.Title,
		Content: req.Post.Content,
		Type:    req.Post.Type,
		Tags:    tags,
		Images:  imgs,
		Cover:   cover,
		Video:   video,
	}
	fmt.Printf("domain # CreatePost m.tags:%+v\n", m.Tags)
	_, err := s.feedbiz.CreatePost(c, m)
	return &feedv1.CreatePostResponse{}, err
}

func (s *FeedServer) UpdatePost(c context.Context, req *feedv1.UpdatePostReq) (*feedv1.UpdatePostResponse, error) {
	tags := &basemodel.Tags{}
	tags.ListFromProto(req.Post.Tags)
	imgs := &basemodel.Images{}
	imgs.ListFromProto(req.Post.Images)

	m := &model.Post{
		Uid:     req.Post.Uid,
		Title:   req.Post.Title,
		Content: req.Post.Content,
		Type:    req.Post.Type,
		Cover: &basemodel.Image{
			Url:    req.Post.Cover.Url,
			Width:  req.Post.Cover.Width,
			Height: req.Post.Cover.Height,
			SizeKb: req.Post.Cover.SizeKb,
		},
		Video: &basemodel.Video{
			Url:    req.Post.Video.Url,
			Width:  req.Post.Video.Width,
			Height: req.Post.Video.Height,
			SizeKb: req.Post.Video.SizeKb,
		},
		Tags:   tags,
		Images: imgs,
	}
	// m.Tags = append(m.Tags, req.Post.Tags...)
	// for _, v := range req.Post.Images {
	// 	m.Images = append(m.Images, basemodel.Image{
	// 		Url:    v.Url,
	// 		Name:   v.Name,
	// 		Width:  v.Width,
	// 		Height: v.Height,
	// 		SizeKb: v.SizeKb,
	// 		Hash:   v.Hash,
	// 	})
	// }
	rv, err := s.feedbiz.UpdatePost(c, m)
	if err != nil {
		return nil, err
	}
	return &feedv1.UpdatePostResponse{Id: rv.Id}, nil
}

func (s *FeedServer) GetPost(c context.Context, req *feedv1.GetPostReq) (*feedv1.GetPostResponse, error) {
	rv, err := s.feedbiz.GetPostById(c, req.Id)
	if err != nil {
		return nil, err
	}
	post := rv.ToProto()
	like_count, err := s.likebiz.GetLikeCountByPost(c, req.Id)
	if err != nil {
		return nil, err
	}
	favor_count, err := s.favorbiz.GetFavorCountByPost(c, req.Id)
	if err != nil {
		return nil, err
	}
	share_count, err := s.sharebiz.GetShareCountByPost(c, req.Id)
	if err != nil {
		return nil, err
	}
	fmt.Println("domain # GetPost favor_count:", favor_count)

	post.LikeCount = like_count
	post.FavorCount = favor_count
	post.ShareCount = share_count
	res := &feedv1.GetPostResponse{
		Post: post,
	}

	// res := &feedv1.GetPostResponse{
	// 	Post: &feedv1.Post{
	// 		Id:      rv.Id,
	// 		Uid:     rv.Uid,
	// 		Title:   rv.Title,
	// 		Content: rv.Content,
	// 		Type:    rv.Type,
	// 		Cover: &basev1.Image{
	// 			Url:    rv.Cover.Url,
	// 			Width:  rv.Cover.Width,
	// 			Height: rv.Cover.Height,
	// 			SizeKb: rv.Cover.SizeKb,
	// 		},
	// 		Tags:       rv.Tags.ListToProto(),
	// 		LikeCount:  like_count,
	// 		FavorCount: favor_count,
	// 		ShareCount: share_count,
	// 	},
	// }
	// if rv.Images != nil {
	// 	images := rv.Images.ListToProto()
	// 	res.Post.Images = images
	// }
	// if rv.Video != nil {
	// 	video := rv.Video.ToProto()
	// 	res.Post.Video = video
	// }
	return res, nil
}

func (s *FeedServer) ListPostCard(c context.Context, req *feedv1.ListPostCardReq) (*feedv1.ListPostCardResponse, error) {
	rv, err := s.feedbiz.ListPostCard(c, req.Offset, req.Limit, req.SortBy, req.SortOrder)
	if err != nil {
		return nil, err
	}
	res := &feedv1.ListPostCardResponse{
		PostCards: rv.ListToProto(),
		Total:     int32(len(rv)),
	}
	return res, nil
}

func (s *FeedServer) ListPostCardByIds(c context.Context, req *feedv1.ListPostCardByIdsReq) (*feedv1.ListPostCardByIdsResponse, error) {
	rv, err := s.feedbiz.ListPostCardByIds(c, req.Ids)
	if err != nil {
		return nil, err
	}
	fmt.Println("start of rpc svr#ListPostCardByIds len(rv):", len(rv))
	res := &feedv1.ListPostCardByIdsResponse{
		// PostCards: make([]*feedv1.PostCard, len(rv)), //bug here
		Total: int32(len(rv)),
	}
	postCards := rv.ListToProto()
	res.PostCards = postCards
	for k, v := range rv {
		like_count, err := s.likebiz.GetLikeCountByPost(c, v.Id)
		if err != nil {
			return nil, err
		}
		res.PostCards[k].LikeCount = like_count
		// res.PostCards = append(res.PostCards, &feedv1.PostCard{
		// 	Id:        v.Id,
		// 	Uid:       v.Uid,
		// 	Title:     v.Title,
		// 	Type:      v.Type,
		// 	LikeCount: like_count,
		// 	Cover: &basev1.Image{
		// 		Url:    v.Cover.Url,
		// 		Width:  v.Cover.Width,
		// 		Height: v.Cover.Height,
		// 		SizeKb: v.Cover.SizeKb,
		// 	},
		// })
	}
	fmt.Println("end of rpc svr#len(res.PostCards):", len(res.PostCards))
	return res, nil
}

func (s *FeedServer) ListPostCardByUid(c context.Context, req *feedv1.ListPostCardByUidReq) (*feedv1.ListPostCardByUidResponse, error) {
	rv, err := s.feedbiz.ListPostCardByUid(c, req.Uid, req.Offset, req.Limit, req.SortBy, req.SortOrder)
	if err != nil {
		return nil, err
	}
	res := &feedv1.ListPostCardByUidResponse{
		PostCards: make([]*feedv1.PostCard, len(rv)),
		Total:     int32(len(rv)),
	}
	for _, v := range rv {
		res.PostCards = append(res.PostCards, &feedv1.PostCard{
			Id:    v.Id,
			Uid:   v.Uid,
			Title: v.Title,
			Type:  v.Type,
			Cover: &basev1.Image{
				Url:    v.Cover.Url,
				Width:  v.Cover.Width,
				Height: v.Cover.Height,
				SizeKb: v.Cover.SizeKb,
			},
		})
	}
	return res, nil
}

func (s *FeedServer) ListTimelinePostCardByTimestamp(c context.Context, req *feedv1.ListTimelinePostCardWithTimestampReq) (*feedv1.ListTimelinePostCardWithTimestampResponse, error) {
	rv, err := s.feedbiz.ListPostCardByUidWithTimestamp(c, req.Uid, req.Offset, req.Limit, req.Timestamp, req.SortBy, req.SortOrder)
	if err != nil {
		return nil, err
	}
	res := &feedv1.ListTimelinePostCardWithTimestampResponse{
		PostCards: make([]*feedv1.PostCard, len(rv)),
		Total:     int32(len(rv)),
	}
	for _, v := range rv {
		res.PostCards = append(res.PostCards, &feedv1.PostCard{
			Id:    v.Id,
			Uid:   v.Uid,
			Title: v.Title,
			Type:  v.Type,
			Cover: &basev1.Image{
				Url:    v.Cover.Url,
				Width:  v.Cover.Width,
				Height: v.Cover.Height,
				SizeKb: v.Cover.SizeKb,
			},
		})
	}
	return res, nil
}

func (s *FeedServer) ListVideoPost(c context.Context, req *feedv1.ListVideoPostReq) (*feedv1.ListVideoPostResponse, error) {
	rv, err := s.feedbiz.ListVideoPost(c, req.Offset, req.Limit, req.SortBy, req.SortOrder)
	if err != nil {
		return nil, err
	}
	fmt.Println("start of rpc svr# ListVideoPost len(rv):", len(rv))
	res := &feedv1.ListVideoPostResponse{
		Posts: make([]*feedv1.Post, 0),
		Total: int32(len(rv)),
	}
	prv := &rv
	res.Posts = prv.ListToProto()
	// TODO: performance issue, don't use this way
	for k, v := range rv {
		like_count, err := s.likebiz.GetLikeCountByPost(c, v.Id)
		if err != nil {
			return nil, err
		}
		favor_count, err := s.favorbiz.GetFavorCountByPost(c, v.Id)
		if err != nil {
			return nil, err
		}
		share_count, err := s.sharebiz.GetShareCountByPost(c, v.Id)
		if err != nil {
			return nil, err
		}
		res.Posts[k].LikeCount = like_count
		res.Posts[k].FavorCount = favor_count
		res.Posts[k].ShareCount = share_count
		// res.Posts = append(res.Posts, &feedv1.Post{
		// 	Id:      v.Id,
		// 	Uid:     v.Uid,
		// 	Title:   v.Title,
		// 	Content: v.Content,
		// 	Type:    v.Type,
		// 	Cover: &basev1.Image{
		// 		Url:    v.Cover.Url,
		// 		Width:  v.Cover.Width,
		// 		Height: v.Cover.Height,
		// 		SizeKb: v.Cover.SizeKb,
		// 	},
		// 	Video: &basev1.Video{
		// 		Url:    v.Video.Url,
		// 		Width:  v.Video.Width,
		// 		Height: v.Video.Height,
		// 		SizeKb: v.Video.SizeKb,
		// 	},
		// 	Tags:       v.Tags.ListToProto(),
		// 	Images:     make([]*basev1.Image, len(v.Images)),
		// 	LikeCount:  like_count,
		// 	FavorCount: favor_count,
		// 	ShareCount: share_count,
		// })
		// for _, img := range v.Images {
		// 	res.Posts[len(res.Posts)-1].Images = append(res.Posts[len(res.Posts)-1].Images, &basev1.Image{
		// 		Url:    img.Url,
		// 		Width:  img.Width,
		// 		Height: img.Height,
		// 		SizeKb: img.SizeKb,
		// 		Hash:   img.Hash,
		// 	})
		// }
	}
	return res, nil
}
