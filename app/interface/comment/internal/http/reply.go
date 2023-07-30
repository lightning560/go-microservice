package server

import (
	"fmt"
	"miopkg/log"
	"net/http"
	"strconv"

	basev1 "redbook/api/redbookpb/v1"
	"redbook/app/interface/comment/internal/biz/model"
	"redbook/common/basemodel"

	xgin "miopkg/http/gin"
	"miopkg/http/gin/resp"
	xlog "miopkg/log"
	interfacecommentv1 "redbook/api/interface/comment/v1"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/gin-gonic/gin"
)

func (ci *CommentInterface) CreateReply(c *gin.Context) {
	var (
		err error
		req interfacecommentv1.CreateReplyReq
	)
	ctx := c.Request.Context()
	if err = c.ShouldBindJSON(&req); err != nil {
		xlog.Errorf("CreateReply ShouldBindJSON err", err)
		resp.JSONErr(c, xgin.StatusGinBindError)
		return
	}
	if err = req.Validate(); err != nil {
		xlog.Errorf("CreateReply Validate err", err)
		resp.JSONErr(c, xgin.StatusHttpRequestValidateError)
		return
	}
	fmt.Printf("CreateReply req:%+v\n", req)
	if err = ci.replybiz.CreateReply(ctx, req.IsFloor, req.SubjectId, req.BizType, req.Uid, req.FloorId, req.Content, req.AtUid); err != nil {
		xlog.Errorf("CreateReply err", err)
		resp.JSONErr(c, xgin.StatusBadRequest)
		return
	}
	resp.JSON(c, http.StatusOK, nil)
}

