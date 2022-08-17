package main

import (
	"fmt"
	"os"
)

// 创建文件
func createFile() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
}

// 创建目录
func makeDir() {
	/* 	// 单目录
	   	err := os.Mkdir("test", os.ModePerm)
	   	if err != nil {
	   		fmt.Printf("err: %v\n", err)
	   	} */

	// 嵌套目录
	err2 := os.MkdirAll("a/b/c", os.ModePerm)
	if err2 != nil {
		fmt.Printf("err: %v\n", err2)
	}

}

func remove() {
	/* 	// 删除单个目录
	   	err:= os.Remove("a.txt")
	   	if err != nil{
	   		fmt.Printf("err: %v\n", err)
	   	} */
	// 删除目录下的所有文件
	err2 := os.RemoveAll("a")
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
}

// 获得当前目录
func getWd() {
	dir, _ := os.Getwd()
	fmt.Printf("dir: %v\n", dir)

	// 更改目录
	os.Chdir("d:/")
	dir, _ = os.Getwd()
	fmt.Printf("dir: %v\n", dir)

	tmp := os.TempDir()
	fmt.Printf("t: %v\n", tmp)
}

// 重命名
func rename() {
	err := os.Rename("test.txt", "a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 读文件
func readFile() {
	b, _ := os.ReadFile("test.txt")
	fmt.Printf("b: %v\n", string(b[:]))
}

// 写文件
func writeFile() {
	os.WriteFile("test.txt", []byte("hello golang"), os.ModePerm)

}

func main() {
	// createFile()
	// makeDir()
	// remove()
	// getWd()
	// rename()
	// readFile()
	writeFile()
}
