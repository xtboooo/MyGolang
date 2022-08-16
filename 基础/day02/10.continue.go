package main

import "fmt"

// continue
func f1() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("i: %v\n", i)
		} else {
			continue
		}
	}
}

// 跳转到某一个标签
func f2() {
	for i := 0; i < 10; i++ {
	MYLABEL:
		for j := 0; j < 10; j++ {
			if i == 2 && j == 2 {
				continue MYLABEL
			}
			fmt.Printf("%v:%v\n", i, j)
		}

	}
}

func main() {
	// f1()
	f2()
}
