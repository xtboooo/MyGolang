package main

import (
	"fmt"
)

func add(a, b int) (c int) {
	c = a + b
	return
}

func calc(a, b int) (sum int, avg int) {
	sum = a + b
	avg = (a + b) / 2
	return
}

// 没有参数的 return 语句返回各个返回变量的当前值
func f1() {
	var a, b int = 1, 2
	c := add(a, b)
	sum, avg := calc(a, b)
	fmt.Println(a, b, c, sum, avg)
}

// Golang返回值不能用容器对象接收多返回值, 只能用多个变量，或 "_" 忽略
func f2() {
	test := func() (int, int) {
		return 1, 2
	}
	// s := make([]int, 2)
	// s = test() // Error: multiple-value test() in single-value context
	x, _ := test()
	println(x)
}

// 多返回值可直接作为其他函数调用实参
func f3() {
	test := func() (int, int) {
		return 1, 2
	}
	add := func(x, y int) int {
		return x + y
	}
	sum := func(n ...int) int {
		var x int
		for _, i := range n {
			x += i
		}
		return x
	}
	println(add(test()))
	println(sum(test()))
}

// 命名返回参数允许 defer 延迟调用通过闭包读取和修改
func f4() {
	add := func(x, y int) (z int) {
		defer func() {
			z += 100
		}()

		z = x + y
		return
	}
	println(add(1, 2))
}

// 显式 return 返回前，会先修改命名返回参数
func f5() {
	add := func(x, y int) (z int) {
		defer func() {
			println(z) // 输出: 203
		}()
		z = x + y
		return z + 200 // 执行顺序: (z = z + 200) -> (call defer) -> (return)
	}
	println(add(1, 2))
}

func main() {
	// f1() // 1 2 3 3 1
	// f2()
	// f3()
	// f4()
	f5()
}
