package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

// defer 是先进后出
func f1() {
	var whatever [5]struct{}

	for i := range whatever {
		defer fmt.Println(i)
	}
}

// defer closure
// 由于闭包用到的变量 i 在执行的时候已经变成4,所以输出全都是4
func f2() {
	var whatever [5]struct{}
	for i := range whatever {
		defer func() { fmt.Println(i) }()
	}
}

type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name, " closed")
}

func f3() {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		t2 := t
		defer t2.Close()
	}
}

// defer后面的语句在执行的时候，函数调用的参数会被保存起来，但是不执行
func f4() {
	defer_test := func(x int) {
		defer println("a")
		defer println("b")
		defer func() {
			println(100 / x) // div0 异常未被捕获，逐步往外传递，最终终止进程。
		}()
		defer println("c")
	}
	defer_test(0)
}

// 延迟调用参数在注册时求值或复制，可用指针或闭包 "延迟" 读取
func f5() {
	x, y := 10, 20
	defer func(i int) {
		println("defer:", i, y) // y 闭包引用
	}(x) // x 被复制
	x += 10
	y += 100
	println("x =", x, "y =", y)
}

// 滥用 defer 可能会导致性能问题，尤其是在一个 "大循环" 里
func f6() {
	var lock sync.Mutex
	test := func() {
		lock.Lock()
		lock.Unlock()
	}

	test_defer := func() {
		lock.Lock()
		defer lock.Unlock()
	}

	func() {
		t1 := time.Now()

		for i := 0; i < 10000; i++ {
			test()
		}
		elapsed := time.Since(t1)
		fmt.Println("test elapsed: ", elapsed)
	}()
	func() {
		t1 := time.Now()

		for i := 0; i < 10000; i++ {
			test_defer()
		}
		elapsed := time.Since(t1)
		fmt.Println("testdefer elapsed: ", elapsed)
	}()
}

// defer & closure
// 如果 defer 后面跟的不是一个 closure 最后执行的时候我们得到的并不是最新的值
func f7() {
	foo := func(a, b int) (i int, err error) {
		// 不是函数
		defer fmt.Printf("first defer err %v\n", err)
		// 引用的是自己作用域的变量
		defer func(err error) { fmt.Printf("second defer err %v\n", err) }(err)
		// closure
		defer func() { fmt.Printf("third defer err %v\n", err) }()
		if b == 0 {
			err = errors.New("divided by zero!")
			return
		}
		i = a / b
		return
	}
	foo(2, 0)
}

// 在有具名返回值的函数中（这里具名返回值为 i），执行 return 2 的时候实际上已经将 i 的值重新赋值为 2。所以defer closure 输出结果为 2 而不是 1
func f8() {
	foo := func() (i int) {
		i = 0
		defer func() {
			fmt.Println(i)
		}()
		return 2
	}
	foo()
}

// 名为 test 的函数一直运行至结束，然后 defer 函数会被执行且会因为值为 nil 而产生 panic 异常。
// 然而值得注意的是，run() 的声明是没有问题，因为在test函数运行完成后它才会被调用
func f9() {
	test := func() {
		var run func() = nil
		defer run()
		fmt.Println("runs")
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	test()
}

// 总是在一次成功的资源分配下面使用 defer ，对于这种情况来说意味着：当且仅当 http.Get 成功执行时才使用 defer
func f10() {
	do := func() error {
		res, err := http.Get("http://xxxxxxxxxx")
		if res != nil {
			defer res.Body.Close()
		}
		if err != nil {
			return err
		}
		return nil
	}
	do()
}

func f11() {
	func() (err error) {
		f, err := os.Open("book.txt")
		if err != nil {
			return err
		}
		if f != nil {
			defer func() {
				if ferr := f.Close(); ferr != nil {
					err = ferr
				}
			}()
		}

		return nil
	}()
}

// 当延迟函数执行时，只有最后一个变量会被用到，因此，f 变量 会成为最后那个资源 (another-book.txt)
// 而且两个 defer 都会将这个资源作为最后的资源来关闭
func f12() {
	func() error {
		f, err := os.Open("book.txt")
		if err != nil {
			return err
		}
		if f != nil {
			defer func(f io.Closer) {
				if err := f.Close(); err != nil {
					fmt.Printf("defer close book.txt err %v\n", err)
				}
			}(f)
		}

		// ..code...

		f, err = os.Open("another-book.txt")
		if err != nil {
			return err
		}
		if f != nil {
			defer func(f io.Closer) {
				if err := f.Close(); err != nil {
					fmt.Printf("defer close another-book.txt err %v\n", err)
				}
			}(f)
		}
		return nil
	}()
}


func main() {
	// f1()
	// f2()
	// f3()
	// f4()
	// f5()
	// f6()
	// f7()
	// f8()
	// f9()
	// f10()
	// f11()
	f12()
}
