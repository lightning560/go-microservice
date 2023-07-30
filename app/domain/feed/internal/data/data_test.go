package data

import (
	"context"
	"fmt"
	"miopkg/conf"
	"redbook/app/domain/feed/internal/biz/model"
	"redbook/common/basemodel"
	"strings"
	"testing"
	"time"

	"github.com/BurntSushi/toml"
)

func _newData(ctx context.Context) (*Data, func(), error) {
	d := &Data{
		feedMG:     NewMongoClient(ctx, "feed"),
		relationMG: NewMongoClient(ctx, "relation"),
		likeMG:     NewMongoClient(ctx, "like"),
		shareMG:    NewMongoClient(ctx, "share"),
		favorMG:    NewMongoClient(ctx, "favor"),
		// redisClient: NewRedisClient(ctx),
		// fanout:      NewFanout(),
	}
	return d, func() { // close db conn
		// if err := d.likeWDB.Close(); err != nil {
		// 	fmt.Println(err)
		// }
	}, nil
}

func loadConfig(t *testing.T) {
	var config = `
	[mongo.dev.feed]
		debug=true
		dsn="mongodb://root:password@127.0.0.1:27017/?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&ssl=false"
	`
	err := conf.LoadFromReader(strings.NewReader(config), toml.Unmarshal)
	if err != nil {
		t.Fatal(err)
	}
}

func addPost(t *testing.T, m *model.Post) {
	loadConfig(t)
	data, _, err := _newData(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	repo := NewFeedRepo(data)

	id, err := repo.CreatePost(context.Background(), m)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("addPost id:", id)
}

func TestCreateVideoPost(t *testing.T) {
	m := &model.Post{
		Id:      1,
		Uid:     101,
		Title:   "Title+test",
		Content: "Content+test",
		Cover: &basemodel.Image{
			Id:     1,
			Name:   "testimg",
			Url:    "https://sns-img-bd.xhscdn.com/57f5f230-ce2c-d965-c7f7-35d6453f30c6?imageView2/2/w/640/format/webp",
			Hash:   "57f5f230-ce2c-d965-c7f7-",
			Width:  960,
			Height: 840,
			SizeKb: 258,
		},
		Type: "video",
		Tags: &basemodel.Tags{
			{
				TagId:   1,
				Name:    "sakura",
				BizType: "feed",
			},
			{
				TagId:   2,
				Name:    "travel",
				BizType: "feed",
			},
		},
		Video: &basemodel.Video{
			Id:        1,
			Url:       "http://sns-video-bd.xhscdn.com/stream/110/259/01e41d73c6b65ab30103700387130e97aa_259.mp4",
			Name:      "testvideo",
			Type:      "mp4",
			Hash:      "01e41d73c6b65ab30103700387130e97aa",
			Width:     960,
			Height:    840,
			SizeKb:    258,
			Length:    10,
			CreatedAt: time.Now().Unix(),
			Cover: &basemodel.Image{
				Id:     1,
				Name:   "testimg",
				Url:    "https://sns-img-bd.xhscdn.com/57f5f230-ce2c-d965-c7f7-35d6453f30c6?imageView2/2/w/640/format/webp",
				Hash:   "57f5f230-ce2c-d965-c7f7-",
				Width:  960,
				Height: 840,
				SizeKb: 258,
			},
		},
		Images:     &basemodel.Images{},
		Deleted:    0,
		State:      1,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
		LikeCount:  0,
		ShareCount: 0,
		FavorCount: 0,
		ViewCount:  1,
		CommentId:  1,
	}
	addPost(t, m)
}

func TestCreateImagePost(t *testing.T) {
	m := &model.Post{
		Id:      1,
		Uid:     101,
		Title:   "Title+test",
		Content: "Content+test",
		Cover: &basemodel.Image{
			Id:     1,
			Name:   "testimg",
			Url:    "https://sns-img-qc.xhscdn.com/e0f53873-fccf-d1ac-d43a-7ca19bccb1c0?imageView2/2/h/1920/format/webp",
			Hash:   "57f5f230-ce2c-d965-c7f7-",
			Width:  960,
			Height: 840,
			SizeKb: 258,
		},
		Type: "image",
		Tags: &basemodel.Tags{
			{
				TagId:   1,
				Name:    "sakura",
				BizType: "feed",
			},
			{
				TagId:   2,
				Name:    "travel",
				BizType: "feed",
			},
		},
		Video: &basemodel.Video{
			Id:        1,
			Url:       "http://sns-video-bd.xhscdn.com/stream/110/259/01e41d73c6b65ab30103700387130e97aa_259.mp4",
			Name:      "testvideo",
			Type:      "mp4",
			Hash:      "01e41d73c6b65ab30103700387130e97aa",
			Width:     960,
			Height:    840,
			SizeKb:    258,
			Length:    10,
			CreatedAt: time.Now().Unix(),
			Cover: &basemodel.Image{
				Id:     1,
				Name:   "testimg",
				Url:    "https://sns-img-bd.xhscdn.com/57f5f230-ce2c-d965-c7f7-35d6453f30c6?imageView2/2/w/640/format/webp",
				Hash:   "57f5f230-ce2c-d965-c7f7-",
				Width:  960,
				Height: 840,
				SizeKb: 258,
			},
		},
		Images: &basemodel.Images{
			{
				Id:     1,
				Name:   "testimg",
				Url:    "https://sns-img-bd.xhscdn.com/62e05c59-e936-d38a-d4f4-11e2c3d5c05e?imageView2/2/h/1920/format/webp",
				Hash:   "57f5f230-ce2c-d965-c7f7-",
				Width:  960,
				Height: 840,
				SizeKb: 258,
			},
			{
				Id:     2,
				Name:   "testimg",
				Url:    "https://sns-img-bd.xhscdn.com/c63e8484-21a7-5968-bb60-6d3abe37bcee?imageView2/2/h/1920/format/webp",
				Hash:   "57f5f230-ce2c-d965-c7f7",
				Width:  960,
				Height: 840,
				SizeKb: 258,
			},
		},
		Deleted:    0,
		State:      1,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
		LikeCount:  0,
		ShareCount: 0,
		FavorCount: 0,
		ViewCount:  1,
		CommentId:  1,
	}
	addPost(t, m)
}
