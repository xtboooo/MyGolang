/* package main

import (
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
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
}

func createTable() {
	db.AutoMigrate(&User{}, &Language{})
}

func insert() {
	l1 := Language{
		Name: "english",
	}
	l2 := Language{
		Name: "chinese",
	}
	user := User{
		Languages: []Language{l1, l2},
	}
	db.Create(&user)
}

func main() {
	// createTable()
	insert()
}
 */