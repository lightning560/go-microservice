package data

import (
	"context"

	"miopkg/client/ekafka"
	"miopkg/errors"
	jobcommentv1 "redbook/api/job/comment/v1"

	"google.golang.org/protobuf/proto"
)

func (r *replyRepo) CreateReply(c context.Context, isFloor bool, subjectId int64, bizType string, owner_uid, floorId int64, content string, atUid int64) error {
	msg := &jobcommentv1.AddReplyMessage{
		IsFloor:   isFloor,
		SubjectId: subjectId,
		BizType:   bizType,
		OwnerUid:  owner_uid,
		FloorId:   floorId,
		Content:   content,
		AtUid:     atUid,
	}
	value, err := proto.Marshal(msg)
	if err != nil {
		return errors.Wrapf(err, "AddReplyMessage marshal msg:%+v", msg)
	}
	if err = r.data.commentPubMQ.WriteMessages(c,
		&ekafka.Message{
			Key:   []byte("ADD_REPLY"),
			Value: []byte(value),
		}); err != nil {
		r.data.logger.Errorf("AddReply commentPubMQ", err)
	}
	return err
}
