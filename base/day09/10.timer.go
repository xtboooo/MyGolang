package main

import (
	"fmt"
	"time"
)

// 1.timer基本使用
func f1() {
	timer1 := time.NewTimer(2 * time.Second)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)
	t2 := <-timer1.C
	fmt.Printf("t2:%v\n", t2)
}

// 2.验证timer只能响应1次
func f2() {
	timer2 := time.NewTimer(time.Second)
	for {
		<-timer2.C
		fmt.Println("时间到")
	}
}

// 3.timer实现延时的功能
func f3() {
	time.Sleep(time.Second)
	timer3 := time.NewTimer(2 * time.Second)
	<-timer3.C
	fmt.Println("2秒到")
	<-time.After(2 * time.Second)
	fmt.Println("2秒到")
}

// 4.停止定时器
func f4() {
	timer4 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer4.C
		fmt.Println("定时器执行了")
	}()
	b := timer4.Stop()
	if b {
		fmt.Println("timer4已经关闭")
	}
}

// 5.重置定时器
func f5() {
	timer5 := time.NewTimer(3 * time.Second)
	timer5.Reset(1 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-timer5.C)
	for {
	}
}

// Ticker
func f6() {
	// 1.获取ticker对象
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	// 子协程
	go func() {
		for {
			//<-ticker.C
			i++
			fmt.Println(<-ticker.C)
			if i == 5 {
				//停止
				ticker.Stop()
			}
		}
	}()
	for {
	}
}

func main() {
	// f1()
	// f2()
	// f3()
	// f4()
	// f5()
	f6()
}
