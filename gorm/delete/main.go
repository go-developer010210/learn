package main

import (
	"git/gorm/connect"
)

func main() {
	db, err := connect.Connect()
	if err != nil {
		panic(err)
	}
	var user = connect.User{}
	// user.ID = 1

	// db.Delete(&user)
	//根据主键删除
	//db.Delete(&connect.User{}, 3)
	//db.Delete(&users, []int{1,2,3})
	// DELETE FROM users WHERE id IN (1,2,3);

	//批量删除
	//db.Where("email LIKE ?", "%jinzhu%").Delete(&Email{})
	// DELETE from emails where email LIKE "%jinzhu%";

	//db.Delete(&Email{}, "email LIKE ?", "%jinzhu%")
	// DELETE from emails where email LIKE "%jinzhu%";
	//var users = []User{{ID: 1}, {ID: 2}, {ID: 3}}
	//db.Delete(&users)

	//查找被软删除的记录
	db.Unscoped().Where("age = ?", 12).Find(&user)
	//永久删除
	//db.Unscoped().Delete(&user)
}
