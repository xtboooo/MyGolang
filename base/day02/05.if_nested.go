package main

import "fmt"

func f1() {
	a := 100
	b := 200
	c := 300
	if a > b {
		if a > c {
			fmt.Printf("a: %v\n", a)
		}
	} else {
		if b > c {
			fmt.Printf("b: %v\n", b)
		} else {
			fmt.Printf("c: %v\n", c)
		}
	}
}


func f2(){
	gender:= "femal"
	age:=15
	if gender == "male"{
		if age >=18{
			fmt.Println("成年男性")
		}else{
			fmt.Println("未成年男性")
		}
	}else{
		if age >= 18{
			fmt.Println("成年女性")
		}else{
			fmt.Println("未成年女性")
		}
	}
}

func main() {
	f1()
	f2()
}
