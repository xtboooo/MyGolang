package main

import "fmt"

func sayHello(name string) {
	fmt.Printf("hello, %s\n", name)
}

// 函数作为参数
func f1(name string, f func(string)) {
	f(name)
}


// 函数作为返回值
func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func cal(operator string) func(int, int) int {
	switch operator {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func main() {
	f1("bob", sayHello)
	ff := cal("+")
	r := ff(1, 2)
	fmt.Printf("r: %v\n", r)
	ff = cal("-")
	r = ff(2, 1)
	fmt.Printf("r: %v\n", r)
}
