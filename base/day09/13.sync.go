package main

import (
	"fmt"
	"strconv"
	"sync"
)

// sync.Once其实内部包含一个互斥锁和一个布尔值，互斥锁保证布尔值和数据的安全，而布尔值用来记录初始化是否完成。这样设计就能保证初始化操作的时候是并发安全的并且初始化操作也不会被执行多次。
// func f1() {
// 	var icons map[string]image.Image
// 	var loadIconsOnce sync.Once

// 	loadIcons := func() {
// 		icons = map[string]image.Image{
// 			"left":  loadIcon("left.png"),
// 			"up":    loadIcon("up.png"),
// 			"right": loadIcon("right.png"),
// 			"down":  loadIcon("down.png"),
// 		}
// 	}
// 	// Icon 是并发安全的
// 	Icon := func(name string) image.Image {
// 		loadIconsOnce.Do(loadIcons)
// 		return icons[name]
// 	}
// }

// Go语言中内置的map不是并发安全的 fatal error: concurrent map writes
func f2() {
	var m = make(map[string]int)
	get := func(key string) int {
		return m[key]
	}
	set := func(key string, value int) {
		m[key] = value
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func f3() {
	var m = sync.Map{}
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func main() {
	// f1()
	// f2()
	f3()
}
