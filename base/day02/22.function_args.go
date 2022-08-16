package main

import (
	"fmt"
)

// 形参
func sum(a int, b int) int {
	return a + b
}

// 参数传递的是copy
func f1(a int) {
	a = 100
}

// 切片传递 可能会导致值被修改
func f2(s []int) {
	s[0] = 200
}

// 可变参数
func f3(args ...int) {
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
}

// 固定参数 + 可变参数
func f4(name string, od bool, args ...int) {
	fmt.Printf("name: %v\n", name)
	fmt.Printf("od: %v\n", od)
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
}

func main() {
	r := sum(1, 2) // 实参
	fmt.Printf("r: %v\n", r)
	x := 200
	f1(x)
	fmt.Printf("x: %v\n", x)
	s := []int{1, 2, 3}
	f2(s)
	fmt.Printf("s: %v\n", s)
	f3(1, 2, 3, 4, 5, 6)
	f4("tom", true, 1, 2, 3, 4, 5)
}
