package main

import "os"

func write() {
	// os.Open() 打开的是只读的权限
	f, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_APPEND|os.O_TRUNC, 0755)
	f.Write([]byte("hello python"))
	f.Close()
}

// 写字符串
func writeString() {
	f, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_TRUNC|os.O_APPEND, 0755)
	f.WriteString("hello golang")
	f.Close()
}

// 在某处写
func writeAt() {
	f, _ := os.OpenFile("test.txt", os.O_RDWR, 0755)
	f.WriteAt([]byte("aaa"), 3)
	f.Close()
}

func main() {
	// write()
	// writeString()
	writeAt()
}
