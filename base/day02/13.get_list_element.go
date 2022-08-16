package main

import "fmt"

func f1() {
	var a1 [2]int
	a1[0] = 100
	a1[1] = 200
	fmt.Printf("a1[0]: %v\n", a1[0])
	fmt.Printf("a1[1]: %v\n", a1[1])
	fmt.Println("-----------")
	a1[0] = 1000
	a1[1] = 2000
	fmt.Printf("a1[0]: %v\n", a1[0])
	fmt.Printf("a1[1]: %v\n", a1[1])

	// 数组长度越界
	// a1[3] = 3000
}

// 数组的长度
func f2() {
	var a1 = [3]int{1, 2, 3}
	fmt.Printf("len(a1): %v\n", len(a1))
	var a2 = [...]int{1, 2, 3, 4, 5}
	fmt.Printf("len(a2): %v\n", len(a2))
}

// 数组的遍历 根据长度和下标
func f3() {
	var a1 = [...]int{1, 2, 3}
	for i := 0; i < len(a1); i++ {
		fmt.Printf("a1[%v]: %v\n", i, a1[i])
	}
}

// 数组的遍历 for range
func f4() {
	var a1 = [...]int{1, 2, 3}
	for i, v := range a1 {
		fmt.Printf("a1[%v]: %v\n", i, v)
	}
}

func main() {
	// f1()
	// f2()
	// f3()
	f4()
}
