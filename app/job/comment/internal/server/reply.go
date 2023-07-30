package server

import (
	"context"
	"fmt"
	"miopkg/errors"

	commentv1 "redbook/api/job/comment/v1"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

func (s *CommentServer) AddReply(ctx context.Context, message *kafka.Message) (err error) {
	req := &commentv1.AddReplyMessage{}
	if err = proto.Unmarshal(message.Value, req); err != nil {
		return errors.Wrapf(err, "proto.Unmarshal(%v)", message.Value)
	}
	fmt.Printf("server#AddReply:%+v\n", req)
	if err = s.replyBiz.AddReply(ctx, req.IsFloor, req.SubjectId, req.FloorId, req.OwnerUid, req.AtUid, req.BizType, req.Content); err != nil {
		return errors.Wrapf(err, "comment.Job.AddReply(%v) error(%v)", req)
	}
	if err = s.commentMQConsumer.CommitMessages(ctx, message); err != nil {
		return errors.Wrapf(err, "comment.Job.AddReply(%v) error(%v) , CommitMessages", req)
	}
	return
}
