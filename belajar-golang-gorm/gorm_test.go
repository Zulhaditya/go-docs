package belajar_golang_gorm

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

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

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

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

// auto create/update pada relasi one to one
func TestAutoCreateUpdate(t *testing.T) {
	user := User{
		ID:       "20",
		Password: "secret",
		Name: Name{
			FirstName: "User 20 bro!",
		},
		Wallet: Wallet{
			ID:      "20",
			UserId:  "20",
			Balance: 2000000,
		},
	}

	err := db.Create(&user).Error
	assert.Nil(t, err)
}

// skip auto create/update menggunakan method omit
func TestSkipAutoCreateUpdate(t *testing.T) {
	user := User{
		ID:       "21",
		Password: "secret",
		Name: Name{
			FirstName: "User 21 bro!",
		},
		Wallet: Wallet{
			ID:      "21",
			UserId:  "21",
			Balance: 3000000,
		},
	}

	err := db.Omit(clause.Associations).Create(&user).Error
	assert.Nil(t, err)
}

// relasi one to many
func TestUserAndAddresses(t *testing.T) {
	user := User{
		ID:       "2",
		Password: "secret",
		Name: Name{
			FirstName: "User 50",
		},
		Wallet: Wallet{
			ID:      "2",
			UserId:  "2",
			Balance: 5000000,
		},
		Addresses: []Address{
			{
				UserId:  "2",
				Address: "Jalan kenangan",
			},
			{
				UserId:  "2",
				Address: "Jalan seroja",
			},
		},
	}

	err := db.Save(&user).Error
	assert.Nil(t, err)
}

// preload dan join pada relasi one to many
func TestPreloadJoinOneToMany(t *testing.T) {
	var usersPreload []User
	err := db.Model(&User{}).Preload("Addresses").Joins("Wallet").Find(&usersPreload).Error
	assert.Nil(t, err)
}

func TestTakePreloadJoinOneToMany(t *testing.T) {
	var user User
	err := db.Model(&User{}).Preload("Addresses").Joins("Wallet").
		Take(&user, "users.id = ?", "50").Error
	assert.Nil(t, err)
}

// preload dan join pada relasi many to one
func TestBelongsTo(t *testing.T) {
	fmt.Println("Preload")
	var addresses []Address
	err := db.Preload("User").Find(&addresses).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(addresses))

	fmt.Println("Joins")
	addresses = []Address{}
	err = db.Joins("User").Find(&addresses).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(addresses))
}

// many to one di relasi one to one
func TestBelongsToOneToOne(t *testing.T) {
	fmt.Println("Preload")
	var wallets []Wallet
	err := db.Preload("User").Find(&wallets).Error
	assert.Nil(t, err)

	fmt.Println("Joins")
	wallets = []Wallet{}
	err = db.Joins("User").Find(&wallets).Error
	assert.Nil(t, err)
}

// create many to many
func TestCreateManyToMany(t *testing.T) {
	// buat product terlebih dahulu
	product := Product{
		ID:    "P001",
		Name:  "Contoh Product",
		Price: 1000000,
	}

	err := db.Create(&product).Error
	assert.Nil(t, err)

	// user 1 menyukai product P001

	err = db.Table("user_like_product").Create(map[string]interface{}{
		"user_id":    "1",
		"product_id": "P001",
	}).Error
	assert.Nil(t, err)

	// user 2 menyukai product P001

	err = db.Table("user_like_product").Create(map[string]interface{}{
		"user_id":    "2",
		"product_id": "P001",
	}).Error
	assert.Nil(t, err)
}

// preload many to many
func TestPreloadManyToMany(t *testing.T) {
	var product Product
	err := db.Preload("LikedByUsers").First(&product, "id = ?", "P001").Error
	assert.Nil(t, err)
	assert.Equal(t, 2, len(product.LikedByUsers))
}

// mencari relasi menggunakan method association
func TestAssociationFind(t *testing.T) {
	var product Product
	err := db.First(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	var users []User
	err = db.Model(&product).Where("users.first_name LIKE ?", "User%").Association("LikedByUsers").
		Find(&users)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
}

// implementasi add relation
func TestAssociationAdd(t *testing.T) {
	var user User
	err := db.Take(&user, "id = ?", "3").Error
	assert.Nil(t, err)

	var product Product
	err = db.Take(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikedByUsers").Append(&user)
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikedByUsers").Append(&user)
	assert.Nil(t, err)
}

// implementasi replace relation
func TestAssociationReplace(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var user User
		err := tx.Take(&user, "id = ?", "1").Error
		assert.Nil(t, err)

		wallet := Wallet{
			ID:      "01",
			UserId:  "1",
			Balance: 2000000,
		}

		err = tx.Model(&user).Association("Wallet").Replace(&wallet)
		return err
	})

	assert.Nil(t, err)
}

