package main

import "fmt"

type Person1 struct {
	id   int
	name string
	age  int
}

// 函数传参 是拷贝
func showPerson(person Person1) {
	person.age = 20
	person.name = "kite"
	person.id = 5
	fmt.Printf("person: %v\n", person) // person: {5 kite 20}
}

// 传递指针 值被修改
func showPerson2(person *Person1) {
	person.id = 8
	person.name = "kite"
	person.age = 24
	fmt.Printf("person: %v\n", person) // person: &{8 kite 24}
}

func main() {
	tom := Person1{	
		6,
		"tom",
		18,
	}
	fmt.Printf("tom: %v\n", tom) // tom: {6 tom 18}
	fmt.Println("---------")
	showPerson(tom)
	fmt.Printf("tom: %v\n", tom) // {6 tom 18}

	per := &tom
	fmt.Printf("per: %v\n", per) // per: &{6 tom 18}
	showPerson2(per)
	fmt.Printf("per: %v\n", *per) // per: {8 kite 24}
}
