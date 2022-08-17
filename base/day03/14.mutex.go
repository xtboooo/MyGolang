package main

import (
	"fmt"
	"sync"
	"time"
)

var i int = 100

var wp1 sync.WaitGroup

var lock sync.Mutex

func add() {
	defer wp1.Done()
	lock.Lock()
	i += 1
	fmt.Printf("i++: %v\n", i)
	time.Sleep(time.Millisecond * 10)
	lock.Unlock()
}

func sub() {
	lock.Lock()
	defer wp1.Done()
	i -= 1
	fmt.Printf("i--: %v\n", i)
	time.Sleep(time.Millisecond * 2)
	lock.Unlock()
}

func main() {
	for i := 0; i < 100; i++ {
		wp1.Add(1)
		go add()
		wp1.Add(1)
		go sub()
	}
	wp1.Wait()
	fmt.Printf("end i: %v\n", i)
}
