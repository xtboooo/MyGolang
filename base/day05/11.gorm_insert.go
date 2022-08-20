/* package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db5 *gorm.DB

type Users struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

func init() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db5 = d
}

var user = Users{
	Name:     "bob",
	Age:      21,
	Birthday: time.Now(),
}

func createTable() {
	db5.AutoMigrate(&Users{})
}

func insert1() {
	result := db5.Create(&user)
	fmt.Printf("result.RowsAffected: %v\n", result.RowsAffected)
	fmt.Printf("user.ID: %v\n", user.ID)
}

func insert3() {
	db5.Select("Name", "Age", "CreatedAt").Create(&user)
}

func insert4() {
	var users = []Users{{Name: "a", Birthday: time.Now()}, {Name: "b", Birthday: time.Now()}, {Name: "c", Birthday: time.Now()}}
	db5.Create(&users)
}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeCreate...")
	return
}

func insert5() {
	db5.Model(&Users{}).Create(map[string]interface{}{"Name": "a", "Age": 20})
}

func main() {
	// createTable()
	// insert1()
	// insert3()
	// insert4()
	insert5()
}
 */