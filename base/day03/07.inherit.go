package main

import "fmt"

type Animal struct {
	name string
	age  int
}

func (a Animal) eat() {
	fmt.Println("eat...")
}

func (a Animal) sleep() {
	fmt.Println("sleep...")
}

type Dog4 struct {
	Animal // 可以理解为继承关系
	color  string
}

type Cat4 struct {
	Animal
	bbb string
}

func main() {
	dog := Dog4{
		Animal{name: "大黄", age: 2}, // 匿名
		"黑色",
	}
	dog.eat()                                // eat...
	dog.sleep()                              // sleep...
	fmt.Printf("dog.color: %v\n", dog.color) // dog.color: 黑色
	fmt.Printf("dog.name: %v\n", dog.name)   // dog.name: 大黄

	cat := Cat4{
		Animal{name: "小花", age: 2},
		"bbb",
	}
	cat.eat()                              // eat...
	cat.sleep()                            // sleep...
	fmt.Printf("cat.name: %v\n", cat.name) // cat.name: 小花
	fmt.Printf("cat.bbb: %v\n", cat.bbb)   // cat.bbb: bbb
}
