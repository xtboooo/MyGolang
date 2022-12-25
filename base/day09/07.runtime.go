package main

import (
	"fmt"
	"runtime"
	"time"
)

// runtime.Gosched()
func f1() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Println("hello")
	}
}

// runtime.Goexit()
func f2() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for {
	}
}

// runtime.GOMAXPROCS
func f3() {
	a := func() {
		for i := 1; i < 10; i++ {
			fmt.Println("A:", i)
		}
	}
	b := func() {
		for i := 1; i < 10; i++ {
			fmt.Println("B:", i)
		}
	}
	runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(time.Second)
}

// runtime.GOMAXPROCS
func f4() {
	a := func() {
		for i := 1; i < 10; i++ {
			fmt.Println("A:", i)
		}
	}
	b := func() {
		for i := 1; i < 10; i++ {
			fmt.Println("B:", i)
		}
	}
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}

func f5() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("main结束")
}

// 循环取值
func f6() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
}

// 单向通道
func f7() {
	counter := func(out chan<- int) {
		for i := 0; i < 100; i++ {
			out <- i
		}
		close(out)
	}

	squarer := func(out chan<- int, in <-chan int) {
		for i := range in {
			out <- i * i
		}
		close(out)
	}
	printer := func(in <-chan int) {
		for i := range in {
			fmt.Println(i)
		}
	}
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}

func main() {
	// f1()
	// f2()
	// f3()
	// f4()
	// f5()
	// f6()
	f7()
}
