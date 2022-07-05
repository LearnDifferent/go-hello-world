package main

import (
	"fmt"
)

func main() {
	// var name type = expression

	// ========Integer========
	var i1 int = 1
	fmt.Println("Integer: ", i1)

	var i0 int
	fmt.Println("Integer 默认值是 : ", i0)

	var i2, i3 int
	fmt.Println("同时声明两个 Integer 默认值：", i2, i3)

	var i4 = 4
	fmt.Println("声明的时候，可以不写 type：", i4)

	var i5, i6 = 5, 6
	fmt.Println("声明两个不写 type 的 Integer：", i5, i6)

	var i7, s1 = 7, "eight"
	fmt.Println("同时声明一个 Integer 和一个 String：", i7, s1)

	// short declaration: name := expression
	i8 := 8
	fmt.Println("short declaration：", i8)

	// ========Pointers========
	x := "参数 x"
	p := &x
	fmt.Println("打印指针 P，该指针指向", x, "在内存中的地址：", p)
	fmt.Println("打印指针 P 指向的", x, "的实际的值：", *p)

	var x1 string = "参数 x1"
	var p1 *string = &x1
	fmt.Println("上面的代码，完整的写法，也可以得出该结果：", x1, p1)

	// ========Type========
	type cusType = int
	var i9 cusType = 9
	fmt.Println("类型也可以声明赋值，比如上面的这个也能获取到：", i9)

	// ========Constants========
	const normalConst = 5
	fmt.Println("正常定义一个常量：", normalConst)

	const (
		a1 = 1
		// B 的值同上
		a2
		b1 = 2
		// d 的值同上
		b2
	)
	fmt.Println("定义多个常量：", a1, a2, b1, b2)

	const (
		zero = iota
		one
		two
	)
	fmt.Println("使用 iota 定义递增的常量：", zero, one, two)

	const (
		three = iota + 3
		four
		five
	)
	fmt.Println("使用 iota 定义从 3 开始的递增常量：", three, four, five)

	// ========untyped constants========
	const (
		_ = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
		eb
		zb
		yb
	)

	//fmt.Println("yb 是非类型化的常量，在实际使用时，会被转换为 Go 的类型这里的 yb 在转换为 int 后，因为数值超过了 int 的范围，所以会报错", yb)
	fmt.Println("但是 yb 在运算的时候，会先被转换为 an array of unsigned integer，"+
		"然后再参与计算，最后得到的值再被转换为 int，所以只要最后的运算结果在 int 的范围内，就不会报错", yb/zb)

	// 常量也不能是需要运算才能得出结果的值，否则无法通过编译。因为这种值是变量
	//const bigNum = math.Pow(2, 100)
}