package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	// 注意，这里要大写
	Name string
	Age  int
}

type teacher struct {
	// 可以使用这种方式，来匹配 Json 中的属性
	Name string `json:"teacher_name"`
	Age  int
}

func main() {
	normalJsonToStruct()
	fmt.Println("==============")
	specJsonToStruct()
	fmt.Println("==============")
	structToJson()
}

func normalJsonToStruct() {
	jsonForm := `{
					"name": "sally", 
					"age": 16
				}`

	fmt.Println("初始化一个 Json，可以使用 `` 符号来定义多行字符串：", jsonForm)

	// 定义一个 struct 变量
	var sally student
	// 使用 Json 工具
	err := json.Unmarshal([]byte(jsonForm), &sally)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sally)
}

func specJsonToStruct() {
	jsonForm := `{
					"teacher_name": "John", 
					"age": 46
				}`
	var john teacher
	err := json.Unmarshal([]byte(jsonForm), &john)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(john)
}

func structToJson() {
	wang := teacher{"Wang", 48}
	wangJson, err := json.Marshal(wang)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("打印的是 ASCII 编码，因为这是一个 []byte", wangJson)
	fmt.Printf("正确的打印方式 %s", wangJson)
}