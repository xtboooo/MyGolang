/* package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db1 *sql.DB

func initDB1() (err error) {
	dsn := "root:123456@/go_db"
	db1, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db1.Ping()
	if err != nil {
		return err
	}
	return nil
}

func insert() {
	s := "insert into user_tb1 (username, password) values(?, ?)"
	r, err := db1.Exec(s, "a", "a1234")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		i, _ := r.LastInsertId()
		fmt.Printf("i: %v\n", i)
	}
}

func insert2(username string, password string) {
	s := "insert into user_tb1 (username, password) values(?, ?)"
	r, err := db1.Exec(s, username, password)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		i, _ := r.LastInsertId()
		fmt.Printf("i: %v\n", i)
	}
}

func main() {
	err := initDB1()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Println("连接成功!")
	}
	// insert()
	insert2("b", "b1234")
}
 */