package main

import (
	"fmt"
	// "io"
	"os"
)

func openClose() {
	/* 	f, err := os.Open("test.txt")
	   	if err != nil {
	   		fmt.Printf("err: %v\n", err)
	   	} else {
	   		fmt.Printf("f.Name(): %v\n", f.Name())
	   		f.Close()
	   	} */

	f, err := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE, 755) // 755 7所有权限 5 只读权限
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
		f.Close()
	}
}

// 创建文件
func createFile() {
	// 等价于 os.OpenFile("b.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	f, _ := os.Create("b.txt")
	fmt.Printf("f.Name(): %v\n", f.Name())
	f2, _ := os.CreateTemp("", "tmp")
	fmt.Printf("f2.Name(): %v\n", f2.Name())
}

// read
func readOps() {
	/* 	// 只打开一次文件
	   	f, _ := os.Open("test.txt")
	   	for {
	   		buf := make([]byte, 3)
	   		n, err := f.Read(buf)
	   		if err == io.EOF {
	   			break
	   		}
	   		fmt.Printf("n: %v\n", n)
	   		fmt.Printf("string(buf): %v\n", string(buf))
	   	} */

	/* 	// 从某个偏移量开始读取
	   	f, _ := os.Open("test.txt")
	   	buf := make([]byte, 4)
	   	n, _ := f.ReadAt(buf, 3)
	   	fmt.Printf("n: %v\n", n)
	   	fmt.Printf("string(buf): %v\n", string(buf)) */

	/* 	// 遍历目录
	   	de, _ := os.ReadDir("a")
	   	for _, v := range de {
	   		fmt.Printf("v.IsDir(): %v\n", v.IsDir())
	   		fmt.Printf("v.Name(): %v\n", v.Name())
	   	} */

	f, _ := os.Open("test.txt")
	f.Seek(3, 0)
	buf := make([]byte, 10)
	n, _ := f.Read(buf)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("string(buf): %v\n", string(buf))
	f.Close()

}

func main() {
	// openClose()
	// createFile()
	readOps()
}
