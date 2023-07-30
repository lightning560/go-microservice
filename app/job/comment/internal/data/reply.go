package data

import (
	"context"
	"fmt"
	"miopkg/errors"
	"miopkg/util/snowflake"
	"redbook/app/job/comment/internal/biz"
	"redbook/app/job/comment/internal/data/entity"
	"time"
)

type replyRepo struct {
	data *Data
}

func NewReplyRepo(data *Data) biz.IReplyRepo {
	return &replyRepo{data: data}
}

func (r *replyRepo) AddReply(ctx context.Context, isFloor bool, subjectId, floorId, ownerUid, atUid int64, bizType, content string) error {
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
		State:        1,
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
		FanGrade:     0,
		Platform:     0,
		Device:       "dev",
		Attr:         0,
		Dialog:       0,
	}
	if isFloor && floorId == 0 {
		fmt.Printf("data#AddReply:isFloor:%v, floorId:%d\n", isFloor, floorId)
		reply.FloorAttr.ReplyCount = 0
		reply.FloorAttr.PinAdmin = 0
		reply.FloorAttr.PinOwner = 0
		reply.FloorAttr.Fold = 0
		reply.FloorAttr.Hot = false
	}
	fmt.Printf("data#AddReply:%+v\n", reply)
	if _, err := r.data.replyMG.InsertOne(ctx, reply); err != nil {
		return errors.Wrapf(err, "CreateReply InsertOne SubjectId:%d, BizType:%s,OwnerUid: %d", subjectId, bizType, ownerUid)
	}
	return nil
}
