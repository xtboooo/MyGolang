/* package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db6 *gorm.DB
func init() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db6 = d
}

type Users struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

var user = Users{
	Name:     "bob",
	Age:      21,
	Birthday: time.Now(),
}


func select1()  {
	// db6.First(&user)
	// fmt.Printf("user.ID: %v\n", user.ID)
	// db6.Take(&user)
	// fmt.Printf("user.ID: %v\n", user.ID)
	db6.Last(&user)
	fmt.Printf("user.ID: %v\n", user.ID)
}


func select2()  {
	// db6.First(&user, 12)
	// fmt.Printf("user.ID: %v\n", user.ID)
	// db6.First(&user, "12")
	// fmt.Printf("user.ID: %v\n", user.ID)
	var users []Users
	db6.Find(&users, []int{11, 12})
	for _, user := range users {
		fmt.Printf("user.ID: %v\n", user.ID)
	}
}



func select3()  {
	var users []Users
	result:=db6.Find(&users)
	fmt.Printf("result.RowsAffected: %v\n", result.RowsAffected)
}

func main()  {
	// select1()
	// select2()
	select3()
}
 */