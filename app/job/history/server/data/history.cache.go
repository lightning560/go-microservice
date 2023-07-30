package data

import (
	"context"
	"fmt"
	"miopkg/errors"
	"miopkg/log"
	"redbook/app/domain/history/model/entity"
	"time"
)

const (
	_prefixUid           = "u_%d"
	_prefixBizType       = "bt_%s"
	_prefixUidAndBizType = "ubt_%d_%s"
	_prefixItemId        = "iid_%d"
	_prefixCreatedAt     = "ca_%d"
)

// func keyUid(uid int64) string {
// 	return fmt.Sprintf(_prefixUid, uid)
// }
// func keyBizType(bizType string) string {
// 	return fmt.Sprintf(_prefixBizType, bizType)
// }

func keyUidAndBizType(uid int64, bizType string) string {
	return fmt.Sprintf(_prefixUidAndBizType, uid, bizType)
}

func keyItemId(itemId int64) string {
	return fmt.Sprintf(_prefixItemId, itemId)
}
func keyCreatedAt(createdAt int64) string {
	return fmt.Sprintf(_prefixCreatedAt, createdAt)
}

func scanKeyUidAndBizType(key string) (uid int64, bizType string, err error) {
	_, err = fmt.Sscanf(key, _prefixUidAndBizType, &uid, &bizType)
	if err != nil {
		return 0, "", errors.Wrapf(errors.FromError(err), "scanKeyUidAndBizType: %s", key)
	}
	return
}

func scanItemId(key string) (itemId int64, err error) {
	_, err = fmt.Sscanf(key, _prefixItemId, &itemId)
	if err != nil {
		return 0, errors.Wrapf(errors.FromError(err), "scanItemId: %s", key)
	}
	return
}
func scanCreatedAt(key string) (createdAt int64, err error) {
	_, err = fmt.Sscanf(key, _prefixCreatedAt, &createdAt)
	if err != nil {
		return 0, errors.Wrapf(errors.FromError(err), "scanCreatedAt: %s", key)
	}
	return
}

func (r *historyRepo) CacheAddHistory(ctx context.Context, uid, itemId int64, bizType string) error {
	key := keyUidAndBizType(uid, bizType)
	err := r.data.historyRedis.HSet(ctx, key, keyItemId(itemId), keyCreatedAt(time.Now().Unix()))
	if err != nil {
		return errors.Wrapf(errors.FromError(err), "HSet: %s", key)
	}
	return nil
}

func (r *historyRepo) CacheListHistoryByTimestamp(ctx context.Context, key string, Timestamp int64) ([]*entity.History, error) {
	rv, err := r.data.historyRedis.HGetAll(ctx, key)
	if err != nil {
		log.Errorf("HGetAll: %s", err)
		return nil, err
	}
	entitys := make([]*entity.History, 0, len(rv))
	uid, bizType, err := scanKeyUidAndBizType(key)
	for itemIdStr, createdAtStr := range rv {
		history := &entity.History{}
		itemId, err := scanItemId(itemIdStr)
		createdAt, err := scanCreatedAt(createdAtStr)
		if err != nil {
			log.Errorf("scanItemId: %s", err)
			return nil, err
		}
		history.Uid = uid
		history.BizType = bizType
		history.ItemId = itemId
		history.CreatedAt = createdAt
		entitys = append(entitys, history)
	}
	return entitys, nil
}
