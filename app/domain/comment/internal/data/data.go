package data

import (
	"context"
	"miopkg/client/eredis"
	"miopkg/db/emongo"
	xlog "miopkg/log"
	"miopkg/sync/fanout"
	"runtime"
)

type Data struct {
	replyMG     *emongo.Collection
	likeMG      *emongo.Collection
	subjectMG   *emongo.Collection
	redisClient *eredis.Component
	fanout      *fanout.Fanout
	logger      *xlog.Logger
}

func NewMongoClient(ctx context.Context, databaseName string) *emongo.Database {
	mongo := emongo.Load("mongo.dev.comment").Build()
	return mongo.Client().Database(databaseName)
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

func NewData(ctx context.Context, replyMG *emongo.Collection, subjectMG *emongo.Collection, likeMG *emongo.Collection, redisClient *eredis.Component, fanout *fanout.Fanout, logger *xlog.Logger) (*Data, func(), error) {

	d := &Data{
		replyMG:     replyMG,
		subjectMG:   subjectMG,
		likeMG:      likeMG,
		redisClient: redisClient,
		fanout:      fanout,
		logger:      logger,
	}
	return d, func() {
		// close db conn
		if err := replyMG.Drop(ctx); err != nil {
			d.logger.Error(err.Error() + "comment data drop db error")
		}
	}, nil
}
