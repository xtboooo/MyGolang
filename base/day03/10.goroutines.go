package main

import (
	"fmt"
	"time"
)

func showMsg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg: %v\n", msg)
		time.Sleep(time.Microsecond * 100)
	}
}

func main() {
	go showMsg("python") // 启动了一个协程来执行了 1
	go showMsg("golang") // 2
	time.Sleep(time.Microsecond  * 2000)
	fmt.Println("main end...") // 3 主函数退出 程序就结束了
}
