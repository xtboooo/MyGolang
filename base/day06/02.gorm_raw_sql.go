/* package main

import (
	"database/sql"
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

func testRawSQL1() {

	var user User
	db.Raw("select id, name, agr from users where name=?", "hello").Scan(&user)
	fmt.Printf("user: %v\n", user)

	var age int
	db.Raw("select sum(age) from users").Scan(&age)
	fmt.Printf("age: %v\n", age)
}

func testRawSQL2() {
	db.Exec("update users set age=? where name=?", 100, "a")
}

func testRawSQL3() {
	var user User
	db.Where("name = @myname", sql.Named("myname", "a")).Find(&user)
	fmt.Printf("user: %v\n", user)
}

func testRawSQL4() {
	var name string
	var age int
	row := db.Table("users").Where("name = ?", "a").Select("name", "age").Row()
	row.Scan(&name, &age)
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)

	rows, _ := db.Model(&User{}).Where("age > ?", 10).Select("name, age, active").Rows()

	for rows.Next(){
		rows.Scan(&name, &age)
		fmt.Printf("name: %v\n", name)
		fmt.Printf("age: %v\n", age)
	}
}

func main() {
	// testRawSQL1()
	// testRawSQL2()
	// testRawSQL3()
	testRawSQL4()
}
 */