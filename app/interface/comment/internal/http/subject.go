package server

import (
	xgin "miopkg/http/gin"
	"miopkg/http/gin/resp"
	xlog "miopkg/log"
	"strconv"

	interfacecommentv1 "redbook/api/interface/comment/v1"
	"redbook/app/interface/comment/internal/biz/model"

	"github.com/gin-gonic/gin"
)

func (ci *CommentInterface) CreateSubject(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(interfacecommentv1.CreateSubjectReq)
	if err := c.ShouldBindJSON(req); err != nil {
		xlog.Errorf("CreateSubject BindJSON", err)
		resp.JSONErr(c, xgin.StatusGinBindError)
		return
	}
	if err := req.Validate(); err != nil {
		xlog.Errorf("CreateSubject validate err", err)
		resp.JSONErr(c, xgin.StatusHttpRequestValidateError)
		return
	}
	var (
		id  int64
		err error
	)
	if id, err = ci.subjectbiz.CreateSubject(ctx, req.BelongId, req.BizType, req.OwnerUid); err != nil {
		xlog.Errorf("CreateSubject err", err)
		resp.JSONErr(c, xgin.StatusInternalServerError)
		return
	}
	idStr := strconv.FormatInt(id, 10)
	data := &interfacecommentv1.CreateSubjectResponse_Data{
		SubjectId: idStr,
	}
	resp.JSONSuccess(c, data)
}

func (ci *CommentInterface) GetSubjectById(c *gin.Context) {
	var (
		id  int64
		rv  *model.Subject
		err error
	)
	ctx := c.Request.Context()
	req := new(interfacecommentv1.GetSubjectByIdReq)
	idStr := c.Param("id")
	if id, err = strconv.ParseInt(idStr, 10, 64); err != nil {
		xlog.Errorf("GetSubjectById BindJSON", err)
		resp.JSONErr(c, xgin.StatusGinBindError)
		return
	}
	req.SubjectId = id
	if err = req.Validate(); err != nil {
		xlog.Errorf("GetSubjectById validate err", err)
		resp.JSONErr(c, xgin.StatusHttpRequestValidateError)
		return
	}
	if rv, err = ci.subjectbiz.GetSubjectById(ctx, req.SubjectId); err != nil {
		xlog.Errorf("GetSubjectById err", err)
		resp.JSONErr(c, xgin.StatusInternalServerError)
		return
	}
	data := &interfacecommentv1.GetSubjectByIdResponse_Data{}
	data.Subject = rv.ToInterfaceProto()
	resp.JSONSuccess(c, data)
}
