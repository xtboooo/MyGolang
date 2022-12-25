package main

import "fmt"

/*
  - 类型 T 方法集包含全部 receiver T 方法。

• 类型 *T 方法集包含全部 receiver T + *T 方法。
• 如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。
• 如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T + *T 方法。
• 不管嵌入 T 或 *T，*S 方法集总是包含 T + *T 方法。
*/

type S struct {
	T
}

type T struct {
	int
}

func (t T) testT() {
	fmt.Println("类型 T1 方法集包含全部 receiver T 方法。")
}

func (t *T) testP() {
	fmt.Println("类型 *T 方法集包含全部 receiver *T 方法。")
}

// 类型 T 方法集包含全部 receiver T 方法。
func f1() {
	t1 := T{1}
	fmt.Printf("t1 is : %v\n", t1)
	t1.testT()
}

// 类型 *T 方法集包含全部 receiver T + *T 方法
func f2() {
	t1 := T{1}
	t2 := &t1
	fmt.Printf("t2 is : %v\n", t2)
	t2.testT()
	t2.testP()
}

// 类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。
func f3() {
	s1 := S{T{1}}
	s2 := &s1
	fmt.Printf("s1 is : %v\n", s1)
	s1.testT()
	fmt.Printf("s2 is : %v\n", s2)
	s2.testT()
}

// 类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T + *T 方法
func f4() {
	s1 := S{T{1}}
	s2 := &s1
	fmt.Printf("s1 is : %v\n", s1)
	s1.testT()
	s1.testP()
	fmt.Printf("s2 is : %v\n", s2)
	s2.testT()
	s2.testP()
}

func main() {
	// f1()
	// f2()
	// f3()
	f4()
}
