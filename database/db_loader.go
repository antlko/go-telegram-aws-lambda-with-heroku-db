package database

import (
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	defaultIdle = 2
	defaultOpen = 2
)

func NewPostgresDB() *sqlx.DB {
	idle, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	if idle == 0 {
		idle = defaultIdle
	}
	open, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
	if open == 0 {
		open = defaultOpen
	}

	config, err := getPostgresDBConfig()
	if err != nil {
		log.Fatal(err)
	}

	db := stdlib.OpenDB(*config)
	ddx := sqlx.NewDb(db, "pgx")
	ddx.SetMaxOpenConns(open)
	ddx.SetMaxIdleConns(idle)
	ddx.SetConnMaxLifetime(5 * time.Minute)

	err = ddx.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return ddx.Unsafe()
}

func getPostgresDBConfig() (*pgx.ConnConfig, error) {
	isHerokuConn, _ := strconv.Atoi(os.Getenv("HEROKU_IS_ACTIVE"))
	if isHerokuConn == 1 {
		return getHerokuConfigURLByPostgresID()
	}
	return getSimplePostgresConfig()
}

func getSimplePostgresConfig() (*pgx.ConnConfig, error) {
	var (
		host     = os.Getenv("PG_HOST")
		port     = os.Getenv("PG_PORT")
		user     = os.Getenv("PG_USER")
		password = os.Getenv("PG_PASSWORD")
		dbname   = os.Getenv("PG_NAME")
	)

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	return pgx.ParseConfig(dsn)
}
