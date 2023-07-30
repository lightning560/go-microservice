package data

import (
	"context"
	"fmt"
	"miopkg/errors"
	"miopkg/log"
	"miopkg/util/constant"
	"miopkg/util/snowflake"
	"redbook/app/domain/comment/internal/biz"
	"redbook/app/domain/comment/model/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type replyRepo struct {
	data *Data
}

func NewReplyRepo(data *Data) biz.IReplyRepo {
	return &replyRepo{data: data}
}

func (r *replyRepo) AddReply(ctx context.Context, isFloor bool, subjectId int64, bizType string, floorId, ownerUid int64, content string, atUid int64) (int64, error) {
	id := snowflake.GenInt64()
	reply := &entity.Reply{
		Id:        id,
		SubjectId: subjectId,
		BizType:   bizType,
		OwnerUid:  ownerUid,
		FloorId:   floorId,
		Content:   content,
		AtUid:     atUid,

		LikeCount:    0,
		DislikeCount: 0,
		Deleted:      0,
		State:        0,
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
		FanGrade:     0,
		Platform:     0,
		Device:       "",
		Attr:         0,
		Dialog:       0,
	}
	if isFloor && floorId == 0 {
		reply.FloorAttr.ReplyCount = 0
		reply.FloorAttr.PinAdmin = 0
		reply.FloorAttr.PinOwner = 0
		reply.FloorAttr.Fold = 0
		reply.FloorAttr.Hot = false
	}
	if _, err := r.data.replyMG.InsertOne(ctx, reply); err != nil {
		return constant.ResultError, errors.Wrapf(err, "AddReply InsertOne SubjectId:%d, BizType:%s,OwnerUid: %d", subjectId, bizType, ownerUid)
	}
	return id, nil
}

func (r *replyRepo) DeleteReply(ctx context.Context, id int64) error {
	if sr := r.data.replyMG.FindOneAndDelete(ctx, bson.M{"id": id}); sr.Err() != nil {
		if sr.Err() == mongo.ErrNoDocuments {
			return nil
		}
		return errors.Wrapf(sr.Err(), "DeleteReply FindOneAndDelete id:%d", id)
	}
	return nil
}

func (r *replyRepo) ListFloorRootBySubjectSortReply(ctx context.Context, subjectId int64, offset, limit int32, by, order string) (entity.Replies, error) {
	fmt.Printf("ListFloorRootBySubjectSortReply subjectId:%d, limit:%d, offset:%d, by:%s, order:%s\n", subjectId, limit, offset, by, order)
	filter := &bson.M{"floor_id": int64(0), "subject_id": subjectId}
	fmt.Printf("ListFloorRootBySubjectSortReply filter:%+v\n", filter)
	// sortOrder := 1
	// if order == "desc" {
	// 	sortOrder = -1
	// }
	// limit64 := int64(limit)
	// skip64 := int64(offset)

	// opts := &options.FindOptions{
	// 	Sort: bson.D{
	// 		bson.E{Key: by, Value: sortOrder},
	// 	},
	// 	Skip:  &skip64,
	// 	Limit: &limit64,
	// }

	offset64 := int64(offset)
	limit64 := int64(limit)
	fmt.Printf("ListFloorRootBySubjectSortReply offset64:%d, limit64:%d\n", offset64, limit64)
	opts := options.Find().
		SetSort(bson.D{{Key: "floor_attr.reply_count", Value: 1}}).SetSkip(offset64).SetLimit(limit64)

	cur, err := r.data.replyMG.Find(ctx, filter, opts)
	if err != nil {
		return nil, errors.Wrapf(err, "ListSubjectFloorSortReply Find subjectId:%d, limit:%d, offset:%d, by:%s, order:%s", subjectId, limit, offset, by, order)
	}
	var replies []*entity.Reply
	for cur.Next(ctx) {
		var reply entity.Reply
		err := cur.Decode(&reply)
		if err != nil {
			return nil, errors.Wrapf(err, "ListSubjectFloorSortReply Decode subjectId:%d, limit:%d, offset:%d, by:%s, order:%s", subjectId, limit, offset, by, order)
		}
		replies = append(replies, &reply)
	}
	fmt.Println("data # ListFloorRootBySubjectSortReply len(replies):", len(replies))
	return replies, nil
}

