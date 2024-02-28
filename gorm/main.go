package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       uint
	Name     string
	Age      uint8
	Birthday time.Time
}

func main() {
	dsn := "root:123456@tcp(localhost:3305)/learn?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接失败")
	}

	// 迁移模式
	db.AutoMigrate(&User{})

	users := []*User{
		{Name: "张然", Age: 23, Birthday: time.Now()},
		{Name: "李涛", Age: 22, Birthday: time.Now()},
	}
	// 通过数据的指针来创建
	result := db.Create(&users)
	// 返回插入数据的主键
	for _, user := range users {
		fmt.Println("插入的ID:", user.ID)
	}
	// 返回 error
	if result.Error != nil {
		panic(result.Error)
	}
	// 返回插入记录的条数
	fmt.Println("插入的记录数量:", result.RowsAffected)
}
