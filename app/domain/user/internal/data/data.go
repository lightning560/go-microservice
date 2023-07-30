package data

import (
	"context"
	"fmt"
	"miopkg/client/eredis"
	"miopkg/log"
	"miopkg/sync/fanout"
	"redbook/app/domain/user/internal/data/ent"
	"runtime"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
)

var (
	_write = "root:password@tcp(127.0.0.1:13306)/account?timeout=2s&readTimeout=2s&writeTimeout=2s&parseTime=true&loc=Local&charset=utf8mb4,utf8"
	_read  = "root:password@tcp(127.0.0.1:13306)/account?timeout=2s&readTimeout=2s&writeTimeout=2s&parseTime=true&loc=Local&charset=utf8mb4,utf8"
)

type Data struct {
	userWDB *ent.Client
	userRDB *ent.Client
	// TODO: config expire time with conf
	rc     *eredis.Component
	fanout *fanout.Fanout
	logger *log.Logger
}

// TODO:config ent with conf
func NewEntClient(ctx context.Context, dsn string) *ent.Client {
	client, err := ent.Open(dialect.MySQL, "root:password@tcp(127.0.0.1:13306)/account?timeout=2s&readTimeout=2s&writeTimeout=2s&parseTime=true&loc=Local&charset=utf8mb4,utf8")
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

func NewRedisClient(ctx context.Context) *eredis.Component {
	return eredis.Load("redis.dev").Build()
}

func NewFanout() *fanout.Fanout {
	return fanout.New("cache", fanout.Worker(runtime.NumCPU()), fanout.Buffer(1024))
}

// TODO: 增加config，增加Data中的属性，读写分离、redis
func NewData(ctx context.Context) (*Data, func(), error) {
	logger := log.MioLogger.With(log.FieldMod("domain.user.data"))
	d := &Data{
		userWDB: NewEntClient(ctx, _write),
		rc:      NewRedisClient(ctx),
		fanout:  NewFanout(),
		logger:  logger,
	}
	return d, func() {
		if err := d.userWDB.Close(); err != nil {
			fmt.Println(err)
		}
	}, nil
}

// rollback calls to tx.Rollback and wraps the given error
// with the rollback error if occurred.
func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
