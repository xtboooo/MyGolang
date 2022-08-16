package main

import "fmt"

func main() {
	/* 	// 类型定义
	   	type MyInt int
	   	var i MyInt
	   	i = 100
	   	fmt.Printf("i: %T,%v\n", i, i) // i: main.MyInt,100 */

	// 类型别名
	type MyInt = int
	var i MyInt
	i = 100
	fmt.Printf("i: %T,%v\n", i, i) // i: int,100
}
