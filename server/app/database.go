package app

import (
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	DB, err := sql.Open("mysql", "root:S@kha7800@tcp(localhost:3306)/monim?parseTime=true")

	if err != nil {
		panic(err)
	}

	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)
	DB.SetConnMaxIdleTime(5 * time.Minute)
	DB.SetConnMaxLifetime(60 * time.Minute)

	return DB
}
