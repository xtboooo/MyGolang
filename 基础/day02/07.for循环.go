package main

import "fmt"

// for
func f1() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
}

// 变量在外面定义
func f2() {
	i := 1
	for ; i <= 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
}

// 变量在外面定义, ++在for外做
func f3() {
	i := 1
	for i <= 10 {
		fmt.Printf("i: %v\n", i)
		i++
	}
}

// 无限循环
func f4() {
	for {
		fmt.Println("run...")
	}
}

func main() {
	f1()
	f2()
	f3()

}
