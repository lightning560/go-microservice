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
	replyMG *emongo.Collection
	likeMG  *emongo.Collection
	logger  *log.Logger
}

func NewMongoClient(ctx context.Context, database string) *emongo.Database {
	mongo := emongo.Load("mongo.dev.comment").Build()
	return mongo.Client().Database(database)
}

func NewMongoCollection(ctx context.Context, db *emongo.Database, collection string) *emongo.Collection {
	return db.Collection(collection)
}
func NewRedisClient(ctx context.Context) *eredis.Component {
	return eredis.Load("redis.dev.comment").Build()
}

func NewFanout() *fanout.Fanout {
	return fanout.New("cache", fanout.Worker(runtime.NumCPU()), fanout.Buffer(1024))
}

func NewData(ctx context.Context, replyMG *emongo.Collection, likeMG *emongo.Collection, logger *log.Logger) (*Data, func(), error) {
	d := &Data{
		replyMG: replyMG,
		likeMG:  likeMG,
		logger:  logger,
	}
	return d, func() { // close db conn
		if err := d.replyMG.Drop(ctx); err != nil {
			fmt.Println(err)
		}
		if err := d.likeMG.Drop(ctx); err != nil {
			fmt.Println(err)
		}
	}, nil
}
