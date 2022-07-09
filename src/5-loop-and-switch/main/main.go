package main

import "fmt"

func main() {
	fmt.Println("====打印从 0 到 9====")
	fmt.Println("----for 加上终止调节----")
	i0 := 0
	for i0 < 10 {
		fmt.Print(i0)
		i0++
	}
	fmt.Println()

	fmt.Println("----将变量声明放到 for 里面----")
	for i1 := 0; i1 < 10; {
		fmt.Print(i1)
		i1++
	}
	fmt.Println()

	fmt.Println("----将 step 也放到 for 里面----")
	for i2 := 0; i2 < 10; i2++ {
		fmt.Print(i2)
	}

	fmt.Println("====打印每个字符及其对应的 index====")
	s := "Hello World"

	fmt.Println("----for loop----")
	for i := 0; i < len(s); i++ {
		fmt.Printf("'%c' is index %d \n", s[i], i)
	}

	fmt.Println("----for range loop----")
	// for (index), (each item) := range (items)
	for i, c := range s {
		fmt.Printf("'%c' is index %d \n", c, i)
	}

	fmt.Println("----for range loop: ignore index and break----")
	for _, c := range s {
		if c == ' ' {
			// 如果 c 是空字符，就断开循环
			break
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()

	fmt.Println("----for range loop: ignore index and continue----")
	for _, c := range s {
		if c == ' ' {
			// 如果 c 是空字符，就跳过
			continue
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()

	fmt.Println("====Label to break Nested for loop====")
outer:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				break outer
			}
			fmt.Printf("%d%d \n", i, j)
		}
	}

	fmt.Println("Go 的 switch 语句不需要 break：")
	//num := 1
	//num := 2
	//num := 3
	//num := 4
	num := 5
	switch num {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3,4:
		fmt.Println("three or four")
	default:
		fmt.Println("Not one to four")
	}
}
