package server

import (
	"fmt"
	xgin "miopkg/http/gin"
	resp "miopkg/http/gin/resp"
	xlog "miopkg/log"
	feedv1 "redbook/api/interface/feed/v1"
	basev1 "redbook/api/redbookpb/v1"
	"redbook/app/interface/feed/internal/biz/model"
	"redbook/common/basemodel"
	"strconv"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/gin-gonic/gin"
)

func (fi *FeedInterface) CreatePost(c *gin.Context) {
	req := new(feedv1.CreatePostReq)
	if err := c.ShouldBindJSON(req); err != nil {
		xlog.Errorf("CreatePost bind err", err)
		resp.JSONErr(c, xgin.StatusGinBindError)
		return
	}
	if err := req.Validate(); err != nil {
		xlog.Errorf("CreatePost validate err", err)
		resp.JSONErr(c, xgin.StatusHttpRequestValidateError)
		return
	}
	fmt.Printf("interface # CreatePost req:%+v\n", req)
	m := &model.Post{}
	m.Title = req.Title
	m.Content = req.Content

	cover := &basemodel.Image{}
	cover.FromProto(req.Cover)
	m.Cover = cover
	m.Uid = req.Uid
	m.Type = req.Type
	fmt.Println("req.Tags", req.Tags)
	tags := &basemodel.Tags{}
	tags.ListFromProto(req.Tags)
	m.Tags = tags
	video := &basemodel.Video{}
	video.FromProto(req.Video)
	m.Video = video
	images := &basemodel.Images{}
	images.ListFromProto(req.Images)
	m.Images = images

	fmt.Printf("interface # CreatePost m:%+v\n", m)
	id, err := fi.feedbiz.CreatePost(c.Request.Context(), m)
	if err != nil {
		xlog.Errorf("FeedInterface CreatePost err", err)
		resp.JSONErr(c, xgin.StatusInternalServerError)
		return
	}
	resPB := &feedv1.CreatePostResponse{
		Data: &feedv1.CreatePostResponse_Data{
			Id: strconv.FormatInt(id, 10),
		},
	}
	resp.JSONSuccess(c, resPB.Data)
}

func (fi *FeedInterface) GetPostById(c *gin.Context) {
	var (
		id  int64
		err error
	)
	strId := c.Param("id")
	if id, err = strconv.ParseInt(strId, 10, 64); err != nil {
		xlog.Errorf("GetPostById ParseInt id", err)
		resp.JSONErr(c, xgin.StatusGinBindError)
		return
	}
	req := &feedv1.GetPostByIdReq{
		Id: id,
	}
	err = req.Validate()
	if err != nil {
		xlog.Errorf("GetPostById validate err", err)
		resp.JSONErr(c, xgin.StatusHttpRequestValidateError)
		return
	}
	m, err := fi.feedbiz.GetPost(c.Request.Context(), req.Id)
	if err != nil {
		xlog.Errorf("FeedInterface GetPost err", err)
		resp.JSONErr(c, xgin.StatusInternalServerError)
		return
	}
	post := m.ToInterfaceProto()
	// post := &feedv1.Post{
	// 	Id:        strconv.FormatInt(m.Id, 10),
	// 	Title:     m.Title,
	// 	Content:   m.Content,
	// 	CreatedAt: m.CreatedAt,
	// 	UpdatedAt: m.UpdatedAt,
	// 	Cover: &basev1.Image{
	// 		Id:     m.Cover.Id,
	// 		Url:    m.Cover.Url,
	// 		Width:  m.Cover.Width,
	// 		Height: m.Cover.Height,
	// 		SizeKb: m.Cover.SizeKb,
	// 		Name:   m.Cover.Name,
	// 		Hash:   m.Cover.Hash,
	// 	},
	// 	Type: m.Type,
	// 	Video: &basev1.Video{
	// 		Id:     m.Video.Id,
	// 		Url:    m.Video.Url,
	// 		Width:  m.Video.Width,
	// 		Height: m.Video.Height,
	// 		SizeKb: m.Video.SizeKb,
	// 		Length: m.Video.Length,
	// 	},
	// 	Images:     make([]*basev1.Image, 0, len(*m.Images)),
	// 	LikeCount:  m.LikeCount,
	// 	ShareCount: m.ShareCount,
	// 	FavorCount: m.FavorCount,
	// 	Uid:        m.Uid,
	// }

	// for _, v := range *m.Images {
	// 	post.Images = append(post.Images, &basev1.Image{
	// 		Id:     v.Id,
	// 		Url:    v.Url,
	// 		Width:  v.Width,
	// 		Height: v.Height,
	// 		SizeKb: v.SizeKb,
	// 		Name:   v.Name,
	// 		Hash:   v.Hash,
	// 	})
	// }
	info, err := fi.userbiz.GetUserInfoByUid(c.Request.Context(), m.Uid)
	if err != nil {
		xlog.Errorf("FeedInterface GetPost GetUserProfileByUid err", err)
		resp.JSONErr(c, xgin.StatusInternalServerError)
		return
	}
	post.AuthorInfo = info.ToProto()
	resPB := &feedv1.GetPostByIdResponse{
		Data: &feedv1.GetPostByIdResponse_Data{
			Post: post,
		},
	}
	resp.JSONSuccess(c, resPB.Data)
}

