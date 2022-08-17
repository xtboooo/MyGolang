package main

import "fmt"

type Pet interface {
	eat(string) string
	eat2(string) string
}

type Dog struct {
	name string
}
// 参数传值 是拷贝
func (dog Dog) eat(name string) string {
	dog.name = name
	fmt.Printf("dog.name: %v\n", dog.name) // dog.name: 小黑
	return "吃的很好"
}

// 参数传指针
func (dog *Dog) eat2(name string) string {
	dog.name = name
	fmt.Printf("dog.name: %v\n", dog.name) // dog.name: 小黑
	return "吃的很好"
}


func main() {
	dog := Dog{
		name: "大黄",
	}
	s:= dog.eat("小黑")
	fmt.Printf("s: %v\n", s)
	fmt.Printf("dog.name: %v\n", dog.name) // dog.name: 大黄


	dog1 := &Dog{
		name: "大黄",
	}
	s = dog1.eat2("小黑")
	fmt.Printf("s: %v\n", s)
	fmt.Printf("dog.name: %v\n", dog1.name) // dog.name: 小黑
}