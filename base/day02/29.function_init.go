package main

import "fmt"

var i = initVar()

// 先执行init 后执行main
func init() {
	fmt.Println("init...")
}

func init()  {
	fmt.Println("init2...")
}

func initVar() int {
	fmt.Println("initVar...")
	return 100
}

func main() {
	fmt.Println("main...")
}