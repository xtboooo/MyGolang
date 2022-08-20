/* package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func testMysqlDriver() {
	db, err := sql.Open("mysql", "root:123456@/go_db")
	if err != nil {
		panic(err)
	} else {
		print("connect success!!!")
	}
	// See "Important settings" section.
	// 最大连接时长
	db.SetConnMaxLifetime(time.Minute * 3)
	// 最大连接数
	db.SetMaxOpenConns(10)
	// 空闲连接数
	db.SetMaxIdleConns(10)
}

var db *sql.DB

func initDB() (err error) {
	dsn := "root:123456@/go_db"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// testMysqlDriver()
	err := initDB()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Println("连接成功!")
	}
}
 */