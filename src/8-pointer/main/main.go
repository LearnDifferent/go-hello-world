package main

import "fmt"

func main() {
	var str string = "123"
	fmt.Println("原始：", str)

	changeStr(str)
	fmt.Println("将值作为参数：", str)

	changeStrPointer(&str)
	fmt.Println("将地址作为参数传入后，原值也会被修改：", str)
}

func changeStr(str string) {
	str = "abc"
}

// *string 表示这是一个 pointer variable
func changeStrPointer(str *string) {
	// *str 表示修改的是 reference
	*str = "456"
}