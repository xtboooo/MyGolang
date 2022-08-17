package main

import (
	"fmt"
	"sync/atomic"
)

func test_add_sub() {
	var j int32 = 100
	atomic.AddInt32(&j, 1)
	fmt.Printf("j: %v\n", j)
	atomic.AddInt32(&j, -1)
	fmt.Printf("j: %v\n", j)

	var h int64 = 100
	atomic.AddInt64(&h, 1)
	fmt.Printf("h: %v\n", h)
}

func test_load_store() {
	var i int32 = 100
	atomic.LoadInt32(&i) // read
	fmt.Printf("i: %v\n", i)

	atomic.StoreInt32(&i, 200) // write
	fmt.Printf("i: %v\n", i)
}

func test_cas() {
	// cas
	var i int32 = 100
	b := atomic.CompareAndSwapInt32(&i, 100, 200)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("i: %v\n", i)
}

// 线程并发问题 原子化操作
func main() {
	// test_add_sub()
	// test_load_store()
	// test_cas
}
