package clickhouse

import (
	"database/sql"
	"fmt"
	"github.com/ClickHouse/clickhouse-go"
	"log"
	"main/config"
)

var (
	connStr string
)

// NewClickHouseDB Return new Click House client
func NewClickHouseDB(cfg *config.Config) (db *sql.DB, err error) {
	connStr = fmt.Sprintf("tcp://%s:%s?username=%s&password=%s&database=%s&read_timeout=10&write_timeout=20&debug=true", cfg.Clickhouse.HOST, cfg.Clickhouse.PORT, cfg.Clickhouse.USER, cfg.Clickhouse.PASS, cfg.Clickhouse.DEFAULT_DB)

	db, err = sql.Open("clickhouse", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
	}

	return
}
