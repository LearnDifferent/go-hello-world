package main

import (
	"errors"
	"fmt"
)

func main() {
	foo()

	sumResult1 := sum1(1, 3)
	fmt.Println(sumResult1)

	sumResult2 := sum2(3, 5)
	fmt.Println(sumResult2)

	firstLetter, length := getFirstLetterAndLength("abc")
	fmt.Println("有多个返回值：", firstLetter, length)

	fmt.Print("有多个返回值时，还可以直接打印：")
	fmt.Println(getFirstLetterAndLength("123456"))

	manyIntegers(1, 3, 5, 7)

	err := oneIsErrorOtherIsNil(2)
	if err != nil {
		fmt.Println("one")
	} else {
		fmt.Println("other")
	}

	num := 1
	fmt.Println("原值为：", num)
	withoutPointer(num)
	fmt.Println("非指针的方法，值不会改变：", num)
	withPointer(&num)
	fmt.Println("有指针的方法，会改变值：", num)

	fmt.Printf("方法也是可以被当成参数：%T \n", returnParam)

	result := functionAsParam(5, returnParam)
	fmt.Println("将方法作为参数传递：", result)

	// 匿名方法
	anonymousFunction := func() {
		fmt.Println("anonymous function 1")
	}
	// 可以调用
	anonymousFunction()

	// 可以将该匿名方法，重新赋给另一个方法
	// 注意，必须是结构相同的方法，比如没有参数的 void 方法，只能重新赋给没有参数的 void
	anonymousFunction = foo
	// 这样再次调用，方法就不同了
	anonymousFunction()

	// 如果想声明一个匿名方法，并直接调用
	// 直接在方法后面加上：()
	func() {
		fmt.Println("anonymous function 2")
	}()

	fmt.Println("正常顺序：")
	printFirst()
	printSecond()
	fmt.Println("使用 defer 将第一个方法放到 main 方法的最后执行：")
	defer printFirst()
	printSecond()
	printSecond()
}

// 普通的方法
func foo() {
	fmt.Println("Hi")
}

// Go 不支持重载，所以下面的方法会失效
//func foo(i int) {
//	fmt.Println("Hi", i)
//}

// 有返回值的方法
func sum1(a int, b int) int {
	return a + b
}

// 还可以这样写：
func sum2(a, b int) int {
	return a + b
}

// 返回多个值
func getFirstLetterAndLength(str string) (string, int) {
	return str[0:1], len(str)
}

// 变长参数
func manyIntegers(integers ...int) {
	for _, i := range integers {
		fmt.Print(i)
	}
	fmt.Println()
}

// 返回 error
func oneIsErrorOtherIsNil(oneOrOther int) error {
	if oneOrOther == 1 {
		return errors.New("AN ERROR")
	}
	return nil
}

// 没有 pointer，那这个参数就是拷贝值进来，不会影响原来的值
func withoutPointer(i int) {
	i++
}

// 如果有 pointer，那这个参数就是传递引用，方法里面进行修改会影响原来的值
func withPointer(i *int) {
	*i++
}

// 返回当前参数的方法
func returnParam(i int) int {
	return i
}

// 参数 1 是 integer
// 参数 2 是一个方法
func functionAsParam(i int, f func(int) int) int {
	// 将参数 1 作为参数，放入参数 2，也就是方法中
	// 返回该方法获取到的值
	return f(i)
}

func printFirst() {
	fmt.Println("First")
}

func printSecond() {
	fmt.Println("Second")
}