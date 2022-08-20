/* package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db3 *sql.DB

func initDB3() (err error) {
	dsn := "root:123456@/go_db"
	db3, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db3.Ping()
	if err != nil {
		return err
	}
	return nil
}

func update()  {
	s := "update user_tb1 set username=?, password=?, where id=?"
	r, err := db3.Exec(s, "bob", "bob", 1)
	defer db3.Close()
	if err!= nil{
		fmt.Printf("err: %v\n", err)
	}else{
		i, _ := r.RowsAffected()
		fmt.Printf("i: %v\n", i)
	}
}


func main() {
	err := initDB3()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Println("连接成功!")
	}
	update()
}
 */