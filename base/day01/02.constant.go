package main

import (
	"fmt"
)

func main() {
	// const PI float32 = 3.14
	// const PI2 = 3.1415
	// fmt.Printf("PI: %v\n", PI)

	// const (
	// 	a = 0
	// 	b = 1
	// )
	// fmt.Printf("a: %v\n", a)
	// fmt.Printf("b: %v\n", b)

	// const m, n = 3, 4
	// fmt.Printf("m: %v\n", m)
	// fmt.Printf("n: %v\n", n)

	// iota 每调用自增1
	// const (
	// 	a1 = iota
	// 	a2 = iota
	// 	a3 = iota
	// )
	// fmt.Printf("a1: %v\n", a1)
	// fmt.Printf("a2: %v\n", a2)
	// fmt.Printf("a3: %v\n", a3)

	// iota _ 跳过
	// const (
	// 	a1 = iota
	// 	_
	// 	a2 = iota
	// )
	// fmt.Printf("a1: %v\n", a1)
	// fmt.Printf("a2: %v\n", a2)

	// iota 中间插队
	const (
		a1=iota
		a2=100
		a3=iota
	)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)
}
