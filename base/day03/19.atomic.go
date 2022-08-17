package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

/* var i = 100

var lock1 sync.Mutex

func add1() {
	lock1.Lock()
	i++
	lock1.Unlock()
}

func sub1() {
	lock1.Lock()
	i--
	lock1.Unlock()
} */

var ii int32 = 100

// 底层实现 cas compare and swap 比较并交换
func add2() {
	atomic.AddInt32(&ii, 1)
}

func sub2() {
	atomic.AddInt32(&ii, -1)
}

func main() {
	for i := 0; i < 100; i++ {
		// go add1()
		// go sub1()

		go add2()
		go sub2()
	}
	time.Sleep(time.Second)
	// fmt.Printf("i: %v\n", i)
	fmt.Printf("ii: %v\n", ii)
}
