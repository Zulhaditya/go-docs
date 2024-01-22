package belajar_golang_gorm

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenConnection() *gorm.DB {
	driver := mysql.Open("root:mysqlpass123@tcp(127.0.0.1:3306)/belajar_golang_gorm?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(driver, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // menambahkan logger mode info
	})
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

// test create user
func TestCreateUser(t *testing.T) {
	user := User{
		ID:       "1",
		Password: "secret",
		Name: Name{
			FirstName:  "Inayah",
			MiddleName: "Fitri",
			LastName:   "Wulandari",
		},
		Information: "Pesan ini akan di ignore",
	}

	response := db.Create(&user)
	assert.Nil(t, response.Error)
	assert.Equal(t, 1, int(response.RowsAffected))
}

// test batch insert => insert data multiple sekaligus dalam satu query
func TestBatchInsert(t *testing.T) {
	var users []User
	for i := 2; i < 10; i++ {
		users = append(users, User{
			ID: strconv.Itoa(i),
			Name: Name{
				FirstName: "User " + strconv.Itoa(i),
			},
			Password: "secret",
		})
	}

	result := db.Create(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 8, int(result.RowsAffected))

}

// implementasi transaction
func TestTransactionSuccess(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{ID: "10", Password: "secret", Name: Name{FirstName: "User 10"}}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&User{ID: "11", Password: "secret", Name: Name{FirstName: "User 11"}}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&User{ID: "12", Password: "secret", Name: Name{FirstName: "User 12"}}).Error
		if err != nil {
			return err
		}

		return nil
	})

	assert.Nil(t, err)
}

func TestTransactionError(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{ID: "13", Password: "secret", Name: Name{FirstName: "User 10"}}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&User{ID: "11", Password: "secret", Name: Name{FirstName: "User 11"}}).Error
		if err != nil {
			return err
		}

		return nil
	})

	assert.Nil(t, err)
}

// implementasi manual transaction
func TestManualTransactionSuccess(t *testing.T) {
	tx := db.Begin()
	defer tx.Rollback()

	err := tx.Create(&User{ID: "13", Password: "secret", Name: Name{FirstName: "User 13"}}).Error
	assert.Nil(t, err)

	err = tx.Create(&User{ID: "14", Password: "secret", Name: Name{FirstName: "User 14"}}).Error
	assert.Nil(t, err)

	if err == nil {
		tx.Commit()
	}
}

func TestManualTransactionError(t *testing.T) {
	tx := db.Begin()
	defer tx.Rollback()

	err := tx.Create(&User{ID: "15", Password: "secret", Name: Name{FirstName: "User 15"}}).Error
	assert.Nil(t, err)

	err = tx.Create(&User{ID: "11", Password: "secret", Name: Name{FirstName: "User 11"}}).Error
	assert.Nil(t, err)

	if err == nil {
		tx.Commit()
	}
}

// implementasi single query object
func TestQuerySingleObject(t *testing.T) {
	user := User{}
	result := db.First(&user)
	assert.Nil(t, result.Error)
	assert.Equal(t, "1", user.ID)

	user = User{}
	result = db.Last(&user)
	assert.Nil(t, result.Error)
	assert.Equal(t, "9", user.ID)
}

func TestQueryInlineCondition(t *testing.T) {
	user := User{}
	result := db.Take(&user, "id = ?", "5")
	assert.Nil(t, result.Error)
	assert.Equal(t, "5", user.ID)
	assert.Equal(t, "User 5", user.Name.FirstName)
}
