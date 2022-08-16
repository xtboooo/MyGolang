package main

import "fmt"

// add
func f1() {
	var s1 = []int{}
	s1 = append(s1, 100)
	s1 = append(s1, 200)
	s1 = append(s1, 300)
	fmt.Printf("s1: %v\n", s1)
}

// delete
func f2() {
	var s1 = []int{1, 2, 3, 4, 5}
	// 删除索引为2的元素
	s1 = append(s1[:2], s1[3:]...)
	fmt.Printf("s1: %v\n", s1)
}

// update
func f3() {
	var s1 = []int{1, 2, 3, 4, 5}
	s1[2] = 100
	fmt.Printf("s1: %v\n", s1)
}

// query
func f4() {
	var s1 = []int{1, 2, 3, 4, 5}
	key := 2
	for i, v := range s1 {
		if v == key {
			fmt.Printf("i: %v\n", i)
		}
	}
}

// 深浅拷贝	
func f5() {
	var s1 = []int{1, 2, 3, 4, 5}
	// s2 := s1 // 浅
	var s2 = make([]int, 5)
	copy(s2, s1) // 深
	s1[0] = 100
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
}

func main() {
	// f1()
	// f2()
	// f3()
	// f4()
	f5()
}
