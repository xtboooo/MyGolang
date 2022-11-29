package main

import (
	"fmt"
	"sort"
)

type student struct {
	id   int `json:"id"`
	name string
	age  int
}

func f1() {
	demo := func(ce []student) {
		//切片是引用传递，是可以改变值的
		ce[1].age = 999
		// ce = append(ce, student{3, "xiaowang", 56})
		// return ce
	}
	var ce []student //定义一个切片类型的结构体
	ce = []student{
		student{1, "xiaoming", 22},
		student{2, "xiaozhang", 33},
	}
	fmt.Println(ce)
	demo(ce)
	fmt.Println(ce)
}

// 删除map类型的结构体
func f2() {
	ce := make(map[int]student)
	ce[1] = student{1, "xiaolizi", 22}
	ce[2] = student{2, "wang", 23}
	fmt.Println(ce)
	delete(ce, 2)
	fmt.Println(ce)
}

// 实现map有序输出
func f3() {
	map1 := make(map[int]string, 5)
	map1[1] = "www.topgoer.com"
	map1[2] = "rpc.topgoer.com"
	map1[5] = "ceshi"
	map1[3] = "xiaohong"
	map1[4] = "xiaohuang"
	sli := []int{}
	for k, _ := range map1 {
		sli = append(sli, k)
	}
	sort.Ints(sli)
	for i := 0; i < len(map1); i++ {
		fmt.Println(map1[sli[i]])
	}
}

func main() {
	// f1()
	f2()
}
