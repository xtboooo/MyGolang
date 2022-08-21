/* package main

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

func delete1() {
	var user User
	db.First(&user)
	db.Delete(&user) // 软删除
}

func delete2() {
	db.Delete(&User{}, 2)
}


func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Println("BeforeDelete")
	return
}

func main() {
	// delete1()
	delete2()
}
 */