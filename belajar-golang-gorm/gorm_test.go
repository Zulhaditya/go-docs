package belajar_golang_gorm

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func TestQueryAllObject(t *testing.T) {
	var users []User
	err := db.Find(&users, "id in ?", []string{"1", "2", "3", "4"}).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))
}

// implementasi query condition
func TestQueryCondition(t *testing.T) {
	var users []User
	result := db.Where("first_name like ?", "%User%").Where("password = ?", "secret").
		Find(&users)
	assert.Nil(t, result.Error)

	assert.Equal(t, 13, len(users))
}

// or operator
func TestOrOperator(t *testing.T) {
	var users []User
	result := db.Where("first_name like ?", "%User%").Or("password = ?", "secret").
		Find(&users)
	assert.Nil(t, result.Error)

	assert.Equal(t, 14, len(users))
}

// not operator
func TestNotOperator(t *testing.T) {
	var users []User
	result := db.Not("first_name like ?", "%User%").Where("password = ?", "secret").
		Find(&users)
	assert.Nil(t, result.Error)

	assert.Equal(t, 1, len(users))
}

// select fields tertentu
func TestSelectFields(t *testing.T) {
	var users []User
	err := db.Select("id", "first_name").Find(&users).Error
	assert.Nil(t, err)

	for _, user := range users {
		assert.NotNil(t, user.ID)
		assert.NotEqual(t, "", user.Name.FirstName)
	}

	assert.Equal(t, 14, len(users))
}

// struct condition
func TestStructCondition(t *testing.T) {
	userCondition := User{
		Name: Name{
			FirstName: "User 7",
			// LastName:  "", tidak bisa dilakukan karena merupakan nilai default value string
		},
	}

	var users []User
	result := db.Where(userCondition).Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 1, len(users))
}

// map condition
func TestMapCondition(t *testing.T) {
	mapCondition := map[string]interface{}{
		"middle_name": "",
	}

	var users []User
	result := db.Where(mapCondition).Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 13, len(users))
}

// order, limit dan offset
func TestOrderLimitOffset(t *testing.T) {
	var users []User
	result := db.Order("id asc, first_name asc").Limit(5).Offset(5).Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 5, len(users))
	assert.Equal(t, "14", users[0].ID)
}

type UserResponse struct {
	ID        string
	FirstName string
	LastName  string
}

// query non model
func TestQueryNonModel(t *testing.T) {
	var users []UserResponse
	err := db.Model(&User{}).Select("id", "first_name", "last_name").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 14, len(users))
	fmt.Println(users)
}

// method update
func TestUpdate(t *testing.T) {
	user := User{}
	result := db.First(&user, "id = ?", "1")
	assert.Nil(t, result.Error)

	// ubah data di struct user
	user.Name.FirstName = "Muhammad"
	user.Name.MiddleName = "Zulhaditya"
	user.Name.LastName = "Hapiz"
	user.Password = "password123"

	result = db.Save(&user)
	assert.Nil(t, result.Error)
}

// update selected column
func TestSelectedUpdate(t *testing.T) {
	// menggunakan map
	result := db.Model(&User{}).Where("id = ?", "1").Updates(map[string]interface{}{
		"middle_name": "",
		"last_name":   "Ackxle",
	})
	assert.Nil(t, result.Error)

	// secara langsung
	result = db.Model(&User{}).Where("id = ?", "1").Update("password", "secret")
	assert.Nil(t, result.Error)

	// menggunakan struct
	result = db.Where("id", "1").Updates(User{
		Name: Name{
			FirstName: "Inayah",
			LastName:  "Wulandari",
		},
	})

	assert.Nil(t, result.Error)
}

// test auto increment
func TestAutoIncrement(t *testing.T) {
	for i := 0; i < 10; i++ {
		userLog := UserLog{
			UserId: "1",
			Action: "Test Action",
		}

		result := db.Create(&userLog)
		assert.Nil(t, result.Error)

		assert.NotEqual(t, 0, userLog.ID)
		fmt.Println(userLog.ID)
	}
}

