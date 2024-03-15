package main

import (
	connect01 "git/gorm_association/connect_01"

	"gorm.io/gorm"
)

// User 有多张 CreditCard，UserID 是外键
type User struct {
	gorm.Model
	CreditCards []CreditCard
}

//想要使用另一个字段作为外键，您可以使用 foreignKey 标签自定义它：
//`gorm:"foreignKey:UserRefer"`
type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

func main() {
	connect01.Init()
	db := connect01.Db
	db.AutoMigrate(&User{})
}
