package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"

	"github.com/ysaito8015/go-todo-web-app/config"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminated server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	c := mysql.Config{
		User:      cfg.DBUser,
		Passwd:    cfg.DBPassword,
		Addr:      cfg.DBHost + ":" + string(rune(cfg.DBPort)),
		DBName:    cfg.DBName,
		ParseTime: true,
	}

	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 稼働待ち
	for {
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}
	log.Println("DB is ready")

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)

	mux, cleanup, err := NewMux(ctx, cfg)
	defer cleanup()
	if err != nil {
		return err
	}
	s := NewServer(l, mux)
	return s.Run(ctx)
}
