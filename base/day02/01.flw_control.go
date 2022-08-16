package main

import "fmt"

func main() {
	a := 200
	if a > 20 {
		fmt.Println("a")
	} else {
		fmt.Println("b")
	}

	switch a {
	case 100:
		fmt.Println("c")
	default:
		fmt.Println("d")
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
	}

	x := [...]int{1, 2, 3}
	for _, v := range x {
		fmt.Printf("v: %v\n", v)
	}
}
