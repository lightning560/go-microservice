package biz

type IDislikeRepo interface{}

type DislikeBiz struct {
	repo IDislikeRepo
}

func NewDislikeBiz(repo IDislikeRepo) *DislikeBiz {
	return &DislikeBiz{repo: repo}
}
