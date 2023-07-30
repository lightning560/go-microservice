package data

import (
	"context"
	"miopkg/client/ekafka"
	"miopkg/errors"
	jobcommentv1 "redbook/api/job/comment/v1"

	"google.golang.org/protobuf/proto"
)

func (r *likeRepo) AddLike(c context.Context, replyId, uid int64) (err error) {
	msg := &jobcommentv1.AddLikeMessage{
		ReplyId: replyId,
		Uid:     uid,
	}
	value, err := proto.Marshal(msg)
	if err != nil {
		return
	}
	if err = r.data.commentPubMQ.WriteMessages(c,
		&ekafka.Message{
			Key:   []byte("ADD_LIKE"),
			Value: []byte(value),
		}); err != nil {
		return errors.Wrapf(err, "AddLike commentPubMQ", err)
	}
	return
}
