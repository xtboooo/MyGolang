package main

import "fmt"

// goto
func f1() {
	i := 1
	if i >= 2 {
		fmt.Println("2")
	} else {
		goto END
	}
END:
	fmt.Println("end...")
}

// goto 双层循环
func f2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i >= 2 && j >= 2 {
				goto END
			}
			fmt.Printf("%v:%v\n", i, j)
		}
	}
END:
	fmt.Println("end...")
}

func main() {
	// f1()
	f2()
}