// implementasi delete relation
func TestAssociationDelete(t *testing.T) {
	var user User
	err := db.Take(&user, "id = ?", "3").Error
	assert.Nil(t, err)

	var product Product
	err = db.Take(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikedByUsers").Delete(&user)
	assert.Nil(t, err)
}

// implementasi clear relation, hapus semua relasi
func TestAssociationClear(t *testing.T) {
	var product Product
	err := db.Take(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikedByUsers").Clear()
	assert.Nil(t, err)
}

// preloading with condition
func TestPreloadWithCondition(t *testing.T) {
	var user User
	err := db.Preload("Wallet", "balance > ?", 100).Take(&user, "id = ?", "1").Error
	assert.Nil(t, err)
	fmt.Println(user) // data wallet kosong
}

// nested preloading
func TestNestedPreloading(t *testing.T) {
	var wallet Wallet
	err := db.Preload("User.Addresses").Take(&wallet, "id = ?", "2").Error
	assert.Nil(t, err)
	fmt.Println(wallet)
	fmt.Println(wallet.User)
	fmt.Println(wallet.User.Addresses)
}

// preload all
func TestPreloadAll(t *testing.T) {
	var user User
	err := db.Preload(clause.Associations).Take(&user, "id = ?", "1").Error
	assert.Nil(t, err)
}

// implementasi join query
func TestJoinQuery(t *testing.T) {
	var users []User
	// inner join => data harus wajib ada di kedua table
	err := db.Joins("join wallets on wallets.user_id = users.id").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))

	users = []User{}
	err = db.Joins("Wallet").Find(&users).Error // menggunakan left join => data hanya ada di satu table saja
	assert.Nil(t, err)
	assert.Equal(t, 16, len(users))
}

func TestJoinQueryCondition(t *testing.T) {
	var users []User
	// inner join => data harus wajib ada di kedua table
	err := db.Joins("join wallets on wallets.user_id = users.id AND wallets.balance > ?", 500000).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))

	users = []User{}
	err = db.Joins("Wallet").Where("Wallet.balance > ?", 500000).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))
}

// implementasi count aggregation
func TestAggregationCount(t *testing.T) {
	var count int64
	err := db.Model(&User{}).Joins("Wallet").Where("Wallet.balance > ?", 500000).
		Count(&count).Error
	assert.Nil(t, err)
	assert.Equal(t, int64(4), count)
}

// aggregation yang lain
type AggregationResult struct {
	TotalBalance int64
	MinBalance   int64
	MaxBalance   int64
	AvgBalance   float64
}

func TestAggregation(t *testing.T) {
	var result AggregationResult
	err := db.Model(&Wallet{}).Select("sum(balance) as total_balance", "min(balance) as min_balance",
		"max(balance) as max_balance", "avg(balance) as avg_balance").Take(&result).Error

	assert.Nil(t, err)
	assert.Equal(t, int64(13000000), result.TotalBalance)
	assert.Equal(t, int64(1000000), result.MinBalance)
	assert.Equal(t, int64(5000000), result.MaxBalance)
	assert.Equal(t, float64(3250000.0000), result.AvgBalance)
}

// implementasi aggregation having dan group by
func TestAggregationGroupByHaving(t *testing.T) {
	var result []AggregationResult
	err := db.Model(&Wallet{}).Select("sum(balance) as total_balance", "min(balance) as min_balance",
		"max(balance) as max_balance", "avg(balance) as avg_balance").Joins("User").Group("User.id").
		Having("sum(balance) > ?", 1000000).Find(&result).Error

	assert.Nil(t, err)
	assert.Equal(t, 3, len(result)) // 3 data dengan wallet diatas satu juta
}

// implementasi context
func TestContext(t *testing.T) {
	ctx := context.Background()

	var users []User
	err := db.WithContext(ctx).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 16, len(users))
}

// implementasi scopes
func BrokenWalletBalance(db *gorm.DB) *gorm.DB {
	return db.Where("balance = ?", 0)
}

func SultanWalletBalance(db *gorm.DB) *gorm.DB {
	return db.Where("balance > ?", 1000000)
}

func TestScopes(t *testing.T) {
	var wallets []Wallet
	err := db.Scopes(BrokenWalletBalance).Find(&wallets).Error
	assert.Nil(t, err)

	wallets = []Wallet{}
	err = db.Scopes(SultanWalletBalance).Find(&wallets).Error
	assert.Nil(t, err)
}

// implementasi migrator untuk membuat table baru
func TestMigrator(t *testing.T) {
	err := db.Migrator().AutoMigrate(&GuestBook{})
	assert.Nil(t, err)
}

// implementasi hook

/*
1. hook untuk create
- begin transaction
BeforeSave()
BeforeCreate()
- save before associations
- insert into database
- save after associations
AfterCreate()
AfterSave()
- commit or rollback transaction
*/

/*
2. hook untuk update
- begin transaction
BeforeSave()
BeforeUpdate()
- save before associations
- insert into database
- save after associations
AfterUpdate()
AfterSave()
- commit or rollback transaction
*/

/*
3. hook untuk delete
- begin transaction
BeforeDelete()
- delete from database
AfterDelete()
- commit or rollback transaction
*/

/*
4. hook untuk find
- load data from database
- preloading (eager loading)
AfterFind()
*/

// implementasi before create

// function untuk menambahkan id jika id user kosong
func (u *User) BeforeCreate(db *gorm.DB) error {
	if u.ID == "" {
		u.ID = "user-" + time.Now().Format("20060102150405")
	}

	return nil
}

func TestUserHook(t *testing.T) {
	user := User{
		Password: "secret",
		Name: Name{
			FirstName: "User 100",
		},
	}

	err := db.Create(&user).Error
	assert.Nil(t, err)
	assert.NotNil(t, user.ID)
	assert.NotEqual(t, "", user.ID)
}
