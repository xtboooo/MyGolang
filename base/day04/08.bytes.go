package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func testTrans() {
	var i int = 100
	var b byte = 10
	b = byte(i)
	i = int(b)

	var s string = "hello world"
	b1 := []byte{1, 2, 3}
	s = string(b1)
	b1 = []byte(s)
}

func testContains() {
	s := "hello golang"
	b := []byte(s)
	b1 := []byte("hello")
	b2 := []byte("HELLO")
	b3 := bytes.Contains(b, b1)
	fmt.Printf("b3: %v\n", b3)
	b4 := bytes.Contains(b, b2)
	fmt.Printf("b4: %v\n", b4)

	b5 := strings.Contains("hello world", "hello")
	fmt.Printf("b5: %v\n", b5)
}

func testCount() {
	// Count
	s := []byte("hellooooooooooo")
	sep1 := []byte("h")
	sep2 := []byte("l")
	sep3 := []byte("o")
	fmt.Printf("bytes.Count(s, sep1): %v\n", bytes.Count(s, sep1)) // 1
	fmt.Printf("bytes.Count(s, sep2): %v\n", bytes.Count(s, sep2)) // 2
	fmt.Printf("bytes.Count(s, sep3): %v\n", bytes.Count(s, sep3)) // 11
}
func testRepeat() {
	// Repeat
	b := []byte("hi")
	fmt.Printf("string(bytes.Repeat(b, 1)): %v\n", string(bytes.Repeat(b, 1))) // hi
	fmt.Printf("string(bytes.Repeat(b, 3)): %v\n", string(bytes.Repeat(b, 3))) // hihihi
}

func testReplace() {
	// Replace
	s := []byte("hello world")
	old := []byte("o")
	new := []byte("ee")
	// -1代表有几次替换几次
	fmt.Printf("string(bytes.Replace(s, old, new, 0)): %v\n", string(bytes.Replace(s, old, new, 0)))   // hello world
	fmt.Printf("string(bytes.Replace(s, old, new, 1)): %v\n", string(bytes.Replace(s, old, new, 1)))   // hellee world
	fmt.Printf("string(bytes.Replace(s, old, new, 2)): %v\n", string(bytes.Replace(s, old, new, 2)))   // hellee weerld
	fmt.Printf("string(bytes.Replace(s, old, new, -1)): %v\n", string(bytes.Replace(s, old, new, -1))) // hellee weerld
}

func testRunes() {
	// Runes
	s := []byte("你好世界!")
	r := bytes.Runes(s)
	fmt.Printf("len(s): %v\n", len(s)) // 13
	fmt.Printf("len(r): %v\n", len(r)) // 5
}

func testJoin() {
	// Join
	s := [][]byte{[]byte("你好"), []byte("世界")}
	sep1 := []byte(",")
	fmt.Printf("string(bytes.Join(s, sep1)): %v\n", string(bytes.Join(s, sep1))) // 你好,世界
	sep2 := []byte("#")
	fmt.Printf("string(bytes.Join(s, sep2)): %v\n", string(bytes.Join(s, sep2))) //
}

func testReader() {
	data := "123456789"
	// 通过[]byte创建Reader
	r := bytes.NewReader([]byte(data))
	// 返回未读取部分的长度
	fmt.Printf("r.Len(): %v\n", r.Len())
	// 返回底层数据总长度
	fmt.Printf("r.Size(): %v\n", r.Size())
	fmt.Println("---------------")
	buf := make([]byte, 2)
	for {
		// 读取数据
		n, err := r.Read(buf)
		if err != nil {
			break
		}
		fmt.Printf("string(buf[:%v]): %v\n", n, string(buf[:n]))
	}
	fmt.Println("----------------")
	// 设置偏移量, 因为上面的操作已经修改了读取位置等信息
	r.Seek(0, 0)
	for {
		b, err := r.ReadByte()
		if err != nil {
			break
		}
		fmt.Printf("string(b): %v\n", string(b))
	}
	fmt.Println("-----------------")

	r.Seek(0, 0)
	off := int64(0)

	for {
		// 指定偏移量读取
		n, err := r.ReadAt(buf, off)
		if err != nil {
			break
		}
		off += int64(n)
		fmt.Println(off, string(buf[:n]))
	}
}

func testBuffer() {
	var b bytes.Buffer
	fmt.Printf("b: %v\n", b)
	var b1 = bytes.NewBufferString("hello")
	fmt.Printf("b1: %v\n", b1)
	var b2 = bytes.NewBuffer([]byte("hello"))
	fmt.Printf("b2: %v\n", b2)
}

func testBuffer2() {
	var b bytes.Buffer
	n, _ := b.WriteString("hello")
	fmt.Printf("n: %v\n", n)
	fmt.Printf("b.Bytes(): %v\n", string(b.Bytes()))
}

func testBuffer3() {
	var b = bytes.NewBufferString("hello world")
	b1 := make([]byte, 2)
	for {
		n, err := b.Read(b1)
		if err == io.EOF {
			break
		}
		fmt.Printf("n: %v\n", n)
		fmt.Printf("string(b1): %v\n", string(b1[0:n]))
	}

}

func main() {
	// testTrans()
	// testContains()
	// testCount()
	// testRepeat()
	// testReplace()
	// testRunes()
	// testJoin()
	// testReader()
	// testBuffer()
	testBuffer3()
}
