package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 数组定义
func f1() {
	var a1 [2]int
	var a2 [3]string
	var a3 [2]bool
	var a4 [2]float32
	fmt.Printf("a1: %T\n", a1)
	fmt.Printf("a2: %T\n", a2)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)
	fmt.Printf("a4: %v\n", a4)
}

// 数组的初始化
func f2() {
	var a1 = [3]int{1, 2}
	fmt.Printf("a1: %v\n", a1)
	var a2 = [2]string{"hello", "world"}
	fmt.Printf("a2: %v\n", a2)
	var a3 = [2]bool{true, false}
	fmt.Printf("a3: %v\n", a3)
}

// 数组的初始化 省略长度
func f3() {
	var a1 = [...]int{1, 2, 3}
	fmt.Printf("len(a1): %v\n", len(a1))
	var a2 = [...]string{"good", "golang"}
	fmt.Printf("len(a2): %v\n", len(a2))
}

// 指定索引来初始化数组
func f4() {
	var a1 = [...]int{0: 1, 2: 3}
	fmt.Printf("a1: %v\n", a1)
	var a2 = [...]bool{3: true, 5: false}
	fmt.Printf("a2: %v\n", a2)
}

// 数组的修改
func f5() {
	var a1 = [...]int{1, 2, 3}
	fmt.Printf("a1: %v\n", a1)
	a1[1] = 100
	fmt.Printf("a1: %v\n", a1)
}

// 二维数组
func f6() {
	var arr0 [5][3]int
	var arr1 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}
	a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	b := [...][2]int{{1, 1}, {2, 2}, {3, 3}} // 第 2 纬度不能用 "..."。
	fmt.Printf("arr0: %v\n", arr0)
	fmt.Printf("arr1: %v\n", arr1)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
}

// 多维数组遍历
func f7() {
	var f [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

	for k1, v1 := range f {
		for k2, v2 := range v1 {
			fmt.Printf("(%d,%d)=%d ", k1, k2, v2)
		}
		fmt.Println()
	}
}

// 数组拷贝和传参
func f8() {
	printArr := func(arr *[5]int) {
		arr[0] = 10
		for i, v := range arr {
			fmt.Println(i, v)
		}
	}
	var arr1 [5]int
	printArr(&arr1)
	fmt.Println(arr1)

	arr2 := [...]int{2, 4, 6, 8, 10}
	printArr(&arr2)
	fmt.Println(arr2)
}

// 求数组所有元素之和
func f9() {
	sumArr := func(a [10]int) int {
		var sum int = 0
		for i := 0; i < len(a); i++ {
			sum += a[i]
		}
		return sum
	}

	printArr := func(arr *[10]int) {
		arr[0] = 10
		for i, v := range arr {
			fmt.Printf("%v %v\n", i, v)
		}
	}

	// 若想做一个真正的随机数，要种子
	// seed()种子默认是1
	//rand.Seed(1)
	rand.Seed(time.Now().Unix())

	var b [10]int
	for i := 0; i < len(b); i++ {
		// 产生一个0到1000随机数
		b[i] = rand.Intn(1000)
	}
	sum := sumArr(b)
	printArr(&b)

	fmt.Printf("sum=%d\n", sum)
}

// 找出数组中和为给定值的两个元素的下标，例如数组[1,3,5,8,7]，找出两个元素之和等于8的下标分别是（0，4）和（1，2）
func f10() {
	findIndexByTarget := func(a [5]int, target int) {
		// 遍历数组
		for i := 0; i < len(a); i++ {
			other := target - a[i]
			// 继续遍历
			for j := i + 1; j < len(a); j++ {
				if a[j] == other {
					fmt.Printf("(%d,%d)\n", i, j)
				}
			}
		}
	}
	b := [5]int{1, 3, 5, 8, 7}
	findIndexByTarget(b, 8)
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
	f10()
}
