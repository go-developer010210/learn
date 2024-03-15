package connect

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Age      uint8
	Birthday *time.Time
	Gender   string
}

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	dsn := "root:123456@tcp(localhost:3305)/learn?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接失败")
	}

	// 迁移模式
	// db.AutoMigrate(&User{})
	DB = db
	return DB, err
}
