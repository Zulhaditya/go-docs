package belajar_golang_gorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	driver := mysql.Open("root:mysqlpass123@tcp(127.0.0.1:3306)/belajar_golang_gorm?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(driver, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

var db = OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, db)
}

// insert execute sql
func TestExecuteSQL(t *testing.T) {
	err := db.Exec("INSERT INTO sample(id, name) values (?, ?)", "1", "Inayah").Error
	assert.Nil(t, err)

	err = db.Exec("INSERT INTO sample(id, name) values (?, ?)", "2", "Fitri").Error
	assert.Nil(t, err)

	err = db.Exec("INSERT INTO sample(id, name) values (?, ?)", "3", "Wulandari").Error
	assert.Nil(t, err)
}

type Sample struct {
	ID   string
	Name string
}

// select execute sql
func TestRawSQL(t *testing.T) {
	var sample Sample
	err := db.Raw("SELECT id, name FROM sample WHERE id = ?", "1").Scan(&sample).Error
	assert.Nil(t, err)
	assert.Equal(t, "1", sample.ID)
	assert.Equal(t, "Inayah", sample.Name)

	var samples []Sample
	err = db.Raw("SELECT id, name FROM sample").Scan(&samples).Error
	assert.Nil(t, err)
	assert.Equal(t, 3, len(samples))

}

// implementasi sql row secara manual
func TestSQLRow(t *testing.T) {
	rows, err := db.Raw("SELECT id, name FROM sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	// iterasi data
	var samples []Sample
	for rows.Next() {
		var id string
		var name string

		err := rows.Scan(&id, &name)
		assert.Nil(t, err)

		samples = append(samples, Sample{
			ID:   id,
			Name: name,
		})
	}

	assert.Equal(t, 3, len(samples))
}

// implementasi scan rows menggunakan gorm secara otomatis
func TestScanRows(t *testing.T) {
	var samples []Sample

	rows, err := db.Raw("SELECT id, name FROM sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	for rows.Next() {
		err := db.ScanRows(rows, &samples)
		assert.Nil(t, err)
	}

	assert.Equal(t, 3, len(samples))

}
