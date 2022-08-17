package main

import "fmt"

type Person2 struct {
	name string
	age  int
}

func (per Person2) eat() {
	fmt.Println("eat...")
}
func (per Person2) sleep() {
	fmt.Println("sleep...")
}

func (per Person2) work() {
	fmt.Println("work....")
}

func main() {
	per := Person2{
		name: "tom",
		age:  20,
	}
	per.eat()
	per.sleep()
	per.work()
	/* 	eat...
	   	sleep...
	   	work.... */
}
