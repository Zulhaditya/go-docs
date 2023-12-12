package belajar_golang_database

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {

	db, err := sql.Open("mysql", "root:Aygek6996!@tcp(localhost:3306)/golangdb?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)                  // minimal connection
	db.SetMaxOpenConns(100)                 // maximal connection
	db.SetConnMaxIdleTime(5 * time.Minute)  // 5 menit idle lalu close connection
	db.SetConnMaxLifetime(60 * time.Minute) // connection apapun setelah 60 menit akan dibuatkan connection baru

	return db
}
