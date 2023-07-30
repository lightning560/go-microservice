package data

import (
	"context"
	"fmt"
	domainfeedv1 "redbook/api/domain/feed/v1"
	basev1 "redbook/api/redbookpb/v1"
	"redbook/app/interface/feed/internal/biz"
	"redbook/app/interface/feed/internal/biz/model"

	"github.com/pkg/errors"
)

type feedRepo struct {
	data *Data
}

func NewFeedRepo(data *Data) biz.IFeedRepo {
	return &feedRepo{data: data}
}

func (r *feedRepo) CreatePost(ctx context.Context, p *model.Post) (id int64, err error) {
	req := &domainfeedv1.CreatePostReq{
		Post: &domainfeedv1.Post{
			Title:   p.Title,
			Content: p.Content,
			Uid:     p.Uid,
			Tags:    make([]*basev1.Tag, 0, len(*p.Tags)),
			Cover: &basev1.Image{
				Url:    p.Cover.Url,
				Width:  p.Cover.Width,
				Height: p.Cover.Height,
				SizeKb: p.Cover.SizeKb,
				Name:   p.Cover.Name,
				Hash:   p.Cover.Hash,
			},
			Type:   p.Type,
			Video:  p.Video.ToProto(),
			Images: make([]*basev1.Image, 0, len(*p.Images)),
		},
	}
	for _, v := range *p.Images {
		req.Post.Images = append(req.Post.Images, &basev1.Image{
			Url:    v.Url,
			Width:  v.Width,
			Height: v.Height,
			SizeKb: v.SizeKb,
			Name:   v.Name,
			Hash:   v.Hash,
		})
	}
	req.Post.Tags = p.Tags.ListToProto()
	rv, err := r.data.feedRpc.CreatePost(ctx, req)
	if err != nil {
		return -1, errors.Wrap(err, "rpc create post error")
	}
	return rv.Id, nil
}

func (r *feedRepo) GetPost(ctx context.Context, id int64) (*model.Post, error) {
	resp, err := r.data.feedRpc.GetPost(ctx, &domainfeedv1.GetPostReq{Id: id})
	if err != nil {
		return nil, err
	}
	res := &model.Post{}
	res.FromDomainProto(resp.Post)
	// tags := &basemodel.Tags{}
	// tags.ListFromProto(resp.Post.Tags)
	// res := &model.Post{
	// 	Id:      resp.Post.Id,
	// 	Title:   resp.Post.Title,
	// 	Content: resp.Post.Content,
	// 	Cover: &basemodel.Image{
	// 		Url:    resp.Post.Cover.Url,
	// 		Width:  resp.Post.Cover.Width,
	// 		Height: resp.Post.Cover.Height,
	// 		SizeKb: resp.Post.Cover.SizeKb,
	// 		Name:   resp.Post.Cover.Name,
	// 		Hash:   resp.Post.Cover.Hash,
	// 	},
	// 	Type: resp.Post.Type,
	// 	Video: &basemodel.Video{
	// 		Url:    resp.Post.Video.Url,
	// 		Width:  resp.Post.Video.Width,
	// 		Height: resp.Post.Video.Height,
	// 		SizeKb: resp.Post.Video.SizeKb,
	// 	},
	// 	Uid:        resp.Post.Uid,
	// 	LikeCount:  resp.Post.LikeCount,
	// 	ShareCount: resp.Post.ShareCount,
	// 	FavorCount: resp.Post.FavorCount,
	// 	Tags:       tags,
	// }
	// imgs := &basemodel.Images{}
	// imgs.ListFromProto(resp.Post.Images)
	// res.Images = imgs
	// for _, v := range resp.Post.Images {
	// 	img := basemodel.Image{}
	// 	res.Images = append(res.Images, basemodel.Image{
	// 		Url:    v.Url,
	// 		Width:  v.Width,
	// 		Height: v.Height,
	// 		SizeKb: v.SizeKb,
	// 		Name:   v.Name,
	// 		Hash:   v.Hash,
	// 	})
	// }
	return res, nil
}

