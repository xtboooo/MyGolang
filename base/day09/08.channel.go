package main

import "fmt"

// 无缓冲的通道只有在有人接收值的时候才能发送值
func f1() {
	ch := make(chan int)
	ch <- 10
	fmt.Println("发送成功")
}

func f2() {
	recv := func(c chan int) {
		ret := <-c
		fmt.Println("接收成功", ret)
	}
	ch := make(chan int)
	go recv(ch) // 启用goroutine从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}

func main() {
	// f1()
	f2()
}
