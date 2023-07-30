package data

import (
	"context"
	"fmt"
	"miopkg/errors"
	"miopkg/log"
	"miopkg/util/snowflake"
	"redbook/app/domain/feed/internal/biz"
	"redbook/app/domain/feed/internal/biz/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type feedRepo struct {
	data *Data
}

func NewFeedRepo(data *Data) biz.IFeedRepo {
	return &feedRepo{data: data}
}

func (r *feedRepo) CreatePost(ctx context.Context, m *model.Post) (int64, error) {
	et := m.ToEntity()
	id := snowflake.GenInt64()
	et.Id = id
	et.CommentId = 0
	et.LikeCount = 0
	et.FavorCount = 0
	et.ShareCount = 0
	et.State = 0
	et.Deleted = 0
	et.ViewCount = 0
	et.CreatedAt = time.Now().Unix()
	et.UpdatedAt = time.Now().Unix()

	_, err := r.data.feedMG.InsertOne(ctx, et)
	if err != nil {
		return -1, errors.Wrapf(err, "CreatePost InsertOne id:%d", id)
	}
	return id, nil
}

func (r *feedRepo) UpdatePost(ctx context.Context, m *model.Post) (res *model.Post, err error) {
	rv := r.data.feedMG.FindOneAndUpdate(ctx, bson.M{"id": m.Id}, bson.M{"$set": m})
	if rv.Err() != nil {
		if rv.Err() == mongo.ErrNoDocuments {
			return nil, errors.Wrapf(errors.ErrNoDocuments, "UpdatePost UpdateOne id:%d", m.Id)
		}
		// update error
		return nil, errors.Wrapf(rv.Err(), "UpdatePost UpdateOne id:%d", m.Id)
	}
	rv.Decode(&res)
	return res, nil
}

func (r *feedRepo) GetPostById(ctx context.Context, id int64) (res *model.Post, err error) {
	if err = r.data.feedMG.FindOne(ctx, bson.M{"id": id}).Decode(&res); err != nil {
		if err == mongo.ErrNoDocuments {
			return res, errors.Wrapf(errors.ErrNoDocuments, "GetPostById id:%d", id)
		}
		return res, errors.Wrapf(err, "GetPostById id:%d", id)
	}
	return
}

func (r *feedRepo) ListPostCard(ctx context.Context, offset, limit int32, by, order string) (model.Posts, error) {
	var sortOrder int32
	var sortBy string
	if by == "like" {
		sortBy = "like_count"
	} else {
		sortBy = "created_at"
	}
	if order == "desc" {
		sortOrder = -1
	} else {
		sortOrder = 1
	}
	limit64 := int64(limit)
	offset64 := int64(offset)
	opts := &options.FindOptions{
		Sort: bson.D{
			bson.E{Key: sortBy, Value: sortOrder},
		},
		Limit: &limit64,
		Skip:  &offset64,
	}
	cur, err := r.data.feedMG.Find(ctx, bson.M{}, opts)
	if err != nil {
		errors.Wrap(err, "ListCardByUser desc Find")
		return nil, err
	}
	defer cur.Close(ctx)
	var res []*model.Post
	for cur.Next(ctx) {
		var m *model.Post
		err = cur.Decode(&m)
		if err != nil {
			r.data.logger.Errorf("ListPostCard cursor Decode", log.Any("error", err))
			return nil, err
		}
		res = append(res, m)
	}
	return res, nil
}

func (r *feedRepo) ListPostCardByIds(ctx context.Context, ids []int64) (res model.Posts, err error) {
	fmt.Println("RawListPostCardByIds ids:", ids)
	if len(ids) == 0 {
		err = errors.Wrap(errors.ErrInvalidParam, "RawListPostCardByIds")
		return
	}
	cur, err := r.data.feedMG.Find(ctx, bson.M{"id": bson.M{"$in": ids}})
	if err != nil {
		r.data.logger.Errorf("ListPostCardByIds Find", log.Any("error", err))
		return
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var m *model.Post
		err = cur.Decode(&m)
		if err != nil {
			r.data.logger.Errorf("ListPostCardByIds Decode", log.Any("error", err))
			return
		}
		res = append(res, m)
	}
	fmt.Println("RawListPostCardByIds len(res):", len(res))
	for _, v := range res {
		fmt.Println("RawListPostCardByIds v:", v)
	}
	return
}

