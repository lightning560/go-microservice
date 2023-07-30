package server

import (
	"fmt"
	"miopkg/errors"
	xgin "miopkg/http/gin"
	"miopkg/http/gin/resp"
	"miopkg/log"

	interfacemallv1 "redbook/api/interface/mall/v1"
	basev1 "redbook/api/redbookpb/v1"
	"redbook/app/interface/mall/model"
	"redbook/common/basemodel"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (mi *MallInterface) CreateCollection(c *gin.Context) {
	ctx := c.Request.Context()
	req := &interfacemallv1.CreateCollectionRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		log.Errorf("CreateCollection c.ShouldBind", err)
		resp.JSONErr(c, xgin.StatusGinBindError)
		return
	}
	cover := &basemodel.Image{}
	cover.FromProto(req.Cover)
	video := &basemodel.Video{}
	video.FromProto(req.Video)
	tags := &basemodel.Tags{}
	tags.ListFromProto(req.Tags)
	skusOnlyId := new(model.SkusOnlyId)
	skusOnlyId.ListFromInterfaceProto(req.SkusOnlyId)

	m := &model.CreateCollection{
		ShopId:     req.ShopId,
		Name:       req.Name,
		SkusOnlyId: *skusOnlyId,
		Cover:      *cover,
		Video:      *video,
		Tags:       *tags,
		PublishAt:  req.PublishAt,
	}
	id, err := mi.collectionBiz.CreateCollection(ctx, m)
	if err != nil {
		resp.JSONErr(c, errors.Code(err))
		return
	}
	idStr := strconv.FormatInt(id, 10)
	data := &interfacemallv1.CreateCollectionResponse_Data{
		Id: idStr,
	}
	resp.JSONSuccess(c, data)
}

func (mi *MallInterface) GetCollectionById(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		resp.JSONErr(c, xgin.StatusHttpRequestInvalidParam)
		return
	}
	req := &interfacemallv1.GetCollectionByIdRequest{
		Id: id,
	}
	if err := req.Validate(); err != nil {
		resp.JSONErr(c, xgin.StatusHttpRequestInvalidParam)
		return
	}
	m, err := mi.collectionBiz.GetCollectionById(ctx, req.Id)
	if err != nil {
		resp.JSONErr(c, errors.Code(err))
		return
	}
	data := &interfacemallv1.GetCollectionByIdResponse_Data{
		Collection: m.ToInterfaceProto(),
	}
	resp.JSONSuccess(c, data)
}

func (mi *MallInterface) GetCollectionCardById(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		resp.JSONErr(c, xgin.StatusHttpRequestInvalidParam)
		return
	}
	req := &interfacemallv1.GetCollectionCardByIdRequest{
		Id: id,
	}
	if err := req.Validate(); err != nil {
		resp.JSONErr(c, xgin.StatusHttpRequestInvalidParam)
		return
	}
	m, err := mi.collectionBiz.GetCollectionCardById(ctx, req.Id)
	if err != nil {
		resp.JSONErr(c, errors.Code(err))
		return
	}
	data := &interfacemallv1.GetCollectionCardByIdResponse_Data{
		CollectionCard: m.ToInterfaceProto(),
	}
	resp.JSONSuccess(c, data)
}

func (mi *MallInterface) ListCollectionCardByIds(c *gin.Context) {
	ctx := c.Request.Context()
	req := &interfacemallv1.ListCollectionCardByIdsRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorf("ListCollectionCardByIds c.ShouldBind:", err)
		resp.JSONErr(c, xgin.StatusGinBindError)
		return
	}
	if err := req.Validate(); err != nil {
		log.Errorf("ListCollectionCardByIds req.Validate", err)
		resp.JSONErr(c, xgin.StatusHttpRequestValidateError)
		return
	}
	fmt.Printf("http req.Ids:%v\n", req.Ids)
	// ids := make([]int64, 0, len(req.Ids))
	// for _, id := range req.Ids {
	// 	ids = append(ids, id.Id)
	// }
	idsInt64 := make([]int64, 0, len(req.Ids))
	for _, idStr := range req.Ids {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			resp.JSONErr(c, xgin.StatusHttpRequestInvalidParam)
			return
		}
		idsInt64 = append(idsInt64, id)
	}
	m, err := mi.collectionBiz.ListCollectionCardByIds(ctx, idsInt64)
	if err != nil {
		resp.JSONErr(c, errors.Code(err))
		return
	}
	data := &interfacemallv1.ListCollectionCardByIdsResponse_Data{
		CollectionCards: m.ListToInterfaceProto(),
	}
	resp.JSONSuccess(c, data)
}

