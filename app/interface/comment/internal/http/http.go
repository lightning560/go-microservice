package server

import (
	"context"
	xgin "miopkg/http/gin"

	"redbook/app/interface/comment/internal/biz"
	"redbook/app/interface/comment/internal/data"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CommentInterface struct {
	subjectbiz *biz.SubjectBiz
	replybiz   *biz.ReplyBiz
	likebiz    *biz.LikeBiz

	userbiz *biz.UserBiz
}

func NewCommentInterface(replybiz *biz.ReplyBiz, likebiz *biz.LikeBiz, subjectbiz *biz.SubjectBiz, userbiz *biz.UserBiz) *CommentInterface {
	return &CommentInterface{
		subjectbiz: subjectbiz,
		replybiz:   replybiz,
		likebiz:    likebiz,
		userbiz:    userbiz,
	}
}

func (eng *Engine) NewHttpServer() error {
	ctx := context.Background()
	commentRpcClient := data.NewCommentRpcClient()
	userRpcClient := data.NewUserRpcClient()
	commentPubMQ := data.NewCommentPubMQ()

	dataInstance, _, _ := data.NewData(ctx, commentRpcClient, userRpcClient, commentPubMQ)

	subjectRepo := data.NewSubjectRepo(dataInstance)
	replyRepo := data.NewReplyRepo(dataInstance)
	likeRepo := data.NewLikeRepo(dataInstance)
	userRepo := data.NewUserRepo(dataInstance)

	subjectBiz := biz.NewSubjectBiz(subjectRepo)
	replyBiz := biz.NewReplyBiz(replyRepo)
	likebiz := biz.NewLikeBiz(likeRepo)
	userbiz := biz.NewUserBiz(userRepo)

	ci := *NewCommentInterface(replyBiz, likebiz, subjectBiz, userbiz)

	server := xgin.StdConfig("http").Build()
	server.Use(cors.Default())
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, "pong")
	})
	apiv1 := server.Group("/v1/comment")
	{
		apiv1.POST("/subject", ci.CreateSubject)
		apiv1.GET("/subject/:id", ci.GetSubjectById)
		//  get floors of subject by cursor
		apiv1.GET("/floors/:subject_id/:offset/:limit/:sort_by/:sort_order", ci.ListFloorBySubject)

		//  get comments of dialog by cursor
		// apiv1.GET("/dialog/cursor", dialogReplysByCursor)
		//  get comments of dialog by Page
		// apiv1.GET("/dialog/page", dialogReplysByPage)

		apiv1.POST("/reply", ci.CreateReply)
		// apiv1.DELETE("/reply", delReply)

		// get replies of floor by cursor
		apiv1.GET("/replies/:subject_id/:floor_id/:offset/:limit/:sort_by/:sort_order", ci.ListReplyByFloor)
		// creator&adminer
		// apiv1.POST("/reply/hide", hideReply)

		// apiv1.POST("/reply/top", addTop)
		// apiv1.DELETE("/reply/top", cancelTop)

		// apiv1.POST("/reply/hot", addHots)
		// apiv1.DELETE("/reply/hot", cancelHots)

		// apiv1.POST("/reply/report", reportReply)

		apiv1.POST("/like", ci.AddLike)
		apiv1.DELETE("/like", ci.CancelLike)

		// apiv1.POST("/dislike", addDislike)
		// apiv1.DELETE("/dislike", cancelDislike)
	}
	return eng.Serve(server)
}
