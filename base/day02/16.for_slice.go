package main

import "fmt"

// for
func f1() {
	var a1 = []int{1, 2, 3, 4, 5, 6}
	l := len(a1)
	for i := 0; i < l; i++ {
		fmt.Printf("a1[%v]: %v\n", i, a1[i])
	}
}

// for range
func f2() {
	var a1 = []int{1, 2, 3, 4, 5, 6}
	for i, v := range a1 {
		fmt.Printf("i: %v\n", i)
		fmt.Printf("v: %v\n", v)
	}
}

func main() {
	// f1()
	f2()
}
