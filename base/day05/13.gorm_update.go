package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func init() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic("failed to connect database")
	}
	db = d
}

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
	Active   bool
}

func createTable() {
	db.AutoMigrate(&User{})
}

func insert() {
	user := User{
		Name:     "a",
		Age:      18,
		Birthday: time.Now(),
		Active:   true,
	}
	db.Create(&user)
}

func update() {
	var user User
	db.First(&user) // id=1
	user.Name = "b"
	user.Age = 20
	db.Save(&user)
}

func update2() {
	// 不需要查的操作
	db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
}

func update3() {
	var user User
	db.First(&user)
	// db.Model(&user).Updates(User{Name: "a", Age: 5, Active: false})
	db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": true})
}

// 回调 hook
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeUpdate")
	return
}
func main() {
	// createTable()
	// insert()
	// update()
	// update2()
	update3()
}
