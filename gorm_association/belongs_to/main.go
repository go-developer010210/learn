package main

import (
	connect01 "git/gorm_association/connect_01"

	"gorm.io/gorm"
)

// `User` 属于 `Company`，`CompanyID` 是外键
type User struct {
	gorm.Model
	Name         string
	CompanyRefer int
	Company      Company `gorm:"foreignKey:CompanyRefer"`
}

type Company struct {
	ID   int
	Name string
}

func main() {
	connect01.Init()
	db := connect01.Db
	db.AutoMigrate(&User{})

}
