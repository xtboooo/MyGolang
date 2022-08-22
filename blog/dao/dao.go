package dao

import (
	"blog/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Manager interface {
	Register(user *models.User)
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic("failed to init database")
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&models.User{})
}

func (mgr *manager) Register(user *models.User) {
	mgr.db.Create(user)
}
