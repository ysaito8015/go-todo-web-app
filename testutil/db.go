package testutil

import (
	"database/sql"
	"os"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/ysaito8015/go-todo-web-app/config"
)

func OpenDBForTest(t *testing.T) *sqlx.DB {
	t.Helper()

	cfg, err := config.New()
	if err != nil {
		t.Fatal(err)
	}

	port := cfg.DBPort
	if _, defined := os.LookupEnv("CI"); defined {
		port = 3306
	}

	c := mysql.Config{
		User:      cfg.DBUser,
		Passwd:    cfg.DBPassword,
		Addr:      cfg.DBHost + ":" + string(rune(port)),
		DBName:    cfg.DBName,
		ParseTime: true,
	}

	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { _ = db.Close() })

	return sqlx.NewDb(db, "mysql")
}
