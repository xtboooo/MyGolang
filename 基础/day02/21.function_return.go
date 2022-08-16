package main

import (
	"fmt"
)

func f1() {
	fmt.Println("既没有参数, 也没有返回值")
}

func sum1(a int, b int) (ret int) {
	ret = a+b
	return ret
}

func sum2(a int, b int) int {
	ret := a+b
	return ret
}

func sum3(a int, b int) int {
	// ret := a+b
	return a+b
}

func f2() (name string, age int){
	name ="tom"
	age = 20
	return name, age
}

func f3() (string, int){
	n :="bob"
	m := 18
	return n, m
}


func main() {
	f1()
	ans1 := sum1(1,2)
	fmt.Printf("ans1: %v\n", ans1)
	ans2 := sum2(1,2)
	fmt.Printf("ans2: %v\n", ans2)
	ans3 := sum3(1,2)
	fmt.Printf("ans3: %v\n", ans3)
	name, age := f2()
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	n, m := f3()
	fmt.Printf("name: %v\n", n)
	fmt.Printf("age: %v\n", m)
}