package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// 读所有
func testReadAll() {
	// r := strings.NewReader("hello golang")

	f, _ := os.Open("test.txt") // File 已经实现了Reader
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("string(b): %v\n", string(b))
}

// 读目录
func testReadDir() {
	fi, _ := ioutil.ReadDir(".")
	for _, v := range fi {
		fmt.Printf("v.Name(): %v\n", v.Name())
	}
}

// 读文件
func testReadFile() {
	b, _ := ioutil.ReadFile("test.txt")
	fmt.Printf("string(b): %v\n", string(b))
}

// 写文件
func testWriteFile() {
	ioutil.WriteFile("test.txt", []byte("hello world"), 0664)
}

// 创建临时文件
func testTempFile() {
	content := []byte("temporary file's content")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tmpfile.Name(): %v\n", tmpfile.Name())
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// testReadAll()
	// testReadDir()
	// testReadFile()
	// testWriteFile()
	testTempFile()
}
