package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	// 禁用全局事务
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{ /* SkipDefaultTransaction: true */ })
	if err != nil {
		panic("failed to connect database")
	}
	db = d
}

func test1() {
	// session级别的禁用事务
	db.Session(&gorm.Session{SkipDefaultTransaction: true})
}

type User struct {
	gorm.Model
	Name string
}

var user = User{
	Name: "bob",
}

// 没有使用事务操作
func test2() {
	db.Create(&user)
	db.Create(nil)
}

// 手动使用事务操作
func test3() {

	tx := db.Begin()
	tx.Create(&user)
	err := tx.Create(nil).Error
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}

// Transaction
func test4() {
	db.Transaction(func(tx *gorm.DB) error {
		// 这里会报主键重复的错误, 不会创建
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		return nil
	})
}

func main() {
	// test2()
	// test3()
	test4()
}
