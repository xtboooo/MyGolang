package main

import "fmt"

func f1(){
	a := 100
	b := 200
	if a >= b {
		fmt.Printf("a: %v\n", a)
	}else{
		fmt.Printf("b: %v\n", b)
	}
}

func main() {
	f1()
}