package main

import "fmt"

// for range
func f1() {
	a := [...]int{1, 2, 3}
	for i, v := range a {
		fmt.Printf("%v:%v\n", i, v)
	}
}

// 切片
func f2() {
	var s = []int{1, 2, 3}
	for _, v := range s {
		fmt.Printf("v: %v\n", v)
	}
}

//map
func f3() {
	m := make(map[string]string, 0)
	m["name"] = "tom"
	m["age"] = "20"
	m["email"] = "xtbo9719@gmail.com"
	for key, value := range m {
		fmt.Printf("%v:%v\n", key, value)
	}
}

//字符串
func f4() {
	s := "hello world"
	for _, v := range s {
		fmt.Printf("v: %c\n", v) // %c 打印当前的字符

	}
}

func main() {
	// f1()
	// f2()
	// f3()
	// f4()
}
