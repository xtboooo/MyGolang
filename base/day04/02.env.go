package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取所有的环境变量
	// fmt.Printf("os.Environ(): %v\n", os.Environ())

	// 查找
	s := os.Getenv("GOPATH")
	fmt.Printf("s: %v\n", s)

	// 查找存不存在
	s2, b := os.LookupEnv("GOPATH")
	if b {
		fmt.Printf("s2: %v\n", s2)
	}

	// 设置
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "usr/gopher")
	fmt.Println(os.ExpandEnv("$NAME lives in ${BURROW}."))
	
	// 清除所有的环境变量
	// os.Clearenv()
}
