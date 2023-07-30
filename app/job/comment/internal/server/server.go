package server

import (
	"fmt"
	"miopkg/application"
	"miopkg/client/ekafka"
	"redbook/app/job/comment/internal/biz"

	xlog "miopkg/log"
)

type Engine struct {
	application.Application
}

func NewEngine() *Engine {
	eng := &Engine{}
	if err := eng.Startup(
		eng.initJob,
		eng.newKafkaWorker,
		// eng.newKafkaC,
	); err != nil {
		xlog.Panic("startup", xlog.Any("err", err))
	}
	return eng
}

func (e *Engine) initJob() error {
	return e.Job(NewJobRunner())
}

type JobRunner struct {
	JobName string
}

func NewJobRunner() *JobRunner {
	return &JobRunner{
		JobName: "comment_job_runner",
	}
}

func (j *JobRunner) Run() {
	fmt.Println("i am job runner")
}

func (j *JobRunner) GetJobName() string {
	return j.JobName
}

type CommentServer struct {
	replyBiz          *biz.ReplyBiz
	likeBiz           *biz.LikeBiz
	commentMQConsumer *ekafka.Consumer
}

func NewCommentServer(replyBiz *biz.ReplyBiz, likeBiz *biz.LikeBiz, commentMQConsumer *ekafka.Consumer) *CommentServer {
	return &CommentServer{
		replyBiz:          replyBiz,
		likeBiz:           likeBiz,
		commentMQConsumer: commentMQConsumer,
	}
}
