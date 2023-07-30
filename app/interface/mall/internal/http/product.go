package server

import (
	"fmt"
	interfacemallv1 "redbook/api/interface/mall/v1"
	"redbook/app/interface/mall/model"
	"strconv"

	"github.com/gin-gonic/gin"

	"miopkg/errors"
	xgin "miopkg/http/gin"
	"miopkg/http/gin/resp"
	"miopkg/log"
)

func (mi *MallInterface) CreateProduct(c *gin.Context) {
	ctx := c.Request.Context()
	// req := &interfacemallv1.CreateProductRequest{}
	var req interfacemallv1.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorf("MallInterface # CreateProduct # ShouldBindJSON: %v", err)
		resp.JSONErr(c, xgin.StatusGinBindError)
		return
	}
	if err := req.Validate(); err != nil {
		resp.JSONErr(c, errors.Code(err))
		return
	}
	fmt.Println("MallInterface# req.CreateProduct:", req.CreateProduct)
	m := &model.CreateProduct{}
	m.FromInterfaceProto(req.CreateProduct)
	id, err := mi.productBiz.CreateProduct(ctx, m)
	if err != nil {
		resp.JSONErr(c, errors.Code(err))
		return
	}
	idStr := strconv.FormatInt(id, 10)
	data := &interfacemallv1.CreateProductResponse_Data{
		Id: idStr,
	}
	resp.JSONSuccess(c, data)
}

func (mi *MallInterface) GetProductById(c *gin.Context) {
	ctx := c.Request.Context()
	req := &interfacemallv1.GetProductByIdRequest{}
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		resp.JSONErr(c, xgin.StatusHttpRequestInvalidParam)
		return
	}
	req.Id = id
	if err := req.Validate(); err != nil {
		resp.JSONErr(c, xgin.StatusHttpRequestValidateError)
		return
	}
	m, err := mi.productBiz.GetProductById(ctx, req.Id)
	if err != nil {
		resp.JSONErr(c, errors.Code(err))
		return
	}
	data := &interfacemallv1.GetProductByIdResponse_Data{
		Product: m.ToInterfaceProto(),
	}
	resp.JSONSuccess(c, data)
}
