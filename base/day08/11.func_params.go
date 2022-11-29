package main

import (
	"fmt"
)

/*
	注意1：无论是值传递，还是引用传递，传递给函数的都是变量的副本，不过，值传递是值的拷贝。引用传递是地址的拷贝，一般来说，地址拷贝更为高效。而值拷贝取决于拷贝的对象大小，对象越大，则性能越低。

注意2：map、slice、chan、指针、interface默认以引用的方式传递。

不定参数传值 就是函数的参数不是固定的，后面的类型是固定的。（可变参数）

Golang 可变参数本质上就是 slice。只能有一个，且必须是最后一个。

在参数赋值时可以不用用一个一个的赋值，可以直接传递一个数组或者切片，特别注意的是在参数后加上“…”即可。
*/

func f1() {
	test := func(s string, n ...int) string {
		var x int
		for _, i := range n {
			x += i
		}
		return fmt.Sprintf(s, x)
	}
	println(test("sum: %d", 1, 2, 3))
}

// 使用 slice 对象做变参时，必须展开
func f2() {
	test := func(s string, n ...int) string {
		var x int
		for _, i := range n {
			x += i
		}
		return fmt.Sprintf(s, x)
	}
	s := []int{1, 2, 3}
	println(test("sum: %d", s...))
}

func main() {
	// f1()
	f2()
}
