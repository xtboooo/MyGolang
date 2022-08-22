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
	// 全局配置
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info), DryRun: true})
	if err != nil {
		panic("failed to connect database")
	}
	db = d
}

func test1()  {
	// 会话配置
	db.Session(&gorm.Session{DryRun: true})
}

func main() {
}
 */