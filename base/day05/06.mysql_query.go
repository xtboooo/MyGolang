/* package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db2 *sql.DB

func initDB2() (err error) {
	dsn := "root:123456@/go_db"
	db2, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db2.Ping()
	if err != nil {
		return err
	}
	return nil
}

type User struct {
	id       int
	username string
	password string
}

// 查单行
func queryOneRow() {
	s := "select * from user_tb1 where id = ?"
	var user User
	err := db2.QueryRow(s, 2).Scan(&user.id, &user.username, &user.password)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("user: %v\n", user)
	}
}

// 查多行
func queryManyRow() {
	s := "select * from user_tb1"
	r, err := db2.Query(s)
	defer r.Close()
	var user User
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		for r.Next() {
			r.Scan(&user.id, &user.username, &user.password)
			fmt.Printf("user: %v\n", user)
		}
	}
}

func main() {
	err := initDB2()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Println("连接成功!")
	}
	// queryOneRow()
	queryManyRow()
}
 */