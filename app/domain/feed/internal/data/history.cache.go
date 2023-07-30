package data

import (
	"context"
	"redbook/app/domain/feed/internal/data/entity"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	_keyHistory = "his_" // uid -> bit(value)
	_bucket     = 10     // bit bucket
)

// KeyHistory return Switch key.
func keyHistory(uid int64) string {
	return _keyHistory + strconv.FormatInt(uid/_bucket, 10)
}

func (r *historyRepo) ListCacheHistoryItems(ctx context.Context, uid int64, offset, limit int32) (items *entity.HistoryItems, err error) {
	key := keyHistory(uid)
	now := time.Now()
	start := now.Add(-time.Hour * 24 * 60).Unix()
	stop := now.Unix()

	// 初始化查询条件， Offset和Count用于分页
	op := &redis.ZRangeBy{
		Min:    strconv.FormatInt(stop, 10),  // 最小分数
		Max:    strconv.FormatInt(start, 10), // 最大分数
		Offset: int64(offset),                // 类似sql的limit, 表示开始偏移量
		Count:  int64(limit),                 // 一次返回多少数据
	}
	rv, err := r.data.redisClient.ZRangeByScoreWithScores(ctx, key, op)
	if err != nil {
		return nil, err
	}
	for _, v := range rv {
		item := &entity.HistoryItem{}
		item.FromCache(&v)
		*items = append(*items, item)
	}
	return
}
