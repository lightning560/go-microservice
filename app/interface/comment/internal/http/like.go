package server

import (
	xgin "miopkg/http/gin"
	"miopkg/http/gin/resp"
	xlog "miopkg/log"
	"net/http"
	interfacecommentv1 "redbook/api/interface/comment/v1"

	"github.com/gin-gonic/gin"
)

func (s *CommentInterface) AddLike(c *gin.Context) {
	var (
		err error
	)
	ctx := c.Request.Context()
	var req interfacecommentv1.AddLikeReq
	if err = c.ShouldBindJSON(&req); err != nil {
		xlog.Errorf("AddLike ShouldBindJSON err", err)
		resp.JSONErr(c, xgin.StatusGinBindError)
		return
	}
	if err = req.Validate(); err != nil {
		xlog.Errorf("AddLike Validate err", err)
		resp.JSONErr(c, xgin.StatusHttpRequestValidateError)
		return
	}
	if err = s.likebiz.AddLike(ctx, req.ReplyId, req.Uid); err != nil {
		xlog.Errorf("AddLike err", err)
		resp.JSONErr(c, xgin.StatusBadRequest)
		return
	}
	resp.JSON(c, http.StatusOK, nil)
}

func (s *CommentInterface) CancelLike(c *gin.Context) {
	var (
		err error
	)
	ctx := c.Request.Context()
	var req interfacecommentv1.CancelLikeReq
	if err = c.ShouldBindJSON(&req); err != nil {
		xlog.Errorf("CancelLike ShouldBindJSON err", err)
		resp.JSONErr(c, xgin.StatusGinBindError)
		return
	}
	if err = req.Validate(); err != nil {
		xlog.Errorf("CancelLike Validate err", err)
		resp.JSONErr(c, xgin.StatusHttpRequestValidateError)
		return
	}
	if err = s.likebiz.CancelLike(ctx, req.ReplyId, req.Uid); err != nil {
		xlog.Errorf("CancelLike err", err)
		resp.JSONErr(c, xgin.StatusBadRequest)
		return
	}
	resp.JSON(c, http.StatusOK, nil)
}
