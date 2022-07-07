package main

import (
	"fmt"
	"math"
)

func main() {
	// 自动决定是 int32 还是 int64
	var i = 1
	var i32 int32 = 1
	var i64 int64 = 1
	fmt.Println("打印出来都是一样的：", i, i32, i64)
	fmt.Printf("但是，如果打印的是类型，就不一样了： %T, %T, %T \n", i, i32, i64)

	// 浮点数
	pi := 3.1415
	fmt.Println("打印一个 short declaration 的浮点数：", pi)
	fmt.Printf("浮点数的默认类型是 float64：%T \n", pi)

	var pi32 float32 = 3.1415
	fmt.Printf("可以将浮点数显示地设置为 float32：%T \n", pi32)

	// 转换 Integer 和 Float
	a := 3.99
	var b int
	b = int(a)
	fmt.Println("将 Float 转换为 Integer 会直接损失小数点后面的内容：", b)

	a = float64(b)
	fmt.Println("将 Integer 转换为 Float，不会显示小数点，看起来和 Integer 一样：", a)

	fmt.Println("Ceiling:", math.Ceil(3.00))
	fmt.Println("Ceiling:", math.Ceil(3.01))
	fmt.Println("Floor:", math.Floor(3.00))
	fmt.Println("Floor:", math.Floor(3.00))

	fmt.Println("Min:", math.Min(1, 2))
	fmt.Println("Abs:", math.Abs(-2))
	fmt.Println("Sqrt:", math.Sqrt(100))
	fmt.Println("Pow:", math.Pow(2, 3))
}
