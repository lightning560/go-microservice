package data

import (
	"context"
	"fmt"
	"miopkg/log"
	"miopkg/util/xgo"
	"redbook/app/job/history/server/biz"
	"sync"
	"time"
)

type historyRepo struct {
	data           *Data
	notify         chan string
	rwLock         sync.RWMutex
	threshold      int32
	historyCounter map[string]int32
	// historyCounterSync sync.Map
}

func NewHistoryRepo(data *Data) biz.IHistoryRepo {
	historyCounter := make(map[string]int32)
	repo := &historyRepo{
		data:           data,
		threshold:      1,
		notify:         make(chan string, 10),
		historyCounter: historyCounter,
	}
	xgo.Go(
		repo.historyproc,
	)
	xgo.Go(
		repo.historyTiktok,
	)
	return repo
}

func (r *historyRepo) historyTiktok() {
	fmt.Println("start historyTiktok")
	for {
		fmt.Println("historyTiktok")
		select {
		case <-r.notify:
			fmt.Println("notify")
			// case <-time.After(time.Second * 10):
			// 	return
		}
		time.Sleep(time.Second * 2)
	}
}

func (r *historyRepo) historyproc() {
	fmt.Println("start historyproc")
	for key := range r.notify {
		fmt.Println("historyproc", key)
		uid, bizType, err := scanKeyUidAndBizType(key)
		if err != nil {
			log.Errorf("scanKeyUidAndBizType: %s", err)
			continue
		}
		r.FlushHistory(context.Background(), uid, bizType)
	}
}

func (r *historyRepo) AddHistory(ctx context.Context, itemId, uid int64, bizType string) error {
	key := keyUidAndBizType(uid, bizType)
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	err := r.CacheAddHistory(ctx, uid, itemId, bizType)
	if err != nil {
		return err
	}
	r.historyCounter[key]++
	if r.historyCounter[key] >= r.threshold {
		r.notify <- key
		fmt.Println("threshold notify", key)
		r.historyCounter[key] = 0
	}
	time.Sleep(time.Millisecond * 60)
	return nil
}

func (r *historyRepo) FlushHistory(ctx context.Context, uid int64, bizType string) error {
	key := keyUidAndBizType(uid, bizType)
	entities, err := r.CacheListHistoryByTimestamp(ctx, key, 0)
	if err != nil {
		log.Errorf("CacheListGetHistory: %s", err)
		return err
	}
	err = r.RawBatchAddHistory(ctx, entities)
	if err != nil {
		log.Errorf("RawBatchAddHistory: %s", err)
		return err
	}
	return nil
}
