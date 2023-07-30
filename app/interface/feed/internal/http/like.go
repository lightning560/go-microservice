package server

import (
	xgin "miopkg/http/gin"
	"miopkg/http/gin/resp"
	xlog "miopkg/log"
	feedv1 "redbook/api/interface/feed/v1"

	"github.com/gin-gonic/gin"
)

func (fi *FeedInterface) AddLike(c *gin.Context) {
	req := new(feedv1.AddLikeReq)
	if err := c.ShouldBindJSON(req); err != nil {
		xlog.Errorf("AddLike bind err", err)
		resp.JSONErr(c, xgin.StatusGinBindError)
		return
	}
	if err := req.Validate(); err != nil {
		xlog.Errorf("AddLike validate err", err)
		resp.JSONErr(c, xgin.StatusHttpRequestValidateError)
		return
	}
	if err := fi.likebiz.AddLike(c.Request.Context(), req.Uid, req.PostId); err != nil {
		xlog.Errorf("FeedInterface AddLike err", err)
		resp.JSONErr(c, xgin.StatusInternalServerError)
		return
	}
	resp.JSONSuccess(c, nil)
}

func (fi *FeedInterface) CancelLike(c *gin.Context) {
	req := new(feedv1.CancelLikeReq)
	if err := c.ShouldBindJSON(req); err != nil {
		xlog.Errorf("CancelLike bind err", err)
		resp.JSONErr(c, xgin.StatusGinBindError)
		return
	}
	if err := req.Validate(); err != nil {
		xlog.Errorf("CancelLike validate err", err)
		resp.JSONErr(c, xgin.StatusHttpRequestValidateError)
		return
	}
	if err := fi.likebiz.CancelLike(c.Request.Context(), req.Uid, req.PostId); err != nil {
		xlog.Errorf("FeedInterface CancelLike err", err)
		resp.JSONErr(c, xgin.StatusInternalServerError)
		return
	}
	resp.JSONSuccess(c, nil)
}
