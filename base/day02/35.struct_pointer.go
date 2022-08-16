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

func main() {
	f1()
	f2()
	f3()
}
