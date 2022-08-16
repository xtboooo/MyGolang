package main

import "fmt"

func main() {
	type Dog struct {
		name  string
		age   int
		color string
	}

	type Person struct {
		dog  Dog
		name string
		age  int
	}

	dog := Dog{
		name:  "大黄",
		age:   2,
		color: "black",
	}

	per := Person{
		dog:  dog,
		age:  12,
		name: "tom",
	}

	fmt.Printf("per: %v\n", per)                   // per: {{大黄 2 black} tom 12}
	fmt.Printf("per.dog.name: %v\n", per.dog.name) // per.dog.name: 大黄
}
