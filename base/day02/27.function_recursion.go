package main

import "fmt"

// 阶乘 for
func f1() {
	var s int = 1
	for i := 1; i <= 5; i++ {
		s *= i
	}
	fmt.Printf("s: %v\n", s)
}

// 阶乘 递归
func f2(a int) int {
	if a == 1 {
		return 1
	} else {
		return a * f2(a-1)
	}
}

// 斐波那契数列
func f3(n int) int {
	if n == 1 || n == 2{
		return 1
	}
	return f3(n-1) + f3(n-2)
}

func main() {
	// f1()
	// i := f2(6)
	// fmt.Printf("i: %v\n", i)
	j := f3(7)
	fmt.Printf("j: %v\n", j)
}
