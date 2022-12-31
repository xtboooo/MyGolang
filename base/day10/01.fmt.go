package main

import (
	"fmt"
	"os"
)

// Print函数直接输出内容
// Printf函数支持格式化输出字符串
// Println函数会在输出内容的结尾添加一个换行符
func f1() {
	fmt.Print("在终端打印该信息。")
	name := "枯藤"
	fmt.Printf("我是：%s\n", name)
	fmt.Println("在终端打印单独一行显示")
}

// Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容
func f2() {
	// 向标准输出写入内容
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./format_demo.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	name := "枯藤"
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)
}

// Sprint系列函数会把传入的数据生成并返回一个字符串。
func f3() {
	s1 := fmt.Sprint("枯藤")
	name := "枯藤"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("枯藤")
	fmt.Println(s1, s2, s3)
}

// Errorf函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。
func f4() {
	err := fmt.Errorf("这是一个错误")
	fmt.Printf("err: %v\n", err)
}

// 格式化占位符
func f5() {
	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)
	o := struct{ name string }{"枯藤"}
	fmt.Printf("%v\n", o)
	fmt.Printf("%+v\n", o)
	fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)
	fmt.Printf("100%%\n")
}

// 整型
func f6() {
	n := 65
	fmt.Printf("%b\n", n)
	fmt.Printf("%c\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X\n", n)
}

// 浮点数与复数
func f7() {
	f := 12.34
	fmt.Printf("%b\n", f)
	fmt.Printf("%e\n", f)
	fmt.Printf("%E\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%g\n", f)
	fmt.Printf("%G\n", f)
}

// 字符串和[]byte
func f8() {
	s := "枯藤"
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("%X\n", s)
}

func main() {
	// f1()
	// f2()
	// f3()
	// f4()
	// f5()
	// f6()
	// f7()
	f8()
}
