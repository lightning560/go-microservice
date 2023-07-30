package server

import (
	"fmt"
	"strconv"

	interfacemallv1 "redbook/api/interface/mall/v1"
	"redbook/common/basemodel"

	"miopkg/errors"
	xgin "miopkg/http/gin"
	"miopkg/http/gin/resp"
	"miopkg/log"

	"github.com/gin-gonic/gin"
)

func (mi *MallInterface) CreateShop(c *gin.Context) {
	ctx := c.Request.Context()
	var req interfacemallv1.CreateShopRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorf("CreateShop c.ShouldBindJSON err:%v", err)
		resp.JSONErr(c, xgin.StatusGinBindError)
		return
	}
	if err := req.Validate(); err != nil {
		resp.JSONErr(c, xgin.StatusHttpRequestValidateError)
		return
	}
	fmt.Println("CreateShop req:", req)
	logo := &basemodel.Image{}
	logo.FromProto(req.Logo)
	shopId, err := mi.shopBiz.CreateShop(ctx, req.Uid, req.Name, req.Sign, logo)
	if err != nil {
		fmt.Println("err:", err)
		e := errors.FromError(err)
		fmt.Println("e:", e)
		resp.JSONErr(c, errors.Code(err))
		return
	}
	shopIdStr := strconv.FormatInt(shopId, 10)
	data := &interfacemallv1.CreateShopResponse_Data{
		Id: shopIdStr,
	}
	resp.JSONSuccess(c, data)
}

func (mi *MallInterface) GetShopById(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		resp.JSONErr(c, xgin.StatusHttpRequestInvalidParam)
		return
	}
	req := &interfacemallv1.GetShopByIdRequest{
		Id: id,
	}
	if err := req.Validate(); err != nil {
		resp.JSONErr(c, xgin.StatusHttpRequestValidateError)
		return
	}
	m, err := mi.shopBiz.GetShopById(ctx, id)
	if err != nil {
		resp.JSONErr(c, errors.Code(err))
		return
	}
	data := &interfacemallv1.GetShopByIdResponse_Data{
		Shop: m.ToInterfaceProto(),
	}
	resp.JSONSuccess(c, data)
}
