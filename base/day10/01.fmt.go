package main

import "fmt"

// Print函数直接输出内容
// Printf函数支持格式化输出字符串
// Println函数会在输出内容的结尾添加一个换行符
func main() {
	fmt.Print("在终端打印该信息。")
	name := "枯藤"
	fmt.Printf("我是：%s\n", name)
	fmt.Println("在终端打印单独一行显示")
}
