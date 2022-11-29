package main

import "fmt"

func length(s string) int {
	println("call length.")
	return len(s)
}
func f1() {
	s := "abcd"
	for i, n := 0, length(s); i < n; i++ { // 避免多次调用 length 函数。
		println(i, s[i])
	}
}

func f2() {
	var b int = 15
	var a int
	numbers := [6]int{1, 2, 3, 5}
	/* for 循环 */
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}
	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
}

func f3() {
	/* 定义局部变量 */
	var i, j int

	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break // 如果发现因子，则不是素数
			}
		}
		if j > (i / j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
}

func main() {
	// f1()
	// f2()
	f3()
}
