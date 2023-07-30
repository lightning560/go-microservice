package data

import (
	"context"
	"fmt"
	"miopkg/client/eredis"
	"miopkg/db/emongo"
	"miopkg/log"
	"miopkg/sync/fanout"
	"runtime"
)

type Data struct {
	historyCollection *emongo.Collection
	historyRedis      *eredis.Component
	logger            *log.Logger
}

func NewMongoClient(ctx context.Context, collection string) *emongo.Collection {
	mongo := emongo.Load("mongo.dev.history").Build()
	return mongo.Client().Database("redbook").Collection(collection)
}

func NewRedisClient(ctx context.Context) *eredis.Component {
	return eredis.Load("redis.dev.history").Build()
}

func NewFanout() *fanout.Fanout {
	return fanout.New("cache", fanout.Worker(runtime.NumCPU()), fanout.Buffer(1024))
}

func NewData(ctx context.Context) (*Data, func(), error) {
	logger := log.MioLogger.With(log.FieldMod("domain.history.data"))
	d := &Data{
		historyCollection: NewMongoClient(ctx, "history"),
		historyRedis:      NewRedisClient(ctx),
		logger:            logger,
	}
	return d, func() { // close db conn
		if err := d.historyCollection.Drop(ctx); err != nil {
			fmt.Println(err)
		}
	}, nil
}
