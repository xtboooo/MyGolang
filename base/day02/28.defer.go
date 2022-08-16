package main

import "fmt"

// stack
func f1() {
	fmt.Println("start...")
	defer fmt.Println("step1...")
	defer fmt.Println("step2...")
	defer fmt.Println("step3...")
	fmt.Println("end...")
}

func main() {
	f1()
	/* 
	start...
	end...
	step3...
	step2...
	step1...
	*/
}