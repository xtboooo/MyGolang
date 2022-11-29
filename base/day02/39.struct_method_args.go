package main

import "fmt"

func foo() {
	type Person3 struct {
		name string
	}
	p1 := Person3{
		name: "tom",
	}

	p2 := &Person3{
		name: "tom",
	}

	fmt.Printf("p1: %T\n", p1) // p1: main.Person3 值类型
	fmt.Printf("p2: %T\n", p2) // p2: *main.Person3 指针类型
}

type Person4 struct {
	name string
}

func showPerson1(per Person4) {
	per.name = "tom ..."
}

func showPerson3(per *Person4) {
	// per.name 自动解引用  (*per).name = "tom .. "
	per.name = "tom ..."
}

func (per Person4) showPerson5() {
	per.name = "tom ..."
}
func (per *Person4) showPerson6() {
	// per.name 自动解引用  (*per).name = "tom .. "
	per.name = "tom ..."

}
func main() {
	p1 := Person4{
		name: "tom",
	}
	p2 := &Person4{
		name: "tom",
	}

	/* 	showPerson1(p1)
	   	fmt.Printf("p1: %v\n", p1) // p1: {tom}
	   	fmt.Println("----------")
	   	showPerson3(p2)
	   	fmt.Printf("p2: %v\n", *p2) // p2: {tom .. } */

	p1.showPerson5()
	fmt.Printf("p1: %v\n", p1) // p1: {tom}
	fmt.Println("----------")
	p2.showPerson6()
	fmt.Printf("p2: %v\n", *p2) // p2: {tom .. }
}
