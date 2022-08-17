package main

import (
	"fmt"
)

type Pet2 interface {
	eat4()
	sleep1()
}

type Dog2 struct {
}

type Cat2 struct {
}

// dog 实现 Pet2 接口
func (dog Dog2) eat4() {
	fmt.Println("dog2 eat4...")
}

func (dog Dog2) sleep1() {
	fmt.Println("dog2 sleep...")
}

// cat 实现 Pet2 接口
func (cat Cat2) eat4() {
	fmt.Println("cat2 eat4...")
}

func (cat Cat2) sleep1() {
	fmt.Println("cat2 sleep...")
}

type Person1 struct {
}

// pet 既可以传递Dog2 也可以传递Cat2
func (person Person1) care(pet Pet2) {
	pet.eat4()
	pet.sleep1()
}

func main() {
	dog := Dog2{}
	cat := Cat2{}
	per := Person1{}
	per.care(dog)
	per.care(cat)
	/* 	dog2 eat4...
	   	dog2 sleep...
	   	cat2 eat4...
	   	cat2 sleep... */
}
