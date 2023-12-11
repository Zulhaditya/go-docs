package belajar_golang_database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

/*
- koneksi mysql di terminal :  mysql -u root -p
*/

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:Aygek6996!@tcp(localhost:3306)/golangdb")
	if err != nil {
		panic(err)
	}

	db.Close() // close database jika tidak digunakan lagi

	// gunakan database

}
