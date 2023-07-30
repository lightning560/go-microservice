package server

import (
	"context"
	"fmt"
	"miopkg/client/ekafka"
	"miopkg/client/ekafka/consumerserver"
	xlog "miopkg/log"
	commentv1 "redbook/api/job/comment/v1"

	"redbook/app/job/comment/internal/biz"
	"redbook/app/job/comment/internal/data"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

func (eng *Engine) newKafkaWorker() error {
	ctx := context.Background()
	// 初始化logger
	logger := xlog.MioLogger.With(xlog.FieldMod("comment.job"))
	// 初始化data
	mongoDB := data.NewMongoClient(ctx, "comment")
	replyMG := data.NewMongoCollection(ctx, mongoDB, "reply")
	likeMG := data.NewMongoCollection(ctx, mongoDB, "like")
	// redisClient := data.NewRedisClient(ctx)
	// fanout := data.NewFanout()
	dataInstance, _, err := data.NewData(ctx, replyMG, likeMG, logger)
	if err != nil {
		xlog.Panicf("data.NewData error(%v)", err)
		return err
	}
	// 初始化biz
	replyRepo := data.NewReplyRepo(dataInstance)
	likeRepo := data.NewLikeRepo(dataInstance)
	// 初始化server
	replyBiz := biz.NewReplyBiz(replyRepo)
	likeBiz := biz.NewLikeBiz(likeRepo)

	// 初始化ekafka组件
	// 依赖 `ekafka` 管理 Kafka consumer
	ekafkaCmp := ekafka.Load("kafka").Build()
	cs := consumerserver.Load("kafkaConsumerServers.cs1").Build(
		consumerserver.WithEkafka(ekafkaCmp),
	)
	svr := NewCommentServer(replyBiz, likeBiz, ekafkaCmp.Consumer("c1"))

	// 用来接收、处理 `kafka-go` 和处理消息的回调产生的错误
	consumptionErrors := make(chan error)
	// 注册处理消息的回调函数
	cs.OnEachMessage(consumptionErrors, func(ctx context.Context, message kafka.Message) error {
		fmt.Printf("got a message: %s\n", string(message.Value))
		// 如果返回错误则会被转发给 `consumptionErrors`
		switch string(message.Key) {
		case "ADD_LIKE":
			fmt.Println("switch ADD_LIKE")
			svr.AddLike(ctx, message)
		case "ADD_REPLY":
			fmt.Println("switch ADD_REPLY")
			svr.AddReply(ctx, &message)
		default:
			fmt.Println("unknown message.Key:", string(message.Key))
			err = fmt.Errorf("unknown message.Key:%s", string(message.Key))
		}
		// srv.consumeCommentMessage(ctx, message)
		return nil
	})

	// srv.consumeComment(ctx, cmp.Consumer("c1"))
	return eng.Serve(cs)
}

func (s *CommentServer) consumeCommentMessage(ctx context.Context, message kafka.Message) error {
	fmt.Println("got a message.Headers:", message.Headers)
	fmt.Println("got a message.Key: ", string(message.Key))
	fmt.Println("got a message.Value:", string(message.Value))

	commentMessage := &commentv1.CommentMessage{}
	if err := proto.Unmarshal(message.Value, commentMessage); err != nil { // protobuf marshal&unmarshal
		xlog.Errorf("proto.Unmarshal(%v) error(%v)", message, err)
	}
	fmt.Println("commentMessage.Uid:", commentMessage.Uid)
	return nil

}

func (eng *Engine) newKafkaC() error {
	// 依赖 `ekafka` 管理 Kafka consumer
	ec := ekafka.Load("kafka").Build()
	cs := consumerserver.Load("kafkaConsumerServers.cs1").Build(
		consumerserver.WithEkafka(ec),
	)

	// 用来接收、处理 `kafka-go` 和处理消息的回调产生的错误
	consumptionErrors := make(chan error)

	// 注册处理消息的回调函数
	cs.OnEachMessage(consumptionErrors, func(ctx context.Context, message kafka.Message) error {
		fmt.Printf("got a message: %s\n", string(message.Value))
		// 如果返回错误则会被转发给 `consumptionErrors`
		return nil
	})

	return eng.Serve(cs)
}
