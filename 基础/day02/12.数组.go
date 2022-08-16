package main

import (
	"fmt"
)

// 数组定义
func f1() {
	var a1 [2]int
	var a2 [3]string
	var a3 [2]bool
	var a4 [2]float32
	fmt.Printf("a1: %T\n", a1)
	fmt.Printf("a2: %T\n", a2)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)
	fmt.Printf("a4: %v\n", a4)
}

// 数组的初始化
func f2() {
	var a1 = [3]int{1, 2}
	fmt.Printf("a1: %v\n", a1)
	var a2 = [2]string{"hello", "world"}
	fmt.Printf("a2: %v\n", a2)
	var a3 = [2]bool{true, false}
	fmt.Printf("a3: %v\n", a3)
}

// 数组的初始化 省略长度
func f3() {
	var a1 = [...]int{1, 2, 3}
	fmt.Printf("len(a1): %v\n", len(a1))
	var a2 = [...]string{"good", "golang"}
	fmt.Printf("len(a2): %v\n", len(a2))
}

// 指定索引来初始化数组
func f4() {
	var a1 = [...]int{0: 1, 2: 3}
	fmt.Printf("a1: %v\n", a1)
	var a2 = [...]bool{3: true, 5: false}
	fmt.Printf("a2: %v\n", a2)
}

// 数组的修改
func f5() {
	var a1 = [...]int{1, 2, 3}
	fmt.Printf("a1: %v\n", a1)
	a1[1] = 100
	fmt.Printf("a1: %v\n", a1)
}

func main() {
	// f1()
	// f2()
	// f3()
	// f4()
	f5()
}
