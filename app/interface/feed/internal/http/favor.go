package server

import (
	xgin "miopkg/http/gin"
	"miopkg/http/gin/resp"
	"miopkg/log"
	feedv1 "redbook/api/interface/feed/v1"

	"github.com/gin-gonic/gin"
)

func (fi *FeedInterface) AddFavor(c *gin.Context) {
	req := new(feedv1.AddFavorReq)
	if err := c.ShouldBindJSON(req); err != nil {
		log.Errorf("AddFavor bind err", err)
		resp.JSONErr(c, xgin.StatusGinBindError)
		return
	}
	if err := req.Validate(); err != nil {
		log.Errorf("AddFavor validate err", err)
		resp.JSONErr(c, xgin.StatusHttpRequestValidateError)
		return
	}
	if err := fi.favorbiz.AddFavor(c.Request.Context(), req.Uid, req.PostId); err != nil {
		log.Errorf("FeedInterface AddFavor err", err)
		resp.JSONErr(c, xgin.StatusInternalServerError)
		return
	}
	resp.JSONSuccess(c, nil)
}

func (fi *FeedInterface) CancelFavor(c *gin.Context) {
	req := new(feedv1.CancelFavorReq)
	if err := c.ShouldBindJSON(req); err != nil {
		log.Errorf("CancelFavor bind err", err)
		resp.JSONErr(c, xgin.StatusGinBindError)
		return
	}
	if err := req.Validate(); err != nil {
		log.Errorf("CancelFavor validate err", err)
		resp.JSONErr(c, xgin.StatusHttpRequestValidateError)
		return
	}
	if err := fi.favorbiz.CancelFavor(c.Request.Context(), req.Uid, req.PostId); err != nil {
		log.Errorf("FeedInterface CancelFavor err", err)
		resp.JSONErr(c, xgin.StatusInternalServerError)
		return
	}
	resp.JSONSuccess(c, nil)
}
