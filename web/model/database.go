package model

import (
	"database/sql"
	"fmt"
	"go_login/logging"
	"os"
	"time"
)

const (
	DB_USER     = "DB_USER"
	DB_PASSWORD = "DB_PASSWORD"
	DB_SCHEMA   = "DB_SCHEMA"
	DB_PORT     = "DB_PORT"
	DB_HOST     = "DB_HOST"
	DB_SOFT     = "mysql"
)

func Open() *sql.DB {
	logger := logging.GetLogger()

	db_user := os.Getenv(DB_USER)
	db_password := os.Getenv(DB_PASSWORD)
	db_schema := os.Getenv(DB_SCHEMA)
	db_port := os.Getenv(DB_PORT)
	db_host := os.Getenv(DB_HOST)

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db_user, db_password, db_host, db_port, db_schema)

	db, err := sql.Open(DB_SOFT, dataSourceName)
	logger.Debug("Connected DB.", dataSourceName)

	if err != nil {
		logger.Fatalln("DB connect error.", err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