func (r *feedRepo) ListPostCard(ctx context.Context, offset, limit int32, sortBy, sortOrder string) (res *model.PostCards, err error) {
	resp, err := r.data.feedRpc.ListPostCard(ctx, &domainfeedv1.ListPostCardReq{
		Offset:    offset,
		Limit:     limit,
		SortBy:    sortBy,
		SortOrder: sortOrder,
	})
	if err != nil {
		return nil, errors.Wrap(err, "data.ListPostCard")
	}
	res = &model.PostCards{}
	res.ListFromDomainProto(resp.PostCards)
	return res, nil
}

func (r *feedRepo) ListPostCardByIds(ctx context.Context, ids []int64) (model.PostCards, error) {
	resp, err := r.data.feedRpc.ListPostCardByIds(ctx, &domainfeedv1.ListPostCardByIdsReq{Ids: ids})
	if err != nil {
		return nil, err
	}
	fmt.Println("len(resp.PostCards)", len(resp.PostCards))
	// res := make([]*model.PostCard, len(resp.PostCards))
	// for k, v := range resp.PostCards {
	// 	tags := &basemodel.Tags{}
	// 	tags.ListFromProto(v.Tags)
	// 	res[k] = &model.PostCard{
	// 		Id:    v.Id,
	// 		Uid:   v.Uid,
	// 		Title: v.Title,
	// 		Cover: &basemodel.Image{
	// 			Url:  v.Cover.Url,
	// 			Hash: v.Cover.Hash,
	// 		},
	// 		Type:      v.Type,
	// 		LikeCount: v.LikeCount,
	// 		Tags:      tags,
	// 		//  TODO: add作者信息
	// 	}
	// }

	res := &model.PostCards{}
	res.ListFromDomainProto(resp.PostCards)
	fmt.Println("interface-data:len(res)", len(*res))
	return *res, nil
}

func (r *feedRepo) UpdatePost(ctx context.Context, p *model.Post) error {
	req := &domainfeedv1.UpdatePostReq{
		Post: &domainfeedv1.Post{
			Id:      p.Id,
			Title:   p.Title,
			Content: p.Content,
			Uid:     p.Uid,
			Cover: &basev1.Image{
				Url:    p.Cover.Url,
				Width:  p.Cover.Width,
				Height: p.Cover.Height,
				SizeKb: p.Cover.SizeKb,
				Name:   p.Cover.Name,
				Hash:   p.Cover.Hash,
			},
			Type: p.Type,
			Video: &basev1.Video{
				Url:    p.Video.Url,
				Width:  p.Video.Width,
				Height: p.Video.Height,
				SizeKb: p.Video.SizeKb,
				Name:   p.Video.Name,
				Hash:   p.Video.Hash,
			},
			Images: make([]*basev1.Image, 0, len(*p.Images)),
			Tags:   p.Tags.ListToProto(),
		},
	}
	_, err := r.data.feedRpc.UpdatePost(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (r *feedRepo) DeletePost(ctx context.Context, id int64) error {
	_, err := r.data.feedRpc.DeletePost(ctx, &domainfeedv1.DeletePostReq{Id: id})
	if err != nil {
		return err
	}
	return nil
}

func (r *feedRepo) ListVideoPost(ctx context.Context, offset, limit int32, sortBy, sortOrder string) (res *model.Posts, err error) {
	resp, err := r.data.feedRpc.ListVideoPost(ctx, &domainfeedv1.ListVideoPostReq{
		Offset:    offset,
		Limit:     limit,
		SortBy:    sortBy,
		SortOrder: sortOrder,
	})
	if err != nil {
		return nil, errors.Wrap(err, "data.ListVideoPost")
	}
	res = &model.Posts{}
	res.ListFromDomainProto(resp.Posts)
	return
}