func (r *feedRepo) ListPostCardByUid(ctx context.Context, uid int64, offset, limit int32, by, order string) (res model.Posts, err error) {
	filter := &bson.M{"uid": uid}
	var (
		sortOrder int32
		sortBy    string
	)
	if by == "like" {
		sortBy = "like_count"
	} else {
		sortBy = "created_at"
	}
	if order == "desc" {
		sortOrder = -1
	} else {
		sortOrder = 1
	}
	lm := int64(limit)
	os := int64(offset)
	opts := &options.FindOptions{
		Sort: bson.D{
			bson.E{Key: sortBy, Value: sortOrder},
		},
		Limit: &lm,
		Skip:  &os,
	}

	cur, err := r.data.feedMG.Find(ctx, filter, opts)
	if err != nil {
		r.data.logger.Errorf("ListCardByUser desc Find", log.Any("error", err))
		return
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var m *model.Post
		err = cur.Decode(&m)
		if err != nil {
			r.data.logger.Errorf("ListCardByUser cursor Decode", log.Any("error", err))
			return
		}
		res = append(res, m)
	}
	return
}

// FIXME: not work
func (r *feedRepo) ListPostCardByUidWithTimestamp(ctx context.Context, uid int64, offset, limit int32, timestamp int64, by, order string) (res model.Posts, err error) {
	filter := &bson.M{"uid": uid, "created_at": bson.M{"$gt": timestamp}}
	limit64 := int64(limit)
	offset64 := int64(offset)
	opts := &options.FindOptions{
		Sort: bson.D{
			bson.E{Key: "created_at", Value: -1},
		},
		Limit: &limit64,
		Skip:  &offset64,
	}

	cur, err := r.data.feedMG.Find(ctx, filter, opts)
	if err != nil {
		r.data.logger.Errorf("ListTimelineCardByTimestamp desc Find", log.Any("error", err))
		return
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var m *model.Post
		err = cur.Decode(&m)
		if err != nil {
			r.data.logger.Errorf("ListTimelineCardByTimestamp cursor Decode", log.Any("error", err))
			return
		}
		res = append(res, m)
	}
	return
}

// FIXME: not work
func (r *feedRepo) DeletePost(ctx context.Context, id int64) error {
	return nil
}

func (r *feedRepo) ListVideoPost(ctx context.Context, offset, limit int32, by, order string) (res model.Posts, err error) {
	filter := &bson.M{"type": "video"}
	var (
		sortOrder int32
		sortBy    string
	)
	if by == "like" {
		sortBy = "like_count"
	} else {
		sortBy = "created_at"
	}
	if order == "desc" {
		sortOrder = -1
	} else {
		sortOrder = 1
	}
	limit64 := int64(limit)
	offset64 := int64(offset)
	opts := &options.FindOptions{
		Sort: bson.D{
			bson.E{Key: sortBy, Value: sortOrder},
		},
		Limit: &limit64,
		Skip:  &offset64,
	}
	cur, err := r.data.feedMG.Find(ctx, filter, opts)
	if err != nil {
		return nil, errors.Wrap(err, "ListVideoPost Find")
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var m *model.Post
		err = cur.Decode(&m)
		if err != nil {
			return nil, errors.Wrap(err, "ListVideoPost Decode")
		}
		res = append(res, m)
	}
	// if len(res) == 0 {
	// 	return nil, errors.Wrap(errors.ErrNotFound, "RawListVideoPost")
	// }
	return
}
