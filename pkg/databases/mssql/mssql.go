package mssql

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"main/config"
)

var (
	connStr string
)

// NewMssqlDB Return new Ms SQL client
func NewMssqlDB(cfg config.Config) (db *sql.DB, err error) {
	println("Driver PostgreSQL Initialized")

	connStr = fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;port=%d", cfg.Mssql.HOST, cfg.Mssql.DEFAULT_DB, cfg.Mssql.USER, cfg.Mssql.PASS, cfg.Mssql.PORT)

	db, err = sql.Open("sqlserver", connStr)
	if err != nil {
		println(err.Error())
	}

	return
}
