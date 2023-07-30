package server

import (
	"fmt"
	"miopkg/errors"
	xgin "miopkg/http/gin"
	"miopkg/http/gin/resp"
	"miopkg/log"
	interfacemallv1 "redbook/api/interface/mall/v1"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (mi *MallInterface) CreateCartItem(c *gin.Context) {
	ctx := c.Request.Context()
	var req interfacemallv1.CreateCartItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorf("CreateCartItem c.ShouldBindJSON", err)
		resp.JSONErr(c, xgin.StatusGinBindError)
		return
	}
	if err := req.Validate(); err != nil {
		log.Errorf("CreateCartItem req.Validate", err)
		resp.JSONErr(c, xgin.StatusHttpRequestInvalidParam)
		return
	}
	rv, err := mi.cartBiz.CreateCartItem(ctx, req.Uid, req.CollectionId, req.ProductId, req.Quantity)
	if err != nil {
		log.Errorf("CreateCartItem mi.cartBiz.CreateCartItem", err)
		resp.JSONErr(c, errors.Code(err))
		return
	}
	data := &interfacemallv1.CreateCartItemResponse_Data{
		Id: rv,
	}
	resp.JSONSuccess(c, data)
}

// func (mi *MallInterface) GetCartItemById(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	var req interfacemallv1.GetCartItemByIdRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	rv, err := mi.cartBiz.GetCartByIdItem(ctx, req.Id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	resp := &interfacemallv1.GetCartItemByIdResponse{
// 		CartItem: rv.ToInterfaceProto(),
// 	}
// 	data, err := anypb.New(resp)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, basev1.Response{Code: 200, Message: "success", Data: data})
// }

func (mi *MallInterface) ListCartItemByUid(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Errorf("invalid uid: %s", idStr)
		resp.JSONErr(c, xgin.StatusHttpRequestInvalidParam)
		return
	}
	var req interfacemallv1.ListCartItemByUidRequest
	req.Uid = id
	if err := req.Validate(); err != nil {
		log.Errorf("invalid request: %v", err)
		resp.JSONErr(c, xgin.StatusHttpRequestValidateError)
		return
	}
	rv, err := mi.cartBiz.ListCartItemByUid(ctx, req.Uid)
	if err != nil {
		log.Errorf("list cart item by uid failed: %v", err)
		resp.JSONErr(c, errors.Code(err))
		return
	}
	fmt.Println("mi.ListCartItemByUid#rv", rv[0])
	ci := rv.ListToInterfaceProto()
	data := &interfacemallv1.ListCartItemByUidResponse_Data{
		CartItems: ci,
	}
	resp.JSONSuccess(c, data)
}

func (mi *MallInterface) UpdateCartItemQuantity(c *gin.Context) {
	ctx := c.Request.Context()
	var req interfacemallv1.UpdateCartItemQuantityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorf("UpdateCartItemQuantity c.ShouldBindJSON", err)
		resp.JSONErr(c, xgin.StatusGinBindError)
		return
	}
	if err := req.Validate(); err != nil {
		log.Errorf("UpdateCartItemQuantity req.Validate", err)
		resp.JSONErr(c, xgin.StatusHttpRequestInvalidParam)
		return
	}
	err := mi.cartBiz.UpdateCartItemQuantity(ctx, req.Id, req.Quantity)
	if err != nil {
		log.Errorf("UpdateCartItemQuantity mi.cartBiz.UpdateCartItemQuantity", err)
		resp.JSONErr(c, errors.Code(err))
		return
	}
	resp.JSONSuccess(c, nil)
}
