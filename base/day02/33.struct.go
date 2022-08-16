package main

import "fmt"

type Person struct {
	id     int
	name   string
	age    int
	gender string
}

type Customer struct {
	id, age      int
	name, gender string
}

func main() {
	var bob Person
	fmt.Printf("bob: %v\n", bob) // bob: {0  0 }

	var kite Customer
	fmt.Printf("kite: %v\n", kite) // kite: {0 0  }

	var tom Person
	tom.age = 18
	tom.gender = "male"
	tom.id = 1
	tom.name = "tom"
	fmt.Printf("tom: %v\n", tom)           // tom: {1 tom 18 male}
	fmt.Printf("tom.name: %v\n", tom.name) // tom.name: tom

	var linda struct {
		id     int
		name   string
		gender string
	}
	linda.gender = "female"
	linda.id = 2
	linda.name = "linda"
	fmt.Printf("linda: %v\n", linda) // linda: {2 linda female}
}
