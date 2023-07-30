package data

import (
	"context"
	"miopkg/client/ekafka"
	"miopkg/log"
	commentv1 "redbook/api/domain/comment/v1"
	userv1 "redbook/api/domain/user/v1"
)

type Data struct {
	userRpcClient    userv1.UserRpcClient
	commentRpcClient commentv1.CommentRpcClient
	logger           *log.Logger
	commentPubMQ     *ekafka.Producer
}

func NewCommentPubMQ() *ekafka.Producer {
	ec := ekafka.Load("kafka").Build()
	// 使用p1生产者生产消息
	return ec.Producer("p1")
}

func NewData(c context.Context, commentRpcClient commentv1.CommentRpcClient, userRpcClient userv1.UserRpcClient, commentPubMQ *ekafka.Producer) (*Data, func(), error) {
	// commentPubMQ := NewCommentPubMQ()
	logger := log.MioLogger.With(log.FieldMod("interface.comment.data"))
	d := &Data{
		userRpcClient:    userRpcClient,
		commentRpcClient: commentRpcClient,
		commentPubMQ:     commentPubMQ,
		logger:           logger,
	}
	return d, func() {
		if err := commentPubMQ.Close(); err != nil {
			logger.Errorf("close commentPubMQ:", err)
		}
	}, nil
}
