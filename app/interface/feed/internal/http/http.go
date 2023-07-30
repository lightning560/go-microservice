package server

import (
	"fmt"
	xgin "miopkg/http/gin"
	"net/http"
	feedv1 "redbook/api/interface/feed/v1"
	"redbook/app/interface/feed/internal/biz"
	"redbook/app/interface/feed/internal/data"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type FeedInterface struct {
	feedbiz  *biz.FeedBiz
	likebiz  *biz.LikeBiz
	sharebiz *biz.ShareBiz
	favorbiz *biz.FavorBiz
	userbiz  *biz.UserBiz
}

func NewFeedInterface(feedbiz *biz.FeedBiz, userbiz *biz.UserBiz, likebiz *biz.LikeBiz, sharebiz *biz.ShareBiz, favorbiz *biz.FavorBiz) *FeedInterface {
	return &FeedInterface{
		feedbiz:  feedbiz,
		likebiz:  likebiz,
		sharebiz: sharebiz,
		favorbiz: favorbiz,
		userbiz:  userbiz,
	}
}

func (eng *Engine) NewHttpServer() error {
	feedRpcClient := data.NewFeedRpcClient()
	userRpcClient := data.NewUserRpcClient()
	dataInstance := data.NewData(feedRpcClient, userRpcClient)

	feedRepo := data.NewFeedRepo(dataInstance)
	likeRepo := data.NewLikeRepo(dataInstance)
	favorRepo := data.NewFavorRepo(dataInstance)
	shareRepo := data.NewShareRepo(dataInstance)
	userRepo := data.NewUserRepo(dataInstance)

	feedbiz := biz.NewFeedBiz(feedRepo)
	likebiz := biz.NewLikeBiz(likeRepo)
	favorbiz := biz.NewFavorBiz(favorRepo)
	sharebiz := biz.NewShareBiz(shareRepo)
	userbiz := biz.NewUserBiz(userRepo)

	fi := *NewFeedInterface(feedbiz, userbiz, likebiz, sharebiz, favorbiz)

	server := xgin.StdConfig("http").Build()
	server.Use(cors.Default())
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello Gin")
	})
	server.GET("/v1/feed/ids", func(c *gin.Context) {
		// var r IdsReq
		// var r ListPostCardByIdsReq
		var r feedv1.ListPostCardByIdsReq
		err := c.ShouldBindJSON(&r)
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusOK, "error")
			return
		}
		fmt.Println("bind r:", r)
		for _, v := range r.Ids {
			fmt.Println("id:", v.Id)
		}
		c.String(http.StatusOK, "ok")
	})

	server.GET("/v1/feed/cards", fi.ListPostCardByIds)
	server.POST("/v1/feed/like", fi.AddLike)
	server.DELETE("/v1/feed/like", fi.CancelLike)
	server.POST("/v1/feed/share", fi.AddShare)
	server.DELETE("/v1/feed/share", fi.CancelShare)
	server.POST("/v1/feed/favor", fi.AddFavor)
	server.DELETE("/v1/feed/favor", fi.CancelFavor)
	apiv1 := server.Group("/v1")
	{
		apiv1.POST("/feed/post", fi.CreatePost)
		apiv1.GET("/feed/post/:id", fi.GetPostById)
		apiv1.GET("/feed/cards/:offset/:limit/:sort_by/:sort_order", fi.ListPostCard)
		// apiv1.GET("/feed/cds/", fi.ListPostCardByIds)
		// apiv1.PUT("/feed/post/:id", fi.UpdatePost)
		// apiv1.DELETE("/feed/post/:id", fi.DeletePost)
		apiv1.GET("/feed/videos/:offset/:limit/:sort_by/:sort_order", fi.ListVideoPost)
	}
	return eng.Serve(server)
}
