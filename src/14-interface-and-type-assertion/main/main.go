package main

import (
	"fmt"
)

//====接口====
type intro interface {
	sayName()
}

//====老师及其方法====
type teacher struct {
	name string
}

//----只要实现了接口的方法，就是实现了接口----
func (teacher) sayName() {
	fmt.Println("I'm a teacher")
}

//====学生及其方法====
type student struct {
	name string
}

//----只要实现了接口的方法，就是实现了接口----
func (student) sayName() {
	fmt.Println("I'm a student")
}

//----学生再添加一个方法----
func (student) study() {
	fmt.Println("I'm studying")
}

func main() {
	// 因为两个 struct 都实现了 intro 接口
	// 所以可以他们的 type 也可以是 intro
	pplInSchool := []intro{teacher{"T"}, student{"S"}}
	for _, ppl := range pplInSchool {
		// 当前 type 是 intro，所以可以使用 intro 接口的方法
		ppl.sayName()
		// 可以强制转换类型，但是如果转换失败，就会 Panic：
		//ppl.(student).study()

		// 所以，在强制转换前，需要判断类型
		// 第一个返回值是强制转换后的值
		// 第二个返回值是 bool，表示是否能强制转换
		stu, ok := ppl.(student)
		if ok {
			stu.study()
		}
	}
}