func (mi *MallInterface) ListCollectionCardByShopId(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	offsetStr := c.Param("offset")
	limitStr := c.Param("limit")
	sortBy := c.Param("by")
	sortOrder := c.Param("order")
	fmt.Println("offsetStr", offsetStr, "limitStr", limitStr, "sortBy", sortBy, "sortOrder", sortOrder)
	id, error := strconv.ParseInt(idStr, 10, 64)
	if error != nil {
		log.Errorf("ListPostCard ParseInt id", error)
		resp.JSONErr(c, xgin.StatusHttpRequestInvalidParam)
		return
	}
	offsetInt64, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		log.Errorf("ListPostCard ParseInt offset", err)
		resp.JSONErr(c, xgin.StatusHttpRequestInvalidParam)
		return
	}
	offset := int32(offsetInt64)

	limitInt64, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		log.Errorf("ListPostCard ParseInt limit", err)
		resp.JSONErr(c, xgin.StatusHttpRequestInvalidParam)
		return
	}
	limit := int32(limitInt64)
	req := &interfacemallv1.ListCollectionCardByShopIdRequest{
		Id: id,
		Cursor: &basev1.Cursor{
			Offset: offset,
			Limit:  limit,
		},
		Sort: &basev1.Sort{
			By:    sortBy,
			Order: sortOrder,
		},
	}
	if err := req.Validate(); err != nil {
		resp.JSONErr(c, xgin.StatusHttpRequestValidateError)
		return
	}
	fmt.Println("req.Cursor.Offset::", req.Cursor.Offset)
	m, err := mi.collectionBiz.ListCollectionCardByShopId(ctx, req.Id, req.Cursor.Offset, req.Cursor.Limit, req.Sort.By, req.Sort.Order)
	if err != nil {
		log.Errorf("http.ListCollectionCardByShopId", err)
		resp.JSONErr(c, errors.Code(err))
		return
	}
	data := &interfacemallv1.ListCollectionCardByShopIdResponse_Data{
		CollectionCards: m.ListToInterfaceProto(),
	}
	resp.JSONSuccess(c, data)
}

func (mi *MallInterface) UpdateCollectionSku(c *gin.Context) {
	ctx := c.Request.Context()
	req := &interfacemallv1.UpdateCollectionSkuRequest{}
	if err := c.ShouldBind(req); err != nil {
		resp.JSONErr(c, xgin.StatusGinBindError)
		return
	}
	if err := req.Validate(); err != nil {
		resp.JSONErr(c, xgin.StatusHttpRequestValidateError)
		return
	}
	skusOnlyId := &model.SkusOnlyId{}
	skusOnlyId.ListFromInterfaceProto(req.SkusOnlyId)
	err := mi.collectionBiz.UpdateCollectionSku(ctx, req.Id, skusOnlyId)
	if err != nil {
		resp.JSONErr(c, errors.Code(err))
		return
	}
	resp.JSONSuccess(c, nil)
}

func (mi *MallInterface) UpdateCollectionState(c *gin.Context) {
	ctx := c.Request.Context()
	req := &interfacemallv1.UpdateCollectionStateRequest{}
	if err := c.ShouldBind(req); err != nil {
		resp.JSONErr(c, xgin.StatusGinBindError)
		return
	}
	if err := req.Validate(); err != nil {
		resp.JSONErr(c, xgin.StatusHttpRequestValidateError)
		return
	}
	err := mi.collectionBiz.UpdateCollectionState(ctx, req.Id, req.State)
	if err != nil {
		resp.JSONErr(c, errors.Code(err))
		return
	}
	resp.JSONSuccess(c, nil)
}
