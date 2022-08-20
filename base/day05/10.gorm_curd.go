package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price int
}

// 创建表
func create_db(db *gorm.DB) {
	db.AutoMigrate(&Product{})
}

// 插入
func insert_data(db *gorm.DB) {
	p := Product{
		Code:  "1001",
		Price: 100,
	}
	db.Create(&p)
}

// 查询
func query_data(db *gorm.DB) {
	var p Product
	db.First(&p, 2)                   // 主键查询
	db.First(&p, "code = ? ", "1001") // 查询code为1001的数据
	fmt.Printf("p: %v\n", p)
}

// 更新
func update_data(db *gorm.DB) {
	var p Product
	db.First(&p, 2)
	db.Model(&p).Update("Price", 1000)
	db.Model(&p).Updates(Product{Price: 1002, Code: "1002"}) // 仅更新非零值字段
	db.Model(&p).Updates(map[string]interface{}{"Price": 1003, "Code": "1003"})
}

// 删除数据
func delete_data(db *gorm.DB)  {
	var p Product
	db.First(&p, 2)
	db.Delete(&p, 2) // 标记删除
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	// create_db(db)
	// insert_data(db)
	// query_data(db)
	// update_data(db)
	delete_data(db)
}
