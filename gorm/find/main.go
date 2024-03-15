package main

import (
	"errors"
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
	var u User
	// 获取第一条记录（主键升序）
	db.First(&u)
	fmt.Println(u)
	// 获取一条记录，没有指定排序字段
	var u1 User
	db.Take(&u1)
	fmt.Println(u1)
	// 获取最后一条记录（主键降序）
	var u2 User
	db.Last(&u2)
	fmt.Println(u2)

	var u4 User
	result := db.First(&u4)
	fmt.Println("返回找到的记录数=", result.RowsAffected)

	if result.Error != nil {
		panic(result.Error)
	}

	// 检查 ErrRecordNotFound 错误
	errors.Is(result.Error, gorm.ErrRecordNotFound)

	//根据主键检索
	var u5 User
	db.Find(&u5, []int{1, 2})
	fmt.Println(u5)
	var u6 User
	db.Find(&u6, User{Name: "李涛"})
	fmt.Println(u6)
	//检索全部对象

	var u7 User
	all := db.Find(&u7)
	fmt.Print(all)

	//string条件
	//=等于 <>不等于 in包含 like模糊查询 and和 time按时间 between？and？和
	var u8 User
	db.Where("name = ?", "李涛").First(&u8)
	fmt.Println(u8)
	//db.Where("name LIKE ?", "%jin%").Find(&users)
	//db.Where("updated_at > ?", lastWeek).Find(&users)
	//db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)

	//struct & map 条件
	var user User
	db.Where(&User{Name: "张然", Age: 23}).First(&user)
	fmt.Println(user)
	//db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)

	//内联条件
	// 	// Get by primary key if it were a non-integer type
	// db.First(&user, "id = ?", "string_primary_key")
	// // SELECT * FROM users WHERE id = 'string_primary_key';

	// // Plain SQL
	// db.Find(&user, "name = ?", "jinzhu")
	// // SELECT * FROM users WHERE name = "jinzhu";

	// db.Find(&users, "name <> ? AND age > ?", "jinzhu", 20)
	// // SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;

	// // Struct
	// db.Find(&users, User{Age: 20})
	// // SELECT * FROM users WHERE age = 20;

	// // Map
	// db.Find(&users, map[string]interface{}{"age": 20})
	// // SELECT * FROM users WHERE age = 20;

	//not条件
	var u9 User
	db.Not("name = ?", "张然").First(&u9)
	fmt.Println(u9)

	//or 条件(或者)
	//db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)

	//选择特定字段
	var u10 User
	db.Select("name", "age").Find(&u10)
	fmt.Println(u10)
	//db.Select([]string{"name", "age"}).Find(&users)
	//db.Table("users").Select("COALESCE(age,?)", 42).Rows()

	var u12 User
	db.Order("age desc, name").Find(&u12)
	fmt.Println(u12)

	//Limit 指定要检索的最大记录数
	//Offset 指定开始返回记录之前要跳过的记录数
	//limit and offset 用法一样
	db.Limit(1).Find(&user)
	//db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
	// SELECT * FROM users LIMIT 10; (users1)
	// SELECT * FROM users; (users2)
	//db.Limit(10).Offset(5).Find(&users)
	// SELECT * FROM users OFFSET 5 LIMIT 10;

	//group by & having 分组依据和拥有
	//db.Model(&User{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "group").Find(&result)
	//从数据库中选择名为name的列和age列的总和作为total，并按name进行分组。
	//使用Having子句来筛选具有特定name值的分组。
	//将结果存储到result变量中。

	//scan扫描
	type Result struct {
		Name string
		Age  int
	}
	var resultscan Result
	db.Table("users").Select("name", "age").Where("name = ?", "Antonio").Scan(&resultscan)
	fmt.Println(resultscan)
}
