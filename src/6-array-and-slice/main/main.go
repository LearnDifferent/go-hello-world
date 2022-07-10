package main

import "fmt"

func main() {
	var defaultArr [3]int
	fmt.Println("声明一个大小为 3 的 Array，查看默认值：", defaultArr)
	fmt.Println("打印 Array 的大小：", len(defaultArr))

	defaultArr[0] = 2
	fmt.Println("为已声明的数组的某个 index 位置赋值：", defaultArr[0])

	//=====================
	fmt.Println("原数组：", defaultArr)
	// 复制数组
	copyArr := defaultArr
	fmt.Println("复制了一个数组：", copyArr)
	// 对比原数组和复制的数组
	fmt.Println("原数组和复制的数组，现在是相等的：", defaultArr == copyArr)
	// 修改原数组
	defaultArr[1] = 6
	fmt.Println("修改了原数组，原数组会发生改变，但是复制的数组并不会被改变：", defaultArr, copyArr)
	// 修改了原数组后，再次对比
	fmt.Println("原数组被修改了，复制的数组还是不变，所以两个数组不相等了：", defaultArr == copyArr)

	lenFourArr := [4]int{1, 2}
	fmt.Println("声明一个长度为 4 的数组，但是只给 index 0 和 1 赋值：", lenFourArr)

	// 注意，是 [...]，而不是 []
	justArr := [...]int{1, 3}
	fmt.Println("不声明长度，也可以直接给数组赋值：", justArr)

	fmt.Println("遍历数组：")
	for i, j := range justArr {
		fmt.Println("index:", i, ", val:", j)
	}

	specArr := [...]int{1: 99, 5: 100}
	fmt.Println("给 index 1 和 index 5 赋值：", specArr)

	strArr := [...]string{"one", "two"}
	fmt.Println("字符串数组：", strArr)

	arr2d := [2][2]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println("二维数组：", arr2d)

	// 声明一个 slice
	sl := []int{1, 2, 3}
	fmt.Println("Slice 的长度，指的是它里面有几个值：", len(sl))

	newSl := append(sl, 4, 5)
	fmt.Println("可以使用 append 方法添加新的元素，并返回新的 slice：", newSl)

	fmt.Println("Slice 有 capacity 的概念，cap >= len：", cap(sl), len(sl))

	// 可以使用这种的方式重新获取 slice
	sl = sl[0:3]
	fmt.Println(sl, len(sl))
	// 如果小于当前 slice 的长度，还会让 slice 变小
	sl = sl[0:2]
	fmt.Println(sl, len(sl))
	// 不能超过当前 slice 的长度，否则会报错：
	//sl = sl[0:100]

	sl = append(sl, 4, 5, 6, 7, 8)
	fmt.Println("如果是 append 之后，重新给当前 slice 赋值的话", sl)
	sl = sl[0:5]
	fmt.Println("此时在范围内的可以打印", sl)
	sl = sl[0:8]
	fmt.Println("在长度范围外，在容量范围内的话，可以自动添加位数", sl)

	madeSlice := make([]int, 5)
	fmt.Println("可以通过 make 方法，在初始的 slice 的时候分配内存空间", madeSlice)

	var a []int
	fmt.Println("如果声明一个 slice，但是没有给它赋值，那么它就是 nil", a == nil)
}