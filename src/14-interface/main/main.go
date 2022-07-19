package main

import "fmt"

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

func main() {
	// 因为两个 struct 都实现了 intro 接口
	// 所以可以他们的 type 也可以是 intro
	pplInSchool := []intro{teacher{"T"}, student{"S"}}
	for _, ppl := range pplInSchool {
		ppl.sayName()
	}
}