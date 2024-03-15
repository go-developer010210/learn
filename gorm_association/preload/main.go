package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Orders   []Order
}
type Order struct {
	gorm.Model
	UserID uint
	Price  float64
}

//使用 Preload通过多个SQL中来直接加载关系
// 查找 user 时预加载相关 Order
//db.Preload("Orders").Find(&users)
// SELECT * FROM users;
// SELECT * FROM orders WHERE user_id IN (1,2,3,4);

//db.Preload("Orders").Preload("Profile").Preload("Role").Find(&users)
// SELECT * FROM users;
// SELECT * FROM orders WHERE user_id IN (1,2,3,4); // has many
// SELECT * FROM profiles WHERE user_id IN (1,2,3,4); // has one
// SELECT * FROM roles WHERE id IN (4,5,6); // belongs to