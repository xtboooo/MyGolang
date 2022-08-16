package main

import "fmt"

// for break
func f1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i >= 5 {
			break
		}
	}
}

//switch break
func f2() {
	i := 2
	switch i {
	case 1:
		fmt.Println("1")
		break
	case 2:
		fmt.Println("2")
		// break
		fallthrough
	case 3:
		fmt.Println("3")
	}
}

//跳转到某一个标签
func f3() {
	MYLABEL:
		for i := 0; i < 10; i++ {
			fmt.Printf("i: %v\n", i)
			if i >= 5 {
				break MYLABEL
			}
		}
		fmt.Println("end...")
}
func main() {
	// f1()
	// f2()
	f3()
}