func (fi *FeedInterface) ListPostCard(c *gin.Context) {
	st := mapset.NewSet[int64]()

	offsetStr := c.Param("offset")
	limitStr := c.Param("limit")
	sortBy := c.Param("sort_by")
	sortOrder := c.Param("sort_order")
	// fmt.Println("offsetStr", offsetStr, "limitStr", limitStr, "sortBy", sortBy, "sortOrder", sortOrder)
	offsetInt64, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		xlog.Errorf("ListPostCard ParseInt offset", err)
		resp.JSONErr(c, xgin.StatusGinBindError)
		return
	}
	offset := int32(offsetInt64)
	limitInt64, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		xlog.Errorf("ListPostCard ParseInt limit", err)
		resp.JSONErr(c, xgin.StatusGinBindError)
		return
	}
	limit := int32(limitInt64)
	req := &feedv1.ListPostCardReq{
		Cursor: &basev1.Cursor{
			Offset: offset,
			Limit:  limit,
		},
		Sort: &basev1.Sort{
			By:    sortBy,
			Order: sortOrder,
		},
	}
	err = req.Validate()
	if err != nil {
		xlog.Errorf("ListPostCard Validate err", err)
		resp.JSONErr(c, xgin.StatusHttpRequestValidateError)
		return
	}
	ms, err := fi.feedbiz.ListPostCard(c.Request.Context(), req.Cursor.Offset, req.Cursor.Limit, req.Sort.By, req.Sort.Order)
	if err != nil {
		xlog.Errorf("FeedInterface ListPostCard err", err)
		resp.JSONErr(c, xgin.StatusInternalServerError)
		return
	}

	for _, v := range *ms {
		st.Add(v.Uid)
	}
	infos, err := fi.userbiz.MapUserInfoByUids(c.Request.Context(), st.ToSlice())
	if err != nil {
		xlog.Errorf("FeedInterface ListPostCard MapUserInfoByUids err", err)
		resp.JSONErr(c, xgin.StatusInternalServerError)
		return
	}
	cards := ms.ListToInterfaceProto()
	for _, v := range cards {
		v.AuthorInfo = infos[v.Uid].ToProto()
	}
	resPB := &feedv1.ListPostCardResponse{
		Data: &feedv1.ListPostCardResponse_Data{
			PostCards: cards,
			Total:     int32(len(*ms)),
			Cursor:    req.Cursor,
		},
	}
	resp.JSONSuccess(c, resPB.Data)
}

