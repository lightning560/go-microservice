package data

import (
	"context"
	"miopkg/client/eredis"
	"miopkg/db/emongo"
	"miopkg/log"
	"miopkg/sync/fanout"
	"runtime"
)

var (
	_write = "root:password@tcp(127.0.0.1:13306)/account?timeout=2s&readTimeout=2s&writeTimeout=2s&parseTime=true&loc=Local&charset=utf8mb4,utf8"
	_read  = "root:password@tcp(127.0.0.1:13306)/account?timeout=2s&readTimeout=2s&writeTimeout=2s&parseTime=true&loc=Local&charset=utf8mb4,utf8"
)

type Data struct {
	// likeWDB *ent.Client
	// likeRDB *ent.Client
	// TODO: config expire time with conf
	rc        *eredis.Component
	likeMG    *emongo.Collection
	dislikeMG *emongo.Collection
	fanout    *fanout.Fanout
	logger    *log.Logger
}

func NewMongoClient(ctx context.Context, collection string) *emongo.Collection {
	// TODO:初始化多个emongo组件，改为1个
	mongo := emongo.Load("mongo").Build()
	return mongo.Client().Database("redbook").Collection(collection)
}

func NewRedisClient(ctx context.Context) *eredis.Component {
	return eredis.Load("redis.dev").Build()
}

func NewFanout() *fanout.Fanout {
	return fanout.New("cache", fanout.Worker(runtime.NumCPU()), fanout.Buffer(1024))
}

func NewData(ctx context.Context) (*Data, func(), error) {
	d := &Data{
		// likeWDB: NewEntClient(ctx, _write),
		likeMG:    NewMongoClient(ctx, "like"),
		dislikeMG: NewMongoClient(ctx, "dislike"),
		rc:        NewRedisClient(ctx),
		fanout:    NewFanout(),
	}
	return d, func() { // close db conn
		// if err := d.likeWDB.Close(); err != nil {
		// 	fmt.Println(err)
		// }
	}, nil
}
