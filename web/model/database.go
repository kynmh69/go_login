package model

import (
	"database/sql"
	"fmt"
	"go_login/utils"
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

func ConnectDb() *sql.DB {
	utils.LoadEnv(".env")
	count := 10
	db := open(count)
	return db
}

func open(count int) *sql.DB {
	db_user := os.Getenv(DB_USER)
	db_password := os.Getenv(DB_PASSWORD)
	db_schema := os.Getenv(DB_SCHEMA)
	db_port := os.Getenv(DB_PORT)
	db_host := os.Getenv(DB_HOST)

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_user, db_password, db_host, db_port, db_schema)

	log.Println("connecting to", dataSourceName)

	db, err := sql.Open(DB_SOFT, dataSourceName)

	if err != nil {
		log.Fatalln("DB connect error.", err.Error())
	}

	if err = db.Ping(); err != nil {
		if count == 0 {
			log.Fatalln("can't connect db.")
		}
		time.Sleep(time.Second * 5)
		count--
		log.Printf("retry... count:%d\n", count)
		open(count)
	}

	log.Println("Connected DB.", dataSourceName)

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
