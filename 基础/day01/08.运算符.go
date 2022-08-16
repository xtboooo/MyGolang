package main

import "fmt"

func main() {
	// + - * / %
	a := 10
	b := 20
	fmt.Printf("(a + b): %v\n", (a + b))
	fmt.Printf("(a - b): %v\n", (a - b))
	fmt.Printf("(a * b): %v\n", (a * b))
	fmt.Printf("(a / b): %v\n", (a / b))
	fmt.Printf("(a %% b): %v\n", (a % b))

	// ++ --
	c := 100
	c++
	fmt.Printf("c: %v\n", c)
	c--
	fmt.Printf("c: %v\n", c)

	// <  > == >= <= !=
	r := a == b
	fmt.Printf("r: %v\n", r)
	r = a > b
	fmt.Printf("r: %v\n", r)
	r = a < b
	fmt.Printf("r: %v\n", r)
	r = a >= b
	fmt.Printf("r: %v\n", r)
	r = a <= b
	fmt.Printf("r: %v\n", r)
	r = a != b
	fmt.Printf("r: %v\n", r)
	// && || !
	d := true
	e := false
	f := d && e
	fmt.Printf("f: %v\n", f)
	f = d || e
	fmt.Printf("f: %v\n", f)
	fmt.Printf("d: %v\n", !d)

	// 位运算 & | ^ << >>
	g := 4 // 0100
	fmt.Printf("g: %b\n", g)
	h := 8 // 1000
	fmt.Printf("h: %b\n", h)
	fmt.Printf("(g & h): %b\n", (g & h))
	fmt.Printf("(g | h): %b\n", (g | h))
	fmt.Printf("(g ^ h): %b\n", (g ^ h))
	fmt.Printf("(g << 2): %b\n", (g << 2))
	fmt.Printf("(g >> 2): %b\n", (g >> 2))

	// 赋值
	j := 200
	j = 300
	fmt.Printf("j: %v\n", j)
	j += 100 // j = j +100
	fmt.Printf("j: %v\n", j)
}
