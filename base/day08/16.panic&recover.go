package main

import (
	"errors"
	"fmt"
)

func f1() {
	test := func() {
		defer func() {
			if err := recover(); err != nil {
				println(err.(string)) // 将 interface{} 转型为具体类型。
			}
		}()
		panic("panic error!")
	}
	test()
}

// 向已关闭的通道发送数据会引发panic
func f2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	ch := make(chan int, 10)
	close(ch)
	ch <- 1
}

// 延迟调用中引发的错误，可被后续延迟调用捕获，但仅最后一个错误可被捕获
func f3() {
	func() {
		defer func() {
			fmt.Println(recover())
		}()

		defer func() {
			panic("defer panic")
		}()

		panic("test panic")
	}()
}

// 捕获函数 recover 只有在延迟调用内直接调用才会终止错误，否则总是返回 nil
// 任何未捕获的错误都会沿调用堆栈向外传递
func f4() {
	func() {
		defer func() {
			fmt.Println(recover()) //有效
		}()
		defer recover()              //无效！
		defer fmt.Println(recover()) //无效！
		defer func() {
			func() {
				println("defer inner")
				recover() //无效！
			}()
		}()
		panic("test panic")
	}()
}

// 使用延迟匿名函数或下面这样都是有效的s
func f5() {
	except := func() {
		fmt.Println(recover())
	}
	func() {
		defer except()
		panic("test panic")
	}()
}

// 如果需要保护代码段，可将代码块重构成匿名函数，如此可确保后续代码被执
func f6() {
	func(x, y int) {
		var z int
		func() {
			defer func() {
				if recover() != nil {
					z = 0
				}
			}()
			panic("test panic")
			z = x / y
			return
		}()

		fmt.Printf("x / y = %d\n", z)
	}(2, 1)
}

// 标准库 errors.New 和 fmt.Errorf 函数用于创建实现 error 接口的错误对象。通过判断错误对象实例来确定具体错误类型
func f7() {
	var ErrDivByZero = errors.New("division by zero")

	div := func(x, y int) (int, error) {
		if y == 0 {
			return 0, ErrDivByZero
		}
		return x / y, nil
	}

	defer func() {
		fmt.Println(recover())
	}()
	switch z, err := div(10, 0); err {
	case nil:
		println(z)
	case ErrDivByZero:
		panic(err)
	}
}

// Go实现类似 try catch 的异常处理
func f8() {
	Try := func(fun func(), handler func(interface{})) {
		defer func() {
			if err := recover(); err != nil {
				handler(err)
			}
		}()
		fun()
	}
	Try(func() {
		panic("test panic")
	}, func(err interface{}) {
		fmt.Println(err)
	})
}

func main() {
	// f1()
	// f2()
	// f3()
	// f4()
	// f5()
	// f6()
	// f7()
	f8()
}
