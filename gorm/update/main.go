package main

import (
	"git/gorm/connect"

	"gorm.io/gorm"
)

func main() {
	db, err := connect.Connect()
	if err != nil {
		panic(err)
	}
	//保存所有字段
	var user connect.User
	db.First(&user)
	// user.Name = "jinzhu 2"
	// user.Age = 100
	// db.Save(&user)
	//db.Save(&connect.User{ID: 1, Name: "jinzhu 2", Age: 20, Birthday: nil})
	//db.Save(&connect.User{ID: 1, Name: "jinzhu", Age: 100})

	//更新单个列
	// 根据条件更新
	//db.Model(&user).Where("Age = ?", 22).Update("name", "hello")
	//更新多列
	// 根据 `struct` 更新属性，只会更新非零值的字段
	//db.Model(&user).Where("id = ?", 2).Updates(connect.User{Name: "hello", Age: 18, Gender: "男"})
	//根据map更新
	//db.Model(&user).Where("id = ?", 1).Updates(map[string]interface{}{"name": "hello", "age": 19})
	//更新选定的字段
	// 选择 Map 的字段
	//db.Model(&user).Select("name", "age").Updates(map[string]interface{}{"name": "qq", "age": 18})
	// 选择 Struct 的字段（会选中零值的字段）
	//db.Model(&user).Select("name", "age").Updates(connect.User{Name: "tao", Age: 10})
	// 选择除 Role 外的所有字段（包括零值字段的所有字段）omit除了该字段
	//db.Model(&user).Select("*").Omit("Role").Updates(User{Name: "jinzhu", Role: "admin", Age: 0})
	//使用sql更新
	db.Model(&user).Update("age", gorm.Expr("age+?", 2))
}
