/* package main

import (
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

func test1() {
	type Company struct {
		gorm.Model
		Name string
		ID   int
	}

	type User struct {
		gorm.Model
		Name      string
		Age       int
		Birthday  time.Time
		Active    bool
		CompanyID int
		Company   Company
	}

	// many to one
	db.AutoMigrate(&Company{}, &User{})
}

func test2() {
	type Company struct {
		gorm.Model
		Name string
		ID   int
	}

	type User struct {
		gorm.Model
		Name         string
		Age          int
		Birthday     time.Time
		Active       bool
		CompanyRefer int
		Company      Company `gorm:"foreignKey:CompanyRefer"`
		// 使用外键 CompanyRefer
	}
	// many to one
	db.AutoMigrate(&Company{}, &User{})
}

func main() {
	// test1()
	test2()
}
 */