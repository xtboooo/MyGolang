package main

import (
	"fmt"
	"sort"
)

type NewInts []uint

func (n NewInts) Len() int {
	return len(n)
}

// 自己实现排序
func (n NewInts) Less(i, j int) bool {
	fmt.Println(i, j, n[i] < n[j], n)
	return n[i] < n[j]
}

func (n NewInts) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func testFloat64() {
	f := []float64{1.1, 4.4, 5.5, 3.3, 2.2}
	sort.Float64s(f)
	fmt.Printf("f: %v\n", f)
}

func testString() {
	s := []string{"100", "44", "43", "2", "1"}
	fmt.Printf("s: %v\n", s) // [100 44 43 2 1]
	sort.Strings(s)
	fmt.Printf("s: %v\n", s) // [1 100 2 43 44]

	ls := []string{"d", "ac", "c", "ab", "e"}
	fmt.Printf("ls: %v\n", ls) // [d ac c ab e]
	sort.Strings(ls)
	fmt.Printf("ls: %v\n", ls) // [ab ac c d e]

	ls = []string{"啊", "波", "次", "得", "额", "周"}
	fmt.Printf("ls: %v\n", ls) // [啊 波 次 得 额 周]
	sort.Strings(ls)
	fmt.Printf("ls: %v\n", ls) // [周 啊 得 次 波 额]

	for _, v := range ls {
		fmt.Println(v, []byte(v))
	}
	/*
		周 [229 145 168]
		啊 [229 149 138]
		得 [229 190 151]
		次 [230 172 161]
		波 [230 179 162]
		额 [233 162 157]
	*/
}

// 复杂结构
type testSlice [][]int

func (l testSlice) Len() int {
	return len(l)
}
func (l testSlice) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
func (l testSlice) Less(i, j int) bool {
	return l[i][1] < l[j][1]
}

// 复杂结构体
type testSlice2 []map[string]float64

func (l testSlice2) Len() int {
	return len(l)
}

func (l testSlice2) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l testSlice2) Less(i, j int) bool {
	return l[i]["a"] < l[j]["a"]
}

// 复杂结构体
type People struct {
	Name string
	Age  int
}

type testSlice3 []People

func (l testSlice3) Len() int {
	return len(l)
}

func (l testSlice3) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l testSlice3) Less(i, j int) bool {
	return l[i].Age < l[j].Age
}

func main() {
	/* 	s := []int{2, 3, 1, 4}
	   	sort.Ints(s)
	   	fmt.Printf("s: %v\n", s) */

	/* 	n := []uint{1, 3, 2}
	   	sort.Sort(NewInts(n))
	   	fmt.Println(n) */

	// testFloat64()

	// testString()

	/* 	ls:=testSlice{
	   		{1,4},
	   		{9,3},
	   		{7,5},
	   	}
	   	fmt.Printf("ls: %v\n", ls) // ls: [[1 4] [9 3] [7 5]]
	   	sort.Sort(ls)
	   	fmt.Printf("ls: %v\n", ls) // ls: [[9 3] [1 4] [7 5]] */

	/* ls := testSlice2{
		{"a": 4, "b": 12},
		{"a": 3, "b": 13},
		{"a": 2, "b": 14},
	}
	fmt.Printf("ls: %v\n", ls) // ls: [map[a:4 b:12] map[a:3 b:13] map[a:2 b:14]]
	sort.Sort(ls)
	fmt.Printf("ls: %v\n", ls) // ls: [map[a:2 b:14] map[a:3 b:13] map[a:4 b:12]] */

	ls := testSlice3{
		{Name: "n1", Age: 12},
		{Name: "n2", Age: 11},
		{Name: "n3", Age: 10},
	}

	fmt.Printf("ls: %v\n", ls) // ls: [{n1 12} {n2 11} {n3 10}]
	sort.Sort(ls)
	fmt.Printf("ls: %v\n", ls) // ls: [{n3 10} {n2 11} {n1 12}]
}
