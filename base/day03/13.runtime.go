package main

import (
	"fmt"
	"runtime"
	"time"
)

/* func showMsg2(msg string) {
	for i := 0; i < 2; i++ {
		fmt.Printf("msg: %v\n", msg)
	}
}

func main() {
	go showMsg2("python")
	for i := 0; i < 2; i++ {
		runtime.Gosched() // 让出当前cpu, 让其他协程先执行
		fmt.Printf("\"golang\": %v\n", "golang")
	}
	fmt.Println("end...")
} */

/* func show() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i >= 5 {
			runtime.Goexit()
		}
	}
}

func main() {
	go show()
	time.Sleep(time.Second)
} */

func a() {
	for i := 0; i < 10; i++ {
		fmt.Printf("a i: %v\n", i)
		// time.Sleep(time.Millisecond*100)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Printf("b i: %v\n", i)
	}
}

func main() {
	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
	runtime.GOMAXPROCS(1) // 设置运行协程的cpu数量
	go a()
	go b()
	time.Sleep(time.Second)
}
