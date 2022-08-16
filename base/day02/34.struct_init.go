package main

import "fmt"

func main() {
	type Person struct {
		name   string
		id     int
		age    int
		gender string
	}

	// 初始化1
	var bob Person
	bob = Person{
		id:     3,
		name:   "bob",
		gender: "male",
		age:    20,
	}
	fmt.Printf("bob: %v\n", bob) // bob: {bob 3 20 male

	// 初始化2
	var tom Person
	tom = Person{
		"tom",
		4,
		21,
		"male",
	}
	fmt.Printf("tom: %v\n", tom) // tom: {tom 4 21 male}

	// 初始化3 部分初始化
	var a Person
	a = Person{
		id:   5,
		name: "a",
	}
	fmt.Printf("a: %v\n", a) // a: {a 5 0 }
}
