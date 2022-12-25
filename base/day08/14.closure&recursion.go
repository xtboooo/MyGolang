package main

import "fmt"

// 闭包demo
func f1() {
	a := func() func() int {
		i := 0
		b := func() int {
			i++
			fmt.Println(i)
			return i
		}
		return b
	}
	c := a()
	c()
	c()
	c()
	a() //不会输出i
}

// 闭包复制的是原对象指针
func f2() {
	test := func() func() {
		x := 100
		fmt.Printf("x (%p) = %d\n", &x, x)

		return func() {
			fmt.Printf("x (%p) = %d\n", &x, x)
		}
	}
	test()()
}

// 外部引用函数参数局部变量
func f3() {
	add := func(base int) func(int) int {
		return func(i int) int {
			base += i
			return base
		}
	}
	tmp1 := add(10)
	fmt.Println(tmp1(1), tmp1(2))
	// 此时tmp1和tmp2不是一个实体了
	tmp2 := add(100)
	fmt.Println(tmp2(1), tmp2(2))
}

// 返回2个闭包
func f4() {
	// 返回2个函数类型的返回值
	test01 := func(base int) (func(int) int, func(int) int) {
		// 定义2个函数，并返回
		// 相加
		add := func(i int) int {
			base += i
			return base
		}
		// 相减
		sub := func(i int) int {
			base -= i
			return base
		}
		// 返回
		return add, sub
	}

	f1, f2 := test01(10)
	// base一直是没有消
	fmt.Println(f1(1), f2(2))
	// 此时base是9
	fmt.Println(f1(3), f2(4))
}

func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}

func f5() {
	var i int = 7
	fmt.Printf("Factorial of %d is %d\n", i, factorial(i))
}

func fibonaci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
}

func f6() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\n", fibonaci(i))
	}
}

func main() {
	// f1()
	// f2()
	// f3()
	// f4()
	// f5()
	f6()
}
