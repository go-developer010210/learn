package connect01

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var Err error

func Init() *gorm.DB {
	dsn := "root:123456@tcp(localhost:3305)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	Db, Err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if Err != nil {
		fmt.Println("连接失败", Err)
	}
	return Db

}
