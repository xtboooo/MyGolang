package quick_sort

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

var arrSize = 10000000
var printResult = false

func TestQuickSort(t *testing.T) {
	var arr = rand.Perm(arrSize)
	quickSort(&arr, 0, len(arr)-1)
	if printResult {
		fmt.Printf("arr: %v\n", arr)
	}
}

func TestConcurrentQuickSortV1(t *testing.T) {
	var arr = rand.Perm(arrSize)
	chanWait := make(chan int)
	go concurrentQuickSortV1(&arr, 0, len(arr)-1, chanWait)
	<-chanWait
	if printResult {
		fmt.Printf("arr: %v\n", arr)
	}
}

func TestConcurrentQuickSortV2(t *testing.T) {
	var arr = rand.Perm(arrSize)
	chanWait := make(chan int)
	go concurrentQuickSortV2(&arr, 0, len(arr)-1, chanWait)
	<-chanWait
	if printResult {
		fmt.Printf("arr: %v\n", arr)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	var arr = rand.Perm(arrSize)
	quickSort(&arr, 0, len(arr)-1)
}

func BenchmarkConcurrentQuickSortV1(b *testing.B) {
	var arr = rand.Perm(arrSize)
	chanWait := make(chan int)
	go concurrentQuickSortV1(&arr, 0, len(arr)-1, chanWait)
	<-chanWait
	if printResult {
		fmt.Printf("arr: %v\n", arr)
	}
}

func BenchmarkConcurrentQuickSortV2(b *testing.B) {
	var arr = rand.Perm(arrSize)
	chanWait := make(chan int)
	go concurrentQuickSortV2(&arr, 0, len(arr)-1, chanWait)
	<-chanWait
	if printResult {
		fmt.Printf("arr: %v\n", arr)
	}
}

func TestMain(m *testing.M) {
	fmt.Printf("arrSize: %v\n", arrSize)
	// printResult = true
	retCode := m.Run()
	os.Exit(retCode)
}
