package main

import (
	"fmt"
)

//switch
func f1() {
	grade := "A"
	switch grade {
	case "A":
		fmt.Println("优秀")
	case "B":
		fmt.Println("良好")
	default:
		fmt.Println("其他")
	}
}
// 多值
func f2() {
	day := 2
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
	case 6, 7:
		fmt.Println("休息日")
	default:
		fmt.Println("非法输入")
	}
}

// 表达式
func f3() {
	score := 80
	switch {
	case score >= 60:
		fmt.Println("及格")
	case score >= 80 && score < 90:
		fmt.Println("优秀")
	default:
		fmt.Println("其他")
	}
}

// fallthrough
func f4(){
	num:=100
	switch num{
	case 100:
		fmt.Println("100")
		fallthrough
	case 200:
		fmt.Println("200")
	case 300:
		fmt.Println("300")
	}
}
func main() {
	f1()
	f2()
	f3()
	f4()
}
