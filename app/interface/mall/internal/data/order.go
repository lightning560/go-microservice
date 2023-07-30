package data

type orderRepo struct {
	data *Data
}

func NewOrderRepo(data *Data) *orderRepo {
	return &orderRepo{
		data: data,
	}
}
