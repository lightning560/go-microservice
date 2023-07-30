package server

import (
	"context"
	"miopkg/errors"

	commentv1 "redbook/api/job/comment/v1"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

func (s *CommentServer) AddLike(ctx context.Context, message kafka.Message) (err error) {
	req := &commentv1.AddLikeMessage{}
	if err = proto.Unmarshal(message.Value, req); err != nil {
		return errors.Wrapf(err, "proto.Unmarshal(%v)", message.Value)
	}
	if err = s.likeBiz.AddLike(ctx, req.ReplyId, req.Uid); err != nil {
		return errors.Wrapf(err, "comment.Job.AddLike(%v) error(%v)", req)
	}
	if err = s.commentMQConsumer.CommitMessages(ctx, &message); err != nil {
		return errors.Wrapf(err, "comment.Job.AddLike(%v) error(%v) , CommitMessages", req)
	}
	return
}
