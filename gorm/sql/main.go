package main

import "git/gorm/connect"

type Result struct {
	ID   int
	Name string
	Age  int
}

func main() {
	db, err := connect.Connect()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Result{})

	//原生sql 查询raw
	//增加 删除 修改 exec
}
