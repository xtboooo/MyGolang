package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	/* 	counter := 0
	   	for _ = range ticker.C {
	   		counter ++
	   		if counter > 5{
	   			ticker.Stop()
	   			break
	   		}
	   		fmt.Println("ticker...")
	   	} */

	chanInt1 := make(chan int)

	go func() {
		for _ = range ticker.C {
			select {
			case chanInt1 <- 1:
			case chanInt1 <- 2:
			case chanInt1 <- 3:
			}
		}
	}()

	sum := 0
	for v := range chanInt1 {
		fmt.Printf("v: %v\n", v)
		sum += v
		if sum >= 10 {
			break
		}
	}

}
