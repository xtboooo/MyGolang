package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db4 *sql.DB

func initDB4() (err error) {
	dsn := "root:123456@/go_db"
	db4, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db4.Ping()
	if err != nil {
		return err
	}
	return nil
}

func delete() {
	s := "delete from user_tb1 where id=?"
	r, err := db4.Exec(s, 1)
	defer db4.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		i, _ := r.RowsAffected()
		fmt.Printf("i: %v\n", i)
	}
}

func main() {
	err := initDB4()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Println("连接成功!")
	}
	delete()
}
