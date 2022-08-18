package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func test1() {
	// r := strings.NewReader("hello world")
	f, _ := os.Open("test.txt")
	defer f.Close()
	r2 := bufio.NewReader(f)
	s, _ := r2.ReadString('\n')
	fmt.Printf("s: %v\n", s)
}

func test2() {
	r := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	r2 := bufio.NewReader(r)
	b := make([]byte, 10)
	for {
		n, err := r2.Read(b)
		if err == io.EOF {
			break
		} else {
			fmt.Printf("string(b[0:%v]): %v\n", n, string(b[0:n]))
		}
	}
}

func test3() {
	r := strings.NewReader("ABCDEFG")
	r2 := bufio.NewReader(r)
	b, _ := r2.ReadByte()
	fmt.Printf("b: %c\n", b)
	b, _ = r2.ReadByte()
	fmt.Printf("b: %c\n", b)
	// r2.UnreadByte()
	b, _ = r2.ReadByte()
	fmt.Printf("b: %c\n", b)
}

func test4() {
	r := strings.NewReader("ABC\nDEF\r\nGHI\r\nGHI")
	r2 := bufio.NewReader(r)
	line, isPrefix, _ := r2.ReadLine()
	fmt.Printf("%q %v\n", line, isPrefix)

	line, isPrefix, _ = r2.ReadLine()
	fmt.Printf("%q %v\n", line, isPrefix)

	line, isPrefix, _ = r2.ReadLine()
	fmt.Printf("%q %v\n", line, isPrefix)

	line, isPrefix, _ = r2.ReadLine()
	fmt.Printf("%q %v\n", line, isPrefix)
}

func test5() {
	r := strings.NewReader("ABC,DEF,GHI,JKL")
	r2 := bufio.NewReader(r)
	line, _ := r2.ReadSlice(',')
	fmt.Printf("line: %q\n", line)
	line, _ = r2.ReadSlice(',')
	fmt.Printf("line: %q\n", line)
	line, _ = r2.ReadSlice(',')
	fmt.Printf("line: %q\n", line)
}

func test6() {
	r := strings.NewReader("ABC DEF GHI JKL")
	r2 := bufio.NewReader(r)
	line, _ := r2.ReadBytes(' ')
	fmt.Printf("line: %q\n", line)
	line, _ = r2.ReadBytes(' ')
	fmt.Printf("line: %q\n", line)
	line, _ = r2.ReadBytes(' ')
	fmt.Printf("line: %q\n", line)
}

func test7() {
	r := strings.NewReader("ABC DEF GHI JKL")
	r2 := bufio.NewReader(r)
	line, _ := r2.ReadString(' ')
	fmt.Printf("line: %q\n", line)
	line, _ = r2.ReadString(' ')
	fmt.Printf("line: %q\n", line)
	line, _ = r2.ReadString(' ')
	fmt.Printf("line: %q\n", line)
}

func test8() {
	r := strings.NewReader("ABC DEF GHI JKL")
	r2 := bufio.NewReader(r)
	// b := bytes.NewBuffer(make([]byte, 0))
	// File 实现了 Reader Writer 接口
	f, _ := os.OpenFile("test.txt", os.O_RDWR, 0777)
	defer f.Close()
	r2.WriteTo(f)
	// fmt.Printf("b: %v\n", b)
}

func test9() {
	f, _ := os.OpenFile("test.txt", os.O_RDWR, 0777)
	defer f.Close()
	w := bufio.NewWriter(f)
	i, _ := w.WriteString("hello world")
	w.Flush()
	fmt.Printf("i: %v\n", i)
}

func test10() {
	b := bytes.NewBuffer(make([]byte, 0))
	w := bufio.NewWriter(b)
	w.WriteString("123456")
	b2 := bytes.NewBuffer(make([]byte, 0))
	w.Reset(b2)
	w.WriteString("789")
	w.Flush()
	fmt.Printf("b: %v\n", b)
	fmt.Printf("b2: %v\n", b2)
}

func test11() {
	b := bytes.NewBuffer(make([]byte, 0))
	w := bufio.NewWriter(b)
	fmt.Printf("w.Available(): %v\n", w.Available())
	fmt.Printf("w.Buffered(): %v\n", w.Buffered())

	w.WriteString("ABCDEFG")
	fmt.Printf("w.Available(): %v\n", w.Available())
	fmt.Printf("w.Buffered(): %v\n", w.Buffered())
	fmt.Printf("w: %q\n", b)

	w.Flush()
	fmt.Printf("w.Available(): %v\n", w.Available())
	fmt.Printf("w.Buffered(): %v\n", w.Buffered())
	fmt.Printf("w: %q\n", b)
}

func test12() {
	b := bytes.NewBuffer(make([]byte, 0))
	w := bufio.NewWriter(b)

	r := strings.NewReader("123")
	r2 := bufio.NewReader(r)

	rw := bufio.NewReadWriter(r2, w)
	s, _ := rw.ReadString('\n')
	fmt.Printf("string(s): %v\n", string(s))
	rw.WriteString("4567")
	rw.Flush()
	fmt.Printf("b: %q\n", b)
}

func test13() {
	r := strings.NewReader("ABC DEF GHI JKL")
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords) // 以空格作为分隔符作为分割
	for s.Scan() {
		fmt.Printf("s.Text(): %v\n", s.Text())
	}
}

func test14() {
	r := strings.NewReader("hello 世界")
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes)
	for s.Scan() {
		fmt.Printf("%s ", s.Text())
	}
}

func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	// test5()
	// test6()
	// test7()
	// test8()
	// test9()
	// test10()
	// test11()
	// test12()
	// test13()
	test14()
}
