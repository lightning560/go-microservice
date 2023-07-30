package biz

type IOrderRepo interface {
}

type OrderBiz struct {
	repo IOrderRepo
}

func NewOrderBiz(repo IOrderRepo) *OrderBiz {
	return &OrderBiz{
		repo: repo,
	}
}
