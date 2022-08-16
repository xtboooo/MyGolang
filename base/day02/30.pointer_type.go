package main

import "fmt"

func main() {
	var ip *int
	fmt.Printf("ip: %v\n", ip) // nil
	fmt.Printf("ip: %T\n", ip)

	var i int = 100
	ip = &i
	fmt.Printf("ip: %v\n", ip)
	fmt.Printf("ip: %v\n", *ip)

	var sp *string
	var s = "hello world"
	sp = &s
	fmt.Printf("sp: %T\n", sp)
	fmt.Printf("sp: %v\n", *sp)

	var bp *bool
	b := true
	bp = &b
	fmt.Printf("bp: %T\n", bp)
	fmt.Printf("bp: %v\n", *bp)
}
