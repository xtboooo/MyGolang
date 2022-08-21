package main

import (
	"fmt"

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

type Address struct {
	gorm.Model
	Address1 string
	UserID   uint
}

type Email struct {
	gorm.Model
	Email  string
	UserID uint
}

type User struct {
	gorm.Model
	Name            string
	BillingAddress  Address
	ShippingAddress Address
	Emails          []Email
	Languages       []*Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name  string
	Users []*User `gorm:"many2many:user_languages;"`
}

func createTable() {
	db.AutoMigrate(&User{}, &Language{}, &Email{}, &Address{})
}

func insert() {
	user := User{
		Name:            "xtbo",
		BillingAddress:  Address{Address1: "BillingAddress - Address1"},
		ShippingAddress: Address{Address1: "ShippingAddress - Address1"},
		Emails: []Email{
			{Email: "xtbo9719@gmail.com"},
			{Email: "xtbo@outlook.com"},
		},
		Languages: []*Language{
			{Name: "EN"},
			{Name: "ZH"},
		},
	}
	// db.Create(&user)

	// If you want to update associations’s data, you should use the FullSaveAssociations mode
	db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)
}

// Find matched associations
func test() {
	var user User
	db.First(&user)
	var languages []Language
	db.Model(&user).Association("Languages").Find(&languages)
	fmt.Printf("languages: %v\n", languages)
}

// Return the count of current associations
func test2() {
	var user User
	db.First(&user)
	i := db.Model(&user).Association("Languages").Count()
	fmt.Printf("i: %v\n", i)
}

func main() {
	// createTable()
	// insert()
	// test()
	test2()
}
