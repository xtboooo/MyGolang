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

func test1() {
	type CreditCard struct {
		gorm.Model
		Number string
		UserID uint
	}
	type User struct {
		gorm.Model
		CreditCard CreditCard
	}
	db.AutoMigrate(&User{}, &CreditCard{})
}

type Toy struct {
	ID        int
	Name      string
	OwnerID   int
	OwnerType string
}

type Cat struct {
	ID   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner;"`
}
type Dog struct {
	ID   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner;"`
}


func test2() {
	db.AutoMigrate(&Toy{}, &Cat{}, &Dog{})
}

func test3()  {
	// db.Create(&Dog{Name: "dog1", Toy: Toy{Name: "toy1"}})
	db.Create(&Cat{Name: "cat1", Toy: Toy{Name: "toy2"}})
}

func main() {
	// test1()
	// test2()
	test3()
}
 */