// method save bisa untuk create dan juga update
func TestSaveOrUpdate(t *testing.T) {
	userLog := UserLog{
		UserId: "1",
		Action: "Test Action",
	}

	// create
	result := db.Save(&userLog)
	assert.Nil(t, result.Error)

	// update
	userLog.UserId = "2"
	result = db.Save(&userLog)
	assert.Nil(t, result.Error)
}

// save non auto increment data
func TestSaveOrUpdateNonAutoIncrement(t *testing.T) {
	user := User{
		ID: "99",
		Name: Name{
			FirstName: "User 99",
		},
	}

	// create
	result := db.Save(&user)
	assert.Nil(t, result.Error)

	// update
	user.Name.FirstName = "User 99 Updated"
	result = db.Save(&user)
	assert.Nil(t, result.Error)
}

func TestConflict(t *testing.T) {
	user := User{
		ID: "69",
		Name: Name{
			FirstName: "User 69",
		},
	}

	result := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&user)
	assert.Nil(t, result.Error)
}

// method delete
func TestDelete(t *testing.T) {
	var user User
	result := db.First(&user, "id = ?", "99")
	assert.Nil(t, result.Error)
	result = db.Delete(&user)
	assert.Nil(t, result.Error)

	result = db.Delete(&User{}, "id = ?", "69")
	assert.Nil(t, result.Error)

	result = db.Where("id = ?", "11").Delete(&User{})
	assert.Nil(t, result.Error)
}

// method soft delete
func TestSoftDelete(t *testing.T) {
	todo := Todo{
		UserId:      "1",
		Title:       "Todo 1",
		Description: "Description 1",
	}

	err := db.Create(&todo).Error
	assert.Nil(t, err)

	err = db.Delete(&todo).Error
	assert.Nil(t, err)
	assert.NotNil(t, todo.DeletedAt)

	var todos []Todo
	err = db.Find(&todos).Error
	assert.Nil(t, err)
	assert.Equal(t, 0, len(todos))
}

// method unscoped akan melakukan hard delete ke database, baik yang soft delete sekalipun
func TestUnscoped(t *testing.T) {
	var todo Todo
	result := db.Unscoped().First(&todo, "id = ?", "1")
	assert.Nil(t, result.Error)

	result = db.Unscoped().Delete(&todo)
	assert.Nil(t, result.Error)

	var todos []Todo
	result = db.Unscoped().Find(&todos)
	assert.Nil(t, result.Error)
	assert.Equal(t, 0, len(todos))
}

// implementasi locking data pada transaction
func TestLock(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var user User
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&user, "id = ?", "1").Error
		if err != nil {
			return err
		}

		user.Name.FirstName = "Inayah"
		user.Name.MiddleName = "Fitri"
		user.Name.LastName = "Wulandari"
		return tx.Save(&user).Error
	})

	assert.Nil(t, err)
}

// relasi one to one

// test create wallet
func TestCreateWallet(t *testing.T) {
	wallet := Wallet{
		ID:      "1",
		UserId:  "1",
		Balance: 1000000,
	}

	err := db.Create(&wallet).Error
	assert.Nil(t, err)
}

// test retrieve relation untuk mengambil data wallet yang berelasi dengan table user
// dua kali query
func TestRetrieveRelationPreload(t *testing.T) {
	var user User
	err := db.Model(&User{}).Preload("Wallet").Take(&user).Error
	assert.Nil(t, err)

	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "1", user.Wallet.ID)
	fmt.Println(user.Wallet.ID)
	fmt.Println(user.Wallet.Balance)
}

// menggunakan joins
// satu kali query atau digabungkan (select kedua table sekaligus)
func TestRetrieveRelationJoin(t *testing.T) {
	var user User
	err := db.Model(&User{}).Joins("Wallet").Take(&user, "users.id = ?", "1").Error
	assert.Nil(t, err)

	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "1", user.Wallet.ID)
	assert.Equal(t, int64(1000000), user.Wallet.Balance)
}
