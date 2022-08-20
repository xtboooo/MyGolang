package main

import (
	"crypto/rand"
	"fmt"
	"math"
)


func testMath()  {
	fmt.Printf("math.Pi: %v\n", math.Pi)	
	fmt.Printf("math.MaxInt16: %v\n", math.MaxInt16)
	fmt.Printf("math.Ceil(3.14): %v\n", math.Ceil(3.14))
}

func main() {
	// testMath()
	fmt.Printf("rand.Int: %d\n", rand.Int)
}