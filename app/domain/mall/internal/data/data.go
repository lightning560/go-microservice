package data

import (
	"context"
	"miopkg/client/eredis"
	"miopkg/db/emongo"
	"miopkg/log"
	"miopkg/sync/fanout"
	"runtime"

	"redbook/app/domain/mall/internal/data/ent"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
)

var (
	_write = "root:password@tcp(127.0.0.1:13306)/account?timeout=2s&readTimeout=2s&writeTimeout=2s&parseTime=true&loc=Local&charset=utf8mb4,utf8"
	_read  = "root:password@tcp(127.0.0.1:13306)/account?timeout=2s&readTimeout=2s&writeTimeout=2s&parseTime=true&loc=Local&charset=utf8mb4,utf8"
)

type Data struct {
	redisClient  *eredis.Component
	collectionMG *emongo.Collection
	productMG    *emongo.Collection
	cartMG       *emongo.Collection
	shopMG       *emongo.Collection
	orderWDB     *ent.Client
	fanout       *fanout.Fanout
	logger       *log.Logger
}

// TODO:config ent with conf
func NewEntClient(ctx context.Context, dsn string) *ent.Client {
	client, err := ent.Open(dialect.MySQL, dsn)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	// defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func NewMongoCollection(ctx context.Context, collection string) *emongo.Collection {
	// TODO:初始化多个emongo组件，改为1个
	mongo := emongo.Load("mongo.dev.mall").Build()
	return mongo.Client().Database("mall").Collection(collection)
}

func NewRedisClient(ctx context.Context) *eredis.Component {
	return eredis.Load("redis.dev.mall").Build()
}

func NewFanout() *fanout.Fanout {
	return fanout.New("cache", fanout.Worker(runtime.NumCPU()), fanout.Buffer(1024))
}

func NewData(ctx context.Context) (*Data, func(), error) {
	logger := log.MioLogger.With(log.FieldMod("domain.mall.data"))
	d := &Data{
		collectionMG: NewMongoCollection(ctx, "collection"),
		productMG:    NewMongoCollection(ctx, "product"),
		cartMG:       NewMongoCollection(ctx, "cart"),
		shopMG:       NewMongoCollection(ctx, "shop"),
		orderWDB:     NewEntClient(ctx, _write),
		redisClient:  NewRedisClient(ctx),
		fanout:       NewFanout(),
		logger:       logger,
	}
	return d, func() { // close db conn
		// if err := d.likeWDB.Close(); err != nil {
		// 	fmt.Println(err)
		// }
	}, nil
}
