package main

import "fmt"

// 面向对象 属性和方法分开来写

type Person2 struct {
	name string
}

// (per Person) 接收者 receiver
func (per Person2) eat() {
	fmt.Printf("%v, eat...\n", per.name)
}

func (per Person2) sleep() {
	fmt.Printf("%v, sleep...\n", per.name)
}

type Customer1 struct {
	name string
}

func (customer Customer1) login(name string, pwd string) bool {
	fmt.Printf("customer.name: %v\n", customer.name)
	if name == "tom" && pwd == "123456" {
		return true
	} else {
		return false
	}
}

func f1() {
	per := Person2{
		name: "tom",
	}
	per.eat()   // tom, eat...
	per.sleep() // tom, sleep...

	cus := Customer1{
		name: "tom",
	}
	r := cus.login("tom", "123456")
	fmt.Printf("r: %v\n", r) // r: true
}

// Person 结构体
type Person4 struct {
	name string
	age  int8
}

// NewPerson 构造函数
func NewPerson(name string, age int8) *Person4 {
	return &Person4{
		name: name,
		age:  age,
	}
}

// Dream Person做梦的方法
func (p Person4) Dream() {
	fmt.Printf("%s的梦想是学好Go语言!\n", p.name)
}

func f2() {
	p1 := NewPerson("测试", 25)
	p1.Dream()
}

func main() {
	// f1()
	f2()
}
