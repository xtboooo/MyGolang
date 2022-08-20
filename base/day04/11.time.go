package main

import (
	"fmt"
	"time"
)

func testTime1() {
	t := time.Now()
	fmt.Printf("t: %T\n", t)
	fmt.Printf("t: %v\n", t)

	year := t.Year()
	month := t.Month()
	day := t.Day()
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	fmt.Printf("%T, %T, %T, %T, %T, %T\n", year, month, day, hour, minute, second)
}

func testTime2() {
	now := time.Now()
	// 当前时间戳 Timestamp type:int64, Timestamp:1660916028
	fmt.Printf("Timestamp type:%T, Timestamp:%v", now.Unix(), now.Unix())
}

func add(h, m, s, mls, ns time.Duration) {
	now := time.Now()
	fmt.Println(now.Add(time.Hour*h + time.Minute*m + time.Second*s + time.Millisecond*mls + time.Nanosecond*ns))
}

func testTimeFormat() {
	t := time.Now()
	fmt.Println(t.Format("2006/01/02 15:04"))
}

func main() {
	// testTime1()
	// testTime2()
	// add(1, 2, 3, 4, 5)
	testTimeFormat()
}
