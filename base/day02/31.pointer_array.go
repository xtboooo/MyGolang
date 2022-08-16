package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	// 指向数组的指针
	var pa [3]*int
	for i := 0; i < len(pa); i++ {
		pa[i] = &a[i]
	}
	fmt.Printf("pa: %v\n", pa)

	for i := 0; i < len(pa); i++ {
		fmt.Printf("%v\n", pa[i])
	}
}