func (r *replyRepo) ListFloorRootBySubjectSortTime(ctx context.Context, subjectId int64, offset, limit int32, by, order string) (entity.Replies, error) {
	filter := &bson.M{"floor_id": int64(0), "subject_id": subjectId}
	sortOrder := 1
	if order == "desc" {
		sortOrder = -1
	}
	limit64, offset64 := int64(limit), int64(offset)
	opts := &options.FindOptions{
		Sort: bson.D{
			bson.E{Key: "created_at", Value: sortOrder},
		},
		Limit: &limit64,
		Skip:  &offset64,
	}
	cur, err := r.data.replyMG.Find(ctx, filter, opts)
	if err != nil {
		return nil, errors.Wrapf(err, "ListSubjectFloorSortTime Find subjectId:%d, limit:%d, offset:%d, by:%s, order:%s", subjectId, limit, offset, by, order)
	}
	var replies []*entity.Reply
	for cur.Next(ctx) {
		var reply entity.Reply
		err := cur.Decode(&reply)
		if err != nil {
			r.data.logger.Errorf("ListSubjectFloorSortTime Decode", log.Any("error", err))
			return nil, err
		}
		replies = append(replies, &reply)
	}
	return replies, nil
}

func (r *replyRepo) ListReplyByIds(ctx context.Context, ids []int64) (entity.Replies, error) {
	var replies []*entity.Reply
	cur, err := r.data.replyMG.Find(ctx, bson.M{"id": bson.M{"$in": ids}})
	if err != nil {
		r.data.logger.Errorf("ListReplyByIds Find", log.Any("error", err))
		return nil, err
	}
	for cur.Next(ctx) {
		var reply entity.Reply
		err := cur.Decode(&reply)
		if err != nil {
			r.data.logger.Errorf("ListReplyByIds Decode", log.Any("error", err))
			return nil, err
		}
		replies = append(replies, &reply)
	}
	return replies, nil
}
func (r *replyRepo) ListReplyByFloorSortLike(ctx context.Context, subjectId, floorId int64, offset, limit int32, by, order string) (entity.Replies, error) {
	fmt.Printf("ListReplyByFloorSortLike subjectId:%d, floorId:%d, limit:%d, offset:%d, by:%s, order:%s\n", subjectId, floorId, limit, offset, by, order)
	filter := &bson.M{"subject_id": subjectId, "floor_id": floorId}
	sortOrder := 1
	if order == "desc" {
		sortOrder = -1
	}
	limit64, offset64 := int64(limit), int64(offset)
	fmt.Printf("ListReplyByFloorSortLike limit64:%d, offset64:%d\n", limit64, offset64)
	opts := &options.FindOptions{
		Sort: bson.D{
			bson.E{Key: "like_count", Value: sortOrder},
		},
		Limit: &limit64,
		Skip:  &offset64,
	}
	cur, err := r.data.replyMG.Find(ctx, filter, opts)
	if err != nil {
		return nil, errors.Wrapf(err, "ListFloorReplySortLike Find floorId:%d, limit:%d, offset:%d, by:%s, order:%s", floorId, limit, offset, by, order)
	}
	var replies entity.Replies
	for cur.Next(ctx) {
		var reply entity.Reply
		err := cur.Decode(&reply)
		if err != nil {
			return nil, errors.Wrapf(err, "ListFloorReplySortLike Decode floorId:%d, limit:%d, offset:%d, by:%s, order:%s", floorId, limit, offset, by, order)
		}
		replies = append(replies, &reply)
	}
	fmt.Println("ListReplyByFloorSortLike len(replies) :", len(replies))
	fmt.Printf("ListReplyByFloorSortLike replies:%+v\n", replies[0])
	return replies, nil
}

func (r *replyRepo) ListReplyByFloorSortTime(ctx context.Context, subjectId, floorId int64, offset, limit int32, by, order string) (entity.Replies, error) {
	filter := &bson.M{"subject_id": subjectId, "floor_id": floorId}
	sortOrder := 1
	if order == "desc" {
		sortOrder = -1
	}
	limit64, offset64 := int64(limit), int64(offset)
	opts := &options.FindOptions{
		Sort: bson.D{
			bson.E{Key: "created_at", Value: sortOrder},
		},
		Limit: &limit64,
		Skip:  &offset64,
	}
	cur, err := r.data.replyMG.Find(ctx, filter, opts)
	if err != nil {
		return nil, errors.Wrapf(err, "ListFloorReplySortTime Find floorId:%d, limit:%d, offset:%d, by:%s, order:%s", floorId, limit, offset, by, order)
	}
	var replies []*entity.Reply
	for cur.Next(ctx) {
		var reply entity.Reply
		err := cur.Decode(&reply)
		if err != nil {
			return nil, errors.Wrapf(err, "ListFloorReplySortTime Decode floorId:%d, limit:%d, offset:%d, by:%s, order:%s", floorId, limit, offset, by, order)
		}
		replies = append(replies, &reply)
	}
	return replies, nil
}
