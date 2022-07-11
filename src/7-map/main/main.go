package main

import "fmt"

func main() {
	var a map[string]string
	fmt.Printf("纯粹声明一个 %T，此时 %T 是否为 nil：%t \n", a, a, a == nil)

	var b = map[string]string{}
	fmt.Printf("声明一个 %T 后，使用 {} 来初始化，此时 %T 是否为 nil: %t \n", b, b, b == nil)

	c := make(map[string]string)
	fmt.Println("还可以使用 make 方法来初始化一个 map", c)

	m := map[string]string{"key1": "value1"}
	fmt.Println("直接在声明的时候赋值", m)

	m["key2"] = "value2"
	fmt.Println("可以继续添加 key value", m)

	fmt.Println("获取某个 key 的 value", m["key1"])

	delete(m, "key1")
	fmt.Println("delete 方法可以删除 map 里面的内容", m)
	fmt.Println("此时再获取该 key，就是空字符串：", m["key1"] == "")

	fmt.Println("获取 key 的时候，可以获取两个返回值，第一个是该 key 的 value，第二个是该 key 是否存在：")
	val, isExisted := m["key1"]
	if isExisted {
		fmt.Println("如果该 key 存在，就打印", val)
	} else {
		fmt.Println("不存在就打印现在这段")
	}

	m["key2"] += "abc"
	fmt.Println("小技巧：这样可以很直接在 string 类型的 value 后面添加字符串", m)

	m["k1"] = "v1"
	m["k2"] = "v2"
	m["k3"] = "v3"
	m["k4"] = "v4"
	fmt.Println("遍历 map 中的 key 和 value:")
	for k, v := range m {
		fmt.Println(k, v)
	}
}
