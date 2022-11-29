package main

import (
	"fmt"
	"unsafe"
)

// 在 Go 中，与 C 数组变量隐式作为指针使用不同，Go 数组是值类型，赋值和函数传参操作都会复制整个数组数据。
func f1() {
	testArray := func(x [2]int) {
		fmt.Printf("func Array : %p , %v\n", &x, x)
	}
	var arrayB [2]int
	arrayA := [2]int{100, 200}
	arrayB = arrayA

	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)
	fmt.Printf("arrayB : %p , %v\n", &arrayB, arrayB)

	testArray(arrayA)

}

func f2() {
	testArrayPoint := func(x *[]int) {
		fmt.Printf("func Array : %p , %v\n", x, *x)
		(*x)[1] += 100
	}
	arrayA := [2]int{100, 200}
	// testArrayPoint(&arrayA) // 1.传数组指针
	arrayB := arrayA[:]
	testArrayPoint(&arrayB) // 2.传切片
	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)

}

// 从 slice 中得到一块内存地址
func f3() {
	s := make([]byte, 200)
	ptr := unsafe.Pointer(&s[0])
	fmt.Printf("ptr: %v\n", ptr)
}

// 从 Go 的内存地址中构造一个 slice
func f4() {
	// var ptr unsafe.Pointer
	// var s1 = struct {
	// 	addr uintptr
	// 	len  int
	// 	cap  int
	// }{ptr, length, length}
	// s := *(*[]byte)(unsafe.Pointer(&s1))
	// fmt.Printf("s: %v\n", s)

	// var o []byte
	// sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&o)))
	// sliceHeader.Cap = length
	// sliceHeader.Len = length
	// sliceHeader.Data = uintptr(ptr)
}

// // make 和切片字面量
// func makeslice(et *_type, len, cap int) slice {
// 	// 根据切片的数据类型，获取切片的最大容量
// 	maxElements := maxSliceCap(et.size)
// 	// 比较切片的长度，长度值域应该在[0,maxElements]之间
// 	if len < 0 || uintptr(len) > maxElements {
// 		panic(errorString("makeslice: len out of range"))
// 	}
// 	// 比较切片的容量，容量值域应该在[len,maxElements]之间
// 	if cap < len || uintptr(cap) > maxElements {
// 		panic(errorString("makeslice: cap out of range"))
// 	}
// 	// 根据切片的容量申请内存
// 	p := mallocgc(et.size*uintptr(cap), et, true)
// 	// 返回申请好内存的切片的首地址
// 	return slice{p, len, cap}
// }

// // int64
// func makeslice64(et *_type, len64, cap64 int64) slice {
//     len := int(len64)
//     if int64(len) != len64 {
//         panic(errorString("makeslice: len out of range"))
//     }

//     cap := int(cap64)
//     if int64(cap) != cap64 {
//         panic(errorString("makeslice: cap out of range"))
//     }

//     return makeslice(et, len, cap)
// }

// 切片的扩容策略
func f5() {
	slice := []int{10, 20, 30, 40}
	newSlice := append(slice, 50)
	fmt.Printf("Before slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("Before newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	newSlice[1] += 10
	fmt.Printf("After slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("After newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
}

// copy
func f6() {
	array := []int{10, 20, 30, 40}
	slice := make([]int, 6)
	n := copy(slice, array)
	fmt.Println(n, slice)
}
func main() {
	// f1()
	// f2()
	// f3()
	// f5()
	f6()
}
