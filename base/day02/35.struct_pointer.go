package main

import (
	"fmt"
)

func f1() {
	var name string
	name = "tom"
	var p_name *string
	p_name = &name

	fmt.Printf("name: %v\n", name)      // name: tom
	fmt.Printf("p_name: %v\n", p_name)  // p_name: 0xc000054250
	fmt.Printf("p_name: %v\n", *p_name) // p_name: tom
}

func f2() {
	type Person struct {
		id   int
		name string
		age  int
	}
	tom := Person{
		5,
		"tom",
		20,
	}
	var p_person *Person
	p_person = &tom
	fmt.Printf("tom: %v\n", tom)            // tom: {5 tom 20}
	fmt.Printf("p_person: %p\n", p_person)  // p_person: 0xc0000563e0
	fmt.Printf("p_person: %v\n", *p_person) // p_person: {5 tom 20}
}

func f3() {
	type Person struct {
		id   int
		name string
		age  int
	}
	var tom = new(Person)
	// fmt.Printf("tom: %p\n", tom) // tom: 0xc000056440
	tom.name = "tom"
	tom.age = 21
	tom.id = 10
	fmt.Printf("tom: %v\n", *tom) // tom: {10 tom 21}
}

// 结构体内存布局
func f4() {
	type test struct {
		a int8
		b int8
		c int8
		d int8
	}
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a) // n.a 0xc000124004
	fmt.Printf("n.b %p\n", &n.b) // n.b 0xc000124005
	fmt.Printf("n.c %p\n", &n.c) // n.c 0xc000124006
	fmt.Printf("n.d %p\n", &n.d) // n.d 0xc000124007
}

func f5() {
	type student struct {
		name string
		age  int
	}
	m := make(map[string]*student)
	students := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}

	for _, stu := range students {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}

func main() {
	// f1()
	// f2()
	// f3()
	// f4()
	f5()
}
