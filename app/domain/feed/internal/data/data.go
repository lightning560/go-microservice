package data

import (
	"context"
	"miopkg/client/eredis"
	"miopkg/db/emongo"
	"miopkg/log"
	"miopkg/sync/fanout"
	"runtime"
)

type Data struct {
	redisClient *eredis.Component
	feedMG      *emongo.Collection
	relationMG  *emongo.Collection
	likeMG      *emongo.Collection
	shareMG     *emongo.Collection
	favorMG     *emongo.Collection
	fanout      *fanout.Fanout
	logger      *log.Logger
}

func NewMongoClient(ctx context.Context, collection string) *emongo.Collection {
	// TODO:初始化多个emongo组件，改为1个
	mongo := emongo.Load("mongo.dev.feed").Build()
	return mongo.Client().Database("redbook").Collection(collection)
}

func NewRedisClient(ctx context.Context) *eredis.Component {
	return eredis.Load("redis.dev.feed").Build()
}

func NewFanout() *fanout.Fanout {
	return fanout.New("cache", fanout.Worker(runtime.NumCPU()), fanout.Buffer(1024))
}

func NewData(ctx context.Context) (*Data, func(), error) {
	logger := log.MioLogger.With(log.FieldMod("domain.feed.data"))
	d := &Data{
		// likeWDB: NewEntClient(ctx, _write),
		feedMG:      NewMongoClient(ctx, "feed"),
		relationMG:  NewMongoClient(ctx, "relation"),
		likeMG:      NewMongoClient(ctx, "like"),
		shareMG:     NewMongoClient(ctx, "share"),
		favorMG:     NewMongoClient(ctx, "favor"),
		redisClient: NewRedisClient(ctx),
		fanout:      NewFanout(),
		logger:      logger,
	}
	return d, func() { // close db conn
		// if err := d.likeWDB.Close(); err != nil {
		// 	fmt.Println(err)
		// }
	}, nil
}
