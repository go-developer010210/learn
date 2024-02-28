package main

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(localhost:3305)/learn?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	user := User{Name: "张然", Age: 18, Birthday: time.Now()}
	// 通过数据的指针来创建
	result := db.Create(&user)
	// 返回插入数据的主键
	user.ID
	// 返回 error
	result.Error(err)
	// 返回插入记录的条数
	result.RowsAffected

}
