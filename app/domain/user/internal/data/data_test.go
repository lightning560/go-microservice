package data

import (
	"context"
	"fmt"
	"log"
	"testing"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"

	// _ "github.com/mattn/go-sqlite3"
	"github.com/smartystreets/goconvey/convey"

	"redbook/app/domain/user/internal/data/ent"
	"redbook/app/domain/user/internal/data/ent/user"
)

var (
	_mysql_local   = "root:password@tcp(127.0.0.1:13306)/account?timeout=2s&readTimeout=2s&writeTimeout=2s&parseTime=true&loc=Local&charset=utf8mb4,utf8"
	_sqlite3_local = "file:ent?mode=memory&cache=shared&_fk=1"
)

func TestCreateUser(t *testing.T) {
	convey.Convey("TestAccountData", t, func() {
		client, err := ent.Open(dialect.MySQL, _mysql_local)
		if err != nil {
			log.Fatalf("failed opening connection to sqlite: %v", err)
		}
		// Run the auto migration tool.
		if err := client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
		// Create a new account.
		a, err := client.User.Create().SetID(1).SetUID(101).SetPassword("000000").SetDeleted(false).SetEmail("101@main.xyz").Save(context.Background())
		if err != nil {
			log.Fatalf("failed creating account: %v", err)
		}
		fmt.Println(a)
	})
}
func TestReadUser(t *testing.T) {
	convey.Convey("TestAccountData", t, func() {
		client, err := ent.Open(dialect.MySQL, _mysql_local)
		if err != nil {
			log.Fatalf("failed opening connection to sqlite: %v", err)
		}
		u, err := client.Debug().User.Query().Where(user.UID(102)).Only(context.Background())
		if err != nil {
			log.Fatalf("can't query user by mid: %v", err)
			// return err
		}
		fmt.Println(u)
	})

}

func TestListUser(t *testing.T) {
	convey.Convey("TestListUser", t, func() {
		client, err := ent.Open(dialect.MySQL, _mysql_local)
		if err != nil {
			t.Fatalf("failed opening connection to sqlite: %v", err)
		}
		u, err := client.Debug().User.Query().Where(user.UIDIn(99, 100, 101)).All(context.Background())
		if err != nil {
			t.Fatalf("can't query user by mid: %v", err)
		}
		fmt.Println(u)
	})
}
