package data

import (
	"context"
	"redbook/app/domain/history/model/entity"
)

func (r *historyRepo) RawBatchAddHistory(ctx context.Context, entities []*entity.History) (err error) {
	docs := make([]interface{}, len(entities))
	for i, v := range entities {
		docs[i] = v
	}
	_, err = r.data.historyCollection.InsertMany(ctx, docs)
	if err != nil {
		return
	}
	return nil
}
