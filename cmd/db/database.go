package db

import (
	"fmt"
	"log"
	"net/url"

	"github.com/sapawarga/auth/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewConnection(config *config.Config) *sqlx.DB {
	var err error
	connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_PORT, config.DB_NAME)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	db, err := sqlx.Connect(config.DB_DRIVER_NAME, dsn)
	if err != nil {
		log.Panic("[CONFIG DB] error connect :", err)
	}

	if err = db.Ping(); err != nil {
		log.Panic("[CONFIG DB] error ping :", err)
	}

	return db
}
