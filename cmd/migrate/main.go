package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jus1d/kypidbot/internal/config"
	"github.com/pressly/goose/v3"
)

func main() {
	c := config.MustLoad()
	p := c.Postgres
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", p.User, p.Password, p.Host, p.Port, p.Name, p.ModeSSL)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("set dialect: %v", err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatalf("up: %v", err)
	}
}
