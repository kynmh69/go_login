package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_USER     = "DB_USER"
	DB_PASSWORD = "DB_PASSWORD"
	DB_SCHEMA   = "DB_SCHEMA"
	DB_PORT     = "DB_PORT"
	DB_HOST     = "DB_HOST"
	DB_SOFT     = "mysql"
)

var Db *sql.DB

func init() {
	ConnectDb()
}

func ConnectDb() {
	count := 10
	open(count)
}

func open(count int) {
	db_user := os.Getenv(DB_USER)
	db_password := os.Getenv(DB_PASSWORD)
	db_schema := os.Getenv(DB_SCHEMA)
	db_port := os.Getenv(DB_PORT)
	db_host := os.Getenv(DB_HOST)

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_user, db_password, db_host, db_port, db_schema)

	Db, err := sql.Open(DB_SOFT, dataSourceName)
	log.Println("Connected DB.", dataSourceName)

	if err != nil {
		log.Fatalln("DB connect error.", err.Error())
	}

	if err = Db.Ping(); err != nil {
		time.Sleep(time.Second * 2)
		count--
		log.Printf("retry... count:%d\n", count)
		open(count)
	}

	Db.SetConnMaxLifetime(time.Minute * 5)
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
}
