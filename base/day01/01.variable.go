package main

import (
	"fmt"
)

func getNameAndAge() (string, int) {
	return "f", 20
}

func main() {
	// var 变量声明
	// var name string
	// var age int
	// var gender string

	// var批量声明
	// var (
	//     name string
	//     age int
	//     gender string
	// )
	// name = "a"
	// age = 12
	// gender = "male"

	// 变量的初始化
	// var name string = "b"
	// var age int = 13
	// var gender string = "female"

	// 类型推断
	// var name = "c"
	// var age = 14
	// var gender = "male"

	// 批量初始化
	// var name, age, gender = "d", 16, "female"

	// 短变量声明 := , 仅在函数中才能用
	// name := "e"
	// age := 17
	// gender := "male"

	// fmt.Printf("name: %v\n", name)
	// fmt.Printf("age: %v\n", age)
	// fmt.Printf("gender: %v\n", gender)

	// 匿名变量
	_, age := getNameAndAge()
	// fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
}
