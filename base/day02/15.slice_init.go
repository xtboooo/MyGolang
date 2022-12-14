package main

import "fmt"

// 切片
func f1() {
	var s1 = []int{1, 2, 3, 4, 5, 6}
	var s2 = s1[0:3]
	fmt.Printf("s2: %v\n", s2)
	s3 := s1[3:]
	fmt.Printf("s3: %v\n", s3)
	s4 := s1[2:5]
	fmt.Printf("s4: %v\n", s4)
	s5 := s1[:]
	fmt.Printf("s5: %v\n", s5)
}

// 数组
func f2() {
	var a1 = [...]int{1, 2, 3, 4, 5, 6}
	a2 := a1[:3]
	fmt.Printf("a2: %v\n", a2)
	a3 := a1[3:]
	fmt.Printf("a3: %v\n", a3)
	a4 := a1[:]
	fmt.Printf("a4: %v\n", a4)
}

// 空切片
func f3() {
	var a1 []int
	fmt.Printf("(a1 == nil): %v\n", (a1 == nil))
	fmt.Printf("len(a1): %v\n", len(a1))
	fmt.Printf("cap(a1): %v\n", cap(a1))
}

func main() {
	// f1()
	// f2()
	f3()
}
