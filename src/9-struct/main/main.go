package main

import "fmt"

func main() {

	// 定义一个 struct
	type student struct {
		name     string
		age      int
		teachers []string
	}

	// 可以这样直接创建一个实例，这样就不需要按顺序对应相应属性
	sally := student{
		age:      16,
		name:     "Sally",
		teachers: []string{"Larry", "John"},
	}
	fmt.Println(sally)

	// 还可以按照相应属性，直接赋值来创建实例（这里是重新赋值给 sally 变量）
	sally = student{"Sally Porter", 15, []string{"Li", "Wang"}}
	fmt.Println(sally)

	// 获取属性
	fmt.Println(sally.age)

	// 直接操作属性
	sally.age++
	fmt.Println(sally.age)

	sally.teachers = append(sally.teachers, "Chen")
	fmt.Println(sally.teachers)

	// struct 也可以作为属性
	type studentParent struct {
		name  string
		child student
	}

	fatherOfSally := studentParent{
		"Father of Sally",
		sally,
	}
	fmt.Println(fatherOfSally)

	fmt.Println("原始：fatherOfSally.child.age:", fatherOfSally.child.age)
	fatherOfSally.child.age++
	fmt.Println("+1 之后：", fatherOfSally.child.age)
	fmt.Println("因为这里传递的是值，所以并不会改变 sally.age 的值：", sally.age)

	// 这里传递的是 struct 指针
	type parentWithStudentReference struct {
		name  string
		child *student
	}
	sallyMother := parentWithStudentReference{"Mother Of Sally:", &sally}
	fmt.Println("sallyMother:", sallyMother)
	fmt.Println("打印地址：&sallyMother.child:", &sallyMother.child)
	fmt.Println("打印指针指向的原值，也就是 * 表示引用：*sallyMother.child:", *sallyMother.child)

	fmt.Println("原值：", sallyMother.child.age)
	sallyMother.child.age++
	fmt.Println("传递地址的时候，传递的是引用，所以可以改变原值：", sallyMother.child.age)

	// 使用 new 关键字，会返回一个 struct 的 pointer，而不是实例
	johnPointer := new(student)
	fmt.Println("因为这是一个指针，所以它存储的值是'存储地址的变量'：", johnPointer)
	fmt.Println("如果想打印原值，就使用 * 关键字传递引用：", *johnPointer)
	//下面的，其实就等于：(*johnPointer).name = "John"
	johnPointer.name = "John"
	johnPointer.age = 20
	fmt.Println("指针可以改变原值：", *johnPointer)

	wang := teacher{"Wang"}
	wang.sayHello()
	wang.intro()
}

type teacher struct {
	name string
}

// 给一个 struct 添加方法
func (teacher) sayHello() {
	fmt.Println("Hello")
}

// 给一个 struct 添加一个使用了该 struct 属性的方法
func (t teacher) intro() {
	fmt.Println("I'm ", t.name)
}