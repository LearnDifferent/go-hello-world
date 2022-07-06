package main

import (
	"2-scope/integers"
	"fmt"
)

var packageLevelVal = "Package 的变量"

func main() {
	{
		// block（代码块）内部的变量，范围只在代码块之内
		blockedVal := 1
		fmt.Println("只有在代码块中可以访问：", blockedVal)
	}

	//fmt.Println("在代码块之外无法访问：", blockedVal)

	fmt.Println("Main 方法可以访问 Package 级别的变量：", packageLevelVal)
	fmt.Println("可以访问同一个 Package 下，其他 go 文件的 Package 变量：", otherVal)

	//fmt.Println("不能访问另一个 Package 下，开头为小写的 Package 变量：", integers.five)
	fmt.Println("在 import 了另一个 Package 后，可以访问另一个 Package 下，开头为大写的 Package 变量：",
		integers.Six)
}

func nonMain() {
	fmt.Println("非 Main 方法可以访问 Package 级别的变量：", packageLevelVal)
}