package data

import "redbook/app/domain/feed/internal/data/entity"

type historyRepo struct {
	data *Data
}

func NewHistoryRepo(data *Data) *historyRepo {
	return &historyRepo{data: data}
}

func (r *historyRepo) ListHistoryItems(uid int64, offset, limit int32) (items []*entity.HistoryItems, err error) {
	//  TODO: implement this method
	return nil, nil
}
