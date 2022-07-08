package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "This is a sentence"

	fmt.Println(str)
	fmt.Println("The length of the sentence is ", len(str))
	fmt.Println("获取 Index 为 0 的字符的 ASCII：", str[0])
	fmt.Printf("打印 Index 为 0 的字符的 ASCII 的字符：%c \n", str[0])
	fmt.Println("直接获取 Index 为 0 的字符，也就是获取 [0, 1) 的字符：", str[:1])
	fmt.Println("获取前两个字符：", str[:2])
	fmt.Println("获取 index 从 1 开始，也就是除去第一个字符：", str[1:])
	fmt.Println("获取 index [2, 4) 的字符，也就是第三个和第四个字符：", str[2:4])
	fmt.Println("使用 Backspace 符号，消除空格：这个 \b中间 \b应该有很 \b多空格")

	abcStr := "abc"
	abcByte := []byte(abcStr)
	fmt.Printf("String 就是 byte 数据的集合，所以可以直接转换为 ASCII 的 byte 数组。如果要打印出来的话，可以这样： %s \n",
		abcByte)

	abcStrFromByteArr := string(abcByte)
	fmt.Println("byte 数组可以直接转换为 String：", abcStrFromByteArr)

	chineseHello := "你好"
	fmt.Println("每个中文汉字的长度为 3：", len(chineseHello))
	fmt.Println("获取第一个汉字，相当于获取 [0, 3)：", chineseHello[:3])

	fmt.Println("大写：", strings.ToUpper("Hello"))

	completeStr := "Hello World"
	prefixStr := "Hello"
	fmt.Println("检查", completeStr, "是否以", prefixStr, "开头：",
		strings.HasPrefix(completeStr, prefixStr))

	oldStr, newStr := "o", "O"
	fmt.Println("替换", completeStr, "中第一个", oldStr, "为", newStr, "：",
		strings.Replace(completeStr, oldStr, newStr, 1))

	fmt.Println("是否包含：", strings.Contains(completeStr, prefixStr))

	var sb strings.Builder
	sb.WriteString("这是 String Builder")
	sb.WriteString("，可以持续拼接 String")

	fmt.Println("String Builder：", sb.String())
	fmt.Println("String Builder 的长度：", sb.Len())
	fmt.Println("String Builder 的容量（大于等于长度）：", sb.Cap())

	sb.Grow(100)
	fmt.Println("将 String Builder 的容量在'原容量 * 2' 的基础上，再 +100：", sb.Cap())

	sb.Reset()
	fmt.Println("将 String Builder 清空后，它会变成空字符串：", sb.String())

	capTripleA := []byte{65, 65, 65}
	sb.Write(capTripleA)
	fmt.Println("String Builder 可以直接存入 byte 数组：", sb.String())

	// 将 integer 转换为 a string
	i := 123
	s := strconv.Itoa(i)
	fmt.Printf("将 %T 转换为 %T \n", i, s)

	// 将 a string 转换为 integer
	// strconv.Atoi 会返回 int, error 两个值
	// 这里的 z 表示返回的 int 值，而这里的下划线 _ 表示"不去理会可能会出现的 error"
	z, _ := strconv.Atoi(s)
	fmt.Printf("将 %T 转换为 %T \n", s, z)
}