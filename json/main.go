package main

import (
	"encoding/json"
	"fmt"
)

// 结构体序列化
type People struct {
	Name string
	Age  int
}

// map序列化
func testMap() {
	//定义一个map
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "火云洞"

	data, err := json.Marshal(a) //将a这个map进行序列化
	if err != nil {
		fmt.Printf("序列化错误 err=%v", err)
	}
	fmt.Printf("a map序列化后=%v\n", string(data))
}

func main() {
	//结构体
	people := People{
		Name: "张然",
		Age:  18,
	}
	//通过json包下的Marshal函数对结构体进行转换
	data, err := json.Marshal(&people)
	if err != nil {
		fmt.Printf("序列化错误 err=%v", err)
	}
	fmt.Printf("people序列化后=%v\n", string(data))
	//map
	testMap()
}