type ListPostCardByIdsReq struct {
	state         string
	sizeCache     string
	unknownFields string

	Ids []*ListPostCardByIdsReq_Id `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

type ListPostCardByIdsReq_Id struct {
	state string
	Id    int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}
type IdsReq struct {
	Ids []Id `json:"ids" binding:"required"`
}
type Id struct {
	Id int64 `json:"id" binding:"required"`
}

func (fi *FeedInterface) ListPostCardByIds(c *gin.Context) {
	uidSet := mapset.NewSet[int64]()
	idSet := mapset.NewSet[int64]()
	// var req *feedv1.ListPostCardByIdsReq
	// var req *ListPostCardByIdsReq
	fmt.Println("start---interface---ListPostCardByIds")
	// var r IdsReq
	// var r ListPostCardByIdsReq
	// err := c.ShouldBindJSON(&r)
	// if err != nil {
	// 	fmt.Println(err)
	// 	c.String(http.StatusOK, "error")
	// 	return
	// }
	// fmt.Println("bind r:", r)
	// for _, v := range r.Ids {
	// 	fmt.Println("id:", v.Id)
	// }
	// c.String(http.StatusOK, "ok")
	// return
	// var req IdsReq
	var req ListPostCardByIdsReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		xlog.Errorf("ListPostCardByIds bind err", err)
		resp.JSONErr(c, xgin.StatusGinBindError)
		return
	}
	// if err = req.Validate();err != nil {
	// 	xlog.Errorf("ListPostCardByIds Validate err", err)
	// 	resp.JSONErr(c, xgin.StatusValidateError)
	// 	return
	// }
	fmt.Println("req", req)
	for _, v := range req.Ids {
		fmt.Println("id:", v.Id)
	}

	for _, v := range req.Ids {
		idSet.Add(v.Id)
	}
	ms, err := fi.feedbiz.ListPostCardByIds(c.Request.Context(), idSet.ToSlice())
	if err != nil {
		xlog.Errorf("FeedInterface ListPostCardByIds err", err)
		resp.JSONErr(c, xgin.StatusInternalServerError)
		return
	}
	// uids := make([]int64, 0)
	for _, v := range ms {
		fmt.Println("http--model.PostCard:", ms)
		// uids = append(uids, v.Uid)
		uidSet.Add(v.Uid)
	}
	infos, err := fi.userbiz.MapUserInfoByUids(c.Request.Context(), uidSet.ToSlice())
	if err != nil {
		xlog.Errorf("FeedInterface ListPostCard MapUserInfoByUids err", err)
		resp.JSONErr(c, xgin.StatusInternalServerError)
		return
	}
	cards := ms.ListToInterfaceProto()
	for _, v := range cards {
		v.AuthorInfo = infos[v.Uid].ToProto()
	}
	resPB := &feedv1.ListPostCardResponse{
		Data: &feedv1.ListPostCardResponse_Data{
			PostCards: cards,
			Total:     int32(len(ms)),
		},
	}
	resp.JSONSuccess(c, resPB.Data)
}

func (fi *FeedInterface) ListVideoPost(c *gin.Context) {
	uidSet := mapset.NewSet[int64]()
	offsetStr := c.Param("offset")
	limitStr := c.Param("limit")
	sortBy := c.Param("sort_by")
	sortOrder := c.Param("sort_order")
	offsetInt64, err := strconv.ParseInt(offsetStr, 10, 32)
	if err != nil {
		xlog.Errorf("ListVideoPost ParseInt offset", err)
		resp.JSONErr(c, xgin.StatusGinBindError)
		return
	}
	offset := int32(offsetInt64)
	limitInt64, err := strconv.ParseInt(limitStr, 10, 32)
	if err != nil {
		xlog.Errorf("ListVideoPost ParseInt limit", err)
		resp.JSONErr(c, xgin.StatusGinBindError)
		return
	}
	limit := int32(limitInt64)
	req := &feedv1.ListVideoPostReq{
		Cursor: &basev1.Cursor{
			Offset: offset,
			Limit:  limit,
		},
		Sort: &basev1.Sort{
			By:    sortBy,
			Order: sortOrder,
		},
	}
	ms, err := fi.feedbiz.ListVideoPost(c.Request.Context(), req.Cursor.Offset, req.Cursor.Limit, req.Sort.By, req.Sort.Order)
	if err != nil {
		xlog.Errorf("FeedInterface ListVideoPost err", err)
		resp.JSONErr(c, xgin.StatusInternalServerError)
		return
	}
	if len(*ms) == 0 {
		resp.JSONErr(c, xgin.StatusNotFound)
		return
	}
	for _, v := range *ms {
		uidSet.Add(v.Uid)
	}
	infos, err := fi.userbiz.MapUserInfoByUids(c.Request.Context(), uidSet.ToSlice())
	if err != nil {
		xlog.Errorf("FeedInterface ListVideoPost MapUserInfoByUids err", err)
		resp.JSONErr(c, xgin.StatusInternalServerError)
		return
	}
	posts := ms.ListToInterfaceProto()
	for _, v := range posts {
		v.AuthorInfo = infos[v.Uid].ToProto()
	}
	resPB := &feedv1.ListVideoPostResponse{
		Data: &feedv1.ListVideoPostResponse_Data{
			Posts:  posts,
			Total:  int32(len(*ms)),
			Cursor: req.Cursor,
		},
	}
	resp.JSONSuccess(c, resPB.Data)
}
