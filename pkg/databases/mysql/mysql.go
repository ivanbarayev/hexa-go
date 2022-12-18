package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"main/config"
)

var (
	ConnStr string
)

// NewMysqlDB Return new MysqlDB client
func NewMysqlDB(cfg *config.Config) (db *sql.DB, err error) {
	println("Driver MySQL Initialized")

	ConnStr = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.Mysql.USER, cfg.Mysql.PASS, cfg.Mysql.HOST, cfg.Mysql.PORT, cfg.Mysql.DEFAULT_DB)

	db, err = sql.Open("mysql", ConnStr)
	if err != nil {
		panic(err)
	} else {
		print("conn ok")
	}

	return
}
