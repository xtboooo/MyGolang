package main

import "fmt"

func f1() {
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


func f2()  {
	a :=10
    b :=&a
    fmt.Printf("a:%d ptr:%p\n", a,&a)
    fmt.Printf("b:%p type:%T\n", b, b)
    fmt.Println(&b)
}

func main() {
	// f1()
	f2()

}