func (ci *CommentInterface) ListFloorBySubject(c *gin.Context) {
	var (
		err                         error
		floors                      *model.Floors
		MapUserInfo                 map[int64]*basemodel.UserInfo
		id, offsetInt64, limitInt64 int64
	)
	ctx := c.Request.Context()
	uidSet := mapset.NewSet[int64]()

	subjectIdStr := c.Param("subject_id")
	offsetStr := c.Param("offset")
	limitStr := c.Param("limit")
	sortBy := c.Param("sort_by")
	sortOrder := c.Param("sort_order")
	if id, err = strconv.ParseInt(subjectIdStr, 10, 64); err != nil {
		xlog.Errorf("ListFloorBySubject ParseInt subject_id", err)
		resp.JSONErr(c, xgin.StatusHttpRequestInvalidParam)
		return
	}
	if offsetInt64, err = strconv.ParseInt(offsetStr, 10, 64); err != nil {
		xlog.Errorf("ListFloorBySubject ParseInt offset", err)
		resp.JSONErr(c, xgin.StatusHttpRequestInvalidParam)
		return
	}
	offset := int32(offsetInt64)
	if limitInt64, err = strconv.ParseInt(limitStr, 10, 64); err != nil {
		xlog.Errorf("ListFloorBySubject ParseInt limit", err)
		resp.JSONErr(c, xgin.StatusHttpRequestInvalidParam)
		return
	}
	limit := int32(limitInt64)
	req := &interfacecommentv1.ListFloorBySubjectReq{
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
	if err = req.Validate(); err != nil {
		xlog.Errorf("ListFloorBySubject Validate err", err)
		resp.JSONErr(c, xgin.StatusBadRequest)
		return
	}
	if sortBy != "reply" && sortBy != "time" {
		xlog.Errorf("ListFloorBySubject sort_by err", err)
		resp.JSONErr(c, xgin.StatusBadRequest)
		return
	}
	// sortby reply,reply_count
	if sortBy == "reply" {
		if floors, err = ci.replybiz.ListFloorBySubjectSortReply(ctx, req.Id, req.Cursor.Offset, req.Cursor.Limit, req.Sort.By, req.Sort.Order); err != nil {
			xlog.Errorf("ListFloorBySubjectSortReply err", err)
			resp.JSONErr(c, xgin.StatusInternalServerError)
			return
		}
	}
	// sortby time,created_at
	if sortBy == "time" {
		if floors, err = ci.replybiz.ListFloorBySubjectSortTime(ctx, req.Id, req.Cursor.Offset, req.Cursor.Limit, req.Sort.By, req.Sort.Order); err != nil {
			xlog.Errorf("ListFloorBySubjectSortTime err", err)
			resp.JSONErr(c, xgin.StatusInternalServerError)
			return
		}
	}
	// get uidInfo
	for _, floor := range *floors {
		uidSet.Add(floor.RootReply.OwnerUid)
		for _, reply := range floor.Replies {
			uidSet.Add(reply.OwnerUid)
		}
	}
	fmt.Println("ListFloorBySubject uidSet:", uidSet)
	if MapUserInfo, err = ci.userbiz.MapUserInfoByUids(ctx, uidSet.ToSlice()); err != nil {
		xlog.Errorf("ListFloorBySubject invoke MapUserInfoByUids err", err)
		resp.JSONErr(c, xgin.StatusInternalServerError)
		return
	}
	data := &interfacecommentv1.ListFloorBySubjectResponse_Data{}
	data.Cursor = req.Cursor
	data.Floors = floors.ListToInterfaceProto()
	data.Total = int32(len(*floors))
	fmt.Printf("data total:%d\n", data.Total)
	for _, floor := range data.Floors {
		rootOwnerUidInt64, _ := strconv.ParseInt(floor.RootReply.OwnerUid, 10, 64)
		floor.RootReply.UserInfo = MapUserInfo[rootOwnerUidInt64].ToProto()
		for _, reply := range floor.Replies {
			replyOwnerUidInt64, _ := strconv.ParseInt(reply.OwnerUid, 10, 64)
			reply.UserInfo = MapUserInfo[replyOwnerUidInt64].ToProto()
		}
	}
	fmt.Printf("data total:%d\n", data.Total)
	resp.JSONSuccess(c, data)
}

func (ci *CommentInterface) ListReplyByFloor(c *gin.Context) {
	// 1.init var
	var (
		err                                         error
		replies                                     *model.Replies
		MapUserInfo                                 map[int64]*basemodel.UserInfo
		subjectId, floorId, offsetInt64, limitInt64 int64
	)
	ctx := c.Request.Context()
	uidSet := mapset.NewSet[int64]()
	// 2.parse params
	subjectIdStr := c.Param("subject_id")
	floorIdStr := c.Param("floor_id")
	offsetStr := c.Param("offset")
	limitStr := c.Param("limit")
	sortBy := c.Param("sort_by")
	sortOrder := c.Param("sort_order")
	if subjectId, err = strconv.ParseInt(subjectIdStr, 10, 64); err != nil {
		xlog.Errorf("ListReplyByFloor ParseInt subject_id", err)
		resp.JSONErr(c, xgin.StatusHttpRequestInvalidParam)
		return
	}
	if floorId, err = strconv.ParseInt(floorIdStr, 10, 64); err != nil {
		xlog.Errorf("ListReplyByFloor ParseInt floor_id", err)
		resp.JSONErr(c, xgin.StatusHttpRequestInvalidParam)
		return
	}
	if offsetInt64, err = strconv.ParseInt(offsetStr, 10, 64); err != nil {
		xlog.Errorf("ListReplyByFloor ParseInt offset", err)
		resp.JSONErr(c, xgin.StatusHttpRequestInvalidParam)
		return
	}
	offset := int32(offsetInt64)
	if limitInt64, err = strconv.ParseInt(limitStr, 10, 64); err != nil {
		xlog.Errorf("ListReplyByFloor ParseInt limit", err)
		resp.JSONErr(c, xgin.StatusHttpRequestInvalidParam)
		return
	}
	limit := int32(limitInt64)
	log.Infof("ListReplyByFloor before insert to req --- subjectId:%d,floorId:%d,offset:%d,limit:%d,sortBy:%s,sortOrder:%s\n", subjectId, floorId, offset, limit, sortBy, sortOrder)
	req := &interfacecommentv1.ListReplyByFloorReq{
		SubjectId: subjectId,
		FloorId:   floorId,
		Cursor: &basev1.Cursor{
			Offset: offset,
			Limit:  limit,
		},
		Sort: &basev1.Sort{
			By:    sortBy,
			Order: sortOrder,
		},
	}
	log.Infof("ListReplyByFloor before validate req:%+v\n", req)
	// 3.validate params
	if err = req.Validate(); err != nil {
		xlog.Errorf("ListReplyByFloor  Validate err", err)
		resp.JSONErr(c, xgin.StatusHttpRequestValidateError)
		return
	}
	log.Infof("ListReplyByFloor after validate req:%+v\n", req)
	log.Infof("ListReplyByFloor after validate req.Cursor.Offset:%d", req.Cursor.Offset)
	// 4. biz
	if sortBy != "like" && sortBy != "time" {
		xlog.Errorf("ListReplyByFloor sort_by err", err)
		resp.JSONErr(c, xgin.StatusHttpRequestInvalidParam)
		return
	}
	// sortby reply,reply_count
	if sortBy == "like" {
		log.Infof("ListReplyByFloor before ListReplyByFloorSortLike req.Cursor.GetOffset():%d", req.Cursor.GetOffset())
		if replies, err = ci.replybiz.ListReplyByFloorSortLike(ctx, req.SubjectId, req.FloorId, req.Cursor.GetOffset(), req.Cursor.GetLimit(), req.Sort.By, req.Sort.Order); err != nil {
			xlog.Errorf("ListReplyByFloorSortLike err", err)
			resp.JSONErr(c, xgin.StatusInternalServerError)
			return
		}
	}
	// sortby time,created_at
	if sortBy == "time" {
		if replies, err = ci.replybiz.ListReplyByFloorSortTime(ctx, req.SubjectId, req.FloorId, req.Cursor.GetOffset(), req.Cursor.GetLimit(), req.Sort.By, req.Sort.Order); err != nil {
			xlog.Errorf("ListReplyByFloorSortTime err", err)
			resp.JSONErr(c, xgin.StatusInternalServerError)
			return
		}
	}
	// 5. get uidInfo
	fmt.Println("http#ListReplyByFloor len(replies)", len(*replies))
	for _, reply := range *replies {
		uidSet.Add(reply.OwnerUid)
	}
	fmt.Println("ListReplyByFloor uidSet", uidSet)
	if MapUserInfo, err = ci.userbiz.MapUserInfoByUids(ctx, uidSet.ToSlice()); err != nil {
		xlog.Errorf("ListReplyByFloor invoke MapUserInfoByUids err", err)
		resp.JSONErr(c, xgin.StatusInternalServerError)
		return
	}
	// 6. resp
	data := &interfacecommentv1.ListReplyByFloorResponse_Data{}
	data.Cursor = req.Cursor
	data.Replies = replies.ListToInterfaceProto()
	data.Total = int32(len(*replies))
	for _, reply := range data.Replies {
		ownerUidInt64, _ := strconv.ParseInt(reply.OwnerUid, 10, 64)
		reply.UserInfo = MapUserInfo[ownerUidInt64].ToProto()
	}
	resp.JSONSuccess(c, data)
}

func subjectFloorsByPage(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
func floorReplysByPage(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func floorReplysByCursor(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
func dialogReplysByPage(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func dialogReplysByCursor(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
func addReply(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
func delReply(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
func hideReply(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
func addTop(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
func cancelTop(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func addHots(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func cancelHots(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
func reportReply(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
func addLike(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
func cancelLike(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
func addDislike(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func cancelDislike(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
