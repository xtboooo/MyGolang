package main

import "fmt"

func f1() {
	m1 := make(map[string]string)
	fmt.Printf("m1: %v\n", m1)
	fmt.Printf("m1: %T\n", m1)
}

// 初始化
func f2() {
	var m1 = map[string]string{"name": "tom", "age": "20", "gender": "male"}
	fmt.Printf("m1: %v\n", m1)

	m2 := make(map[string]string)
	m2["name"] = "bob"
	m2["age"] = "18"
	m2["gender"] = "male"
	fmt.Printf("m2: %v\n", m2)
}

// 取值, key不存在不会报错
func f3() {
	var m1 = map[string]string{"name": "tom", "age": "20", "gender": "male"}
	fmt.Printf("m1[\"name\"]: %v\n", m1["name"])
	fmt.Printf("m1[\"height\"]: %v\n", m1["height"])
}

// 取值返回两个值
func f4() {
	var m1 = map[string]string{"name": "tom", "age": "20", "gender": "male"}
	k, ok := m1["name"]
	fmt.Printf("k: %v\n", k)
	fmt.Printf("ok: %v\n", ok)
	k1, ok1 := m1["height"]
	fmt.Printf("k1: %v\n", k1)
	fmt.Printf("ok1: %v\n", ok1)
}

func main() {
	// f1()
	// f2()
	// f3()
	f4()
}
