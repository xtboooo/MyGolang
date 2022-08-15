package main

import (
	"fmt"
	"strings"
)

func main() {
	/* 	s := "hello world"
	   	fmt.Printf("s: %v\n", s)
	   	// `` 多行字符串
	   	s2 := `
	   	line 1
	   	line 2
	   	line 3
	   	`
	   	fmt.Printf("s2: %v\n", s2) */

	// 字符串的连接
	/* 	s1 := "s1"
	   	s2 := "s2"
	   	msg := s1+s2
	   	fmt.Printf("msg: %v\n", msg) */

	// Sprintf
	/* 	s3 := "s3"
	   	s4 := "s4"
	   	msg:= fmt.Sprintf("%s,%s", s3, s4)
	   	fmt.Printf("msg: %v\n", msg) */

	// string.Join
	/* 	name := "tom"
	   	age :="20"
	   	msg := strings.Join([]string{name, age}, ",")
	   	fmt.Printf("mas: %v\n", msg) */

	// buffer
	/* 	var buffer bytes.Buffer
	   	buffer.WriteString("tom")
	   	buffer.WriteString(",")
	   	buffer.WriteString("20")
	   	fmt.Printf("%v\n", buffer.String()) */

	/* 	// 转义字符
	   	hello :="hello"
	   	print(hello + "\n")
	   	print(hello)

	   	s := "hello \t world"
	   	fmt.Printf("s: %v\n", s)

	   	s:= "c:\\programs file\\a"
	   	fmt.Printf("s: %v\n", s) */

	/* 	// 字符串的切片操作
	   	s := "hello world"
	   	a := 2
	   	b := 5
	   	fmt.Printf("s[a]: %c\n", s[a])
	   	fmt.Printf("s[a:b]: %v\n", s[a:b])
	   	fmt.Printf("s[a:]: %v\n", s[a:])
	   	fmt.Printf("s[0:b]: %v\n", s[0:b]) */

	// 字符串函数 len
	s := "Hello World" 
	fmt.Printf("len(s): %v\n", len(s))
	fmt.Printf("strings.Split(s, \" \"): %v\n", strings.Split(s, " "))
	fmt.Printf("strings.Contains(\"hello\"): %v\n", strings.Contains(s, "hello"))
	fmt.Printf("strings.ToLower(s): %v\n", strings.ToLower(s))
	fmt.Printf("strings.ToUpper(s): %v\n", strings.ToUpper(s))
	fmt.Printf("strings.HasPrefix(s, \"hello\"): %v\n", strings.HasPrefix(s, "Hello"))
	fmt.Printf("strings.HasSuffix(s, \"World\"): %v\n", strings.HasSuffix(s, "world"))
	fmt.Printf("strings.Index(s, \"ll\"): %v\n", strings.Index(s, "ll"))
}

