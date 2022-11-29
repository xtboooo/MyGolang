package main

import "fmt"

func f1() {
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type:*int
	fmt.Println(&b)                    // 0xc00000e018
}

// 指针取值
func f2() {
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
}

// 指针传值
func f3() {
	modify1 := func(x int) {
		x = 100
	}

	modify2 := func(x *int) {
		*x = 100
	}
	a := 10
	modify1(a)
	fmt.Printf("a modify1: %v\n", a)
	modify2(&a)
	fmt.Printf("a modify2: %v\n", a)
}

// 空指针
func f4() {
	var p *string
	fmt.Println(p)
	fmt.Printf("p的值是%v\n", p)
	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空值")
	}
}

// new
func f5() {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Println(*a)       // 0
	fmt.Println(*b)       // false
}

// make
func f6() {
	// 声明
	var b map[string]int
	// 初始化
	b = make(map[string]int, 10)
	b["测试"] = 100
	fmt.Println(b)
}

// 将num的地址赋给指针ptr，并通过ptr去修改num的值
func f7() {
	var a int
	fmt.Println(&a)
	var p *int
	p = &a
	*p = 20

}

func main() {
	// f1()
	// f2()
	// f3()
	// f4()
	// f5()
	// f6()
	f7()
}
