package main

import "fmt"

func main() {

	defer safeExit()

	if 1 != 2 {
		panic("Panic 会停止程序，后面的代码不会运作。")
	}
}

func safeExit() {
	if r := recover(); r != nil {
		fmt.Println("要 recover 一个 Panic 的话，要使用 defer 来确保 recover 能成功")
	}
}