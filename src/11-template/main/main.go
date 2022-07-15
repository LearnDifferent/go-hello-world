package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	textTemplate()
	fmt.Println("打印 HTML 的文字时，会转换特殊字符：")
	htmlTemplate()
}

type sentence struct {
	// 注意，只有变量名大写，才能被外界读取
	Name   string
	Secret string
}

func textTemplate() {
	sentence := sentence{"John", "I have no secret."}

	// 模版文件的路径（如果使用相对路径，应该从项目的根路径开始）
	templatePath := "./main/template-demo.txt"

	// template.New(模版文件名).ParseFiles(模版文件所在的完整路径)
	file, err := template.New("template-demo.txt").ParseFiles(templatePath)
	if err != nil {
		fmt.Println(err)
	}

	// 第一个参数表示打印
	// os.Stdout: print out the result using standard out
	// 第二个参数：将该参数，填入模版中
	err = file.Execute(os.Stdout, sentence)
	if err != nil {
		fmt.Println(err)
	}
}

type page struct {
	Title   string
	Content string
}

func htmlTemplate() {
	p := page{"I'm here", "This is the content"}
	templatePath := "./main/html-template-demo.html"
	file, err := template.New("html-template-demo.html").ParseFiles(templatePath)
	if err != nil {
		fmt.Println(err)
	}

	err = file.Execute(os.Stdout, p)
	if err != nil {
		fmt.Println(err)
	}
}
