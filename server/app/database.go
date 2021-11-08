package app

import (
	"database/sql"
	"os"
	"time"
)

func NewDB() *sql.DB {

	NAME := os.Getenv("DB_NAME")
	PASSWORD := os.Getenv("DB_PASSWORD")

	DB, err := sql.Open("mysql", NAME+":"+PASSWORD+"@tcp(localhost:3306)/monim?parseTime=true")

	if err != nil {
		panic(err)
	}

	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)
	DB.SetConnMaxIdleTime(5 * time.Minute)
	DB.SetConnMaxLifetime(60 * time.Minute)

	return DB
}
