package data

import (
	"context"
	"miopkg/conf"
	"strings"
	"testing"
	"time"

	"github.com/BurntSushi/toml"
)

func _newData(ctx context.Context) (*Data, func(), error) {
	d := &Data{
		historyCollection: NewMongoClient(ctx, "history"),
		historyRedis:      NewRedisClient(ctx),
	}
	return d, func() { // close db conn
		// if err := d.likeWDB.Close(); err != nil {
		// 	fmt.Println(err)
		// }
	}, nil
}

func loadConfig(t *testing.T) {
	var config = `
	[mongo.dev.history]
		debug=true
		dsn="mongodb://root:password@127.0.0.1:27017/?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&ssl=false"
	[redis.dev.history]
		debug = true # mio增加redis debug,打开后可以看到,配置名、地址、耗时、请求数据、响应数据
		addr = "127.0.0.1:6379"
		enableAccessInterceptor = true
		enableAccessInterceptorReq = true
		enableAccessInterceptorRes = true
	`
	err := conf.LoadFromReader(strings.NewReader(config), toml.Unmarshal)
	if err != nil {
		t.Fatal(err)
	}
}
func addPost(t *testing.T, uid, itemId int64, bizType string) {
	loadConfig(t)
	data, _, err := _newData(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	repo := NewHistoryRepo(data)

	err = repo.AddHistory(context.Background(), uid, itemId, bizType)
	if err != nil {
		t.Fatal(err)
	}
}
func TestAddHistory(t *testing.T) {
	loadConfig(t)
	data, _, err := _newData(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	repo := NewHistoryRepo(data)
	err = repo.AddHistory(context.Background(), 101, 1, "video")
	time.Sleep(20 * time.Second)
	if err != nil {
		t.Fatal(err)
	}
}
