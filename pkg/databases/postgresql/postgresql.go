package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"main/config"
)

var (
	ConnStr string
)

const (
	maxOpenConns    = 250
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

// NewPostgresqlDB Return new Postgresql client
func NewPostgresqlDB(cfg *config.Config) (db *pgxpool.Pool, err error) {
	println("Driver PostgreSQL Initialized")
	ConnStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s pool_max_conns=%d", cfg.Postgresql.HOST, cfg.Postgresql.PORT, cfg.Postgresql.USER, cfg.Postgresql.PASS, cfg.Postgresql.DEFAULT_DB, cfg.Postgresql.MAX_CONN)

	db, err = pgxpool.Connect(context.Background(), ConnStr)
	if err != nil {
		println(err.Error())
	} else {
		print("conn ok")
	}

	if err = db.Ping(context.Background()); err != nil {
		println(err.Error())
	}

	return
}
