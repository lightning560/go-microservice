package server

import (
	xgin "miopkg/http/gin"
	"miopkg/http/gin/resp"
	xlog "miopkg/log"
	feedv1 "redbook/api/interface/feed/v1"

	"github.com/gin-gonic/gin"
)

func (fi *FeedInterface) AddShare(c *gin.Context) {
	req := new(feedv1.CancelShareReq)
	if err := c.ShouldBindJSON(req); err != nil {
		xlog.Errorf("AddShare bind err", err)
		resp.JSONErr(c, xgin.StatusGinBindError)
		return
	}
	if err := req.Validate(); err != nil {
		xlog.Errorf("AddShare validate err", err)
		resp.JSONErr(c, xgin.StatusHttpRequestValidateError)
		return
	}
	if err := fi.sharebiz.AddShare(c.Request.Context(), req.Uid, req.PostId); err != nil {
		xlog.Errorf("FeedInterface AddShare err", err)
		resp.JSONErr(c, xgin.StatusInternalServerError)
		return
	}
	resp.JSONSuccess(c, nil)
}

func (fi *FeedInterface) CancelShare(c *gin.Context) {
	req := new(feedv1.CancelShareReq)
	if err := c.ShouldBindJSON(req); err != nil {
		xlog.Errorf("CancelShare bind err", err)
		resp.JSONErr(c, xgin.StatusGinBindError)
		return
	}
	if err := req.Validate(); err != nil {
		xlog.Errorf("CancelShare validate err", err)
		resp.JSONErr(c, xgin.StatusHttpRequestValidateError)
		return
	}
	if err := fi.sharebiz.CancelShare(c.Request.Context(), req.Uid, req.PostId); err != nil {
		xlog.Errorf("FeedInterface CancelShare err", err)
		resp.JSONErr(c, xgin.StatusInternalServerError)
		return
	}
	resp.JSONSuccess(c, nil)
}
