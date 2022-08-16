package main

import "fmt"

func f1() {
	m1 := map[string]string{"name": "bob", "age": "20", "gender": "male"}
/* 	for k := range m1 {
		fmt.Printf("k: %v\n", k)
	} */
	for k, v := range m1 {
		fmt.Printf("%v:%v\n", k, v)
	}
}

func main() {
	f1()
}
