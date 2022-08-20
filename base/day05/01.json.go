package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name   string
	Age    int
	Gender string
}

func testJsonMarshal() {
	p := Person{
		Name:   "tom",
		Age:    20,
		Gender: "male",
	}
	b, _ := json.Marshal(p)
	fmt.Printf("string(b): %v\n", string(b)) // string(b): {"Name":"tom","Age":20,"Gender":"male"}
}

func testJsonUnmarshal() {
	b := []byte(`{"Name":"tom","Age":20,"Gender":"male"}`)
	var p Person
	json.Unmarshal(b, &p)
	fmt.Printf("p: %v\n", p) // p: {tom 20 male}
}

// 解析嵌套类型
func testJsonUnmarshal2() {
	b := []byte(`{"Name":"tom","Age":20,"Gender":"male", "Parents":["a", "b"]}`)
	var f map[string]interface{}
	json.Unmarshal(b, &f)
	fmt.Printf("f: %v\n", f) // f: map[Age:20 Gender:male Name:tom Parents:[a b]]
	for k, v := range f {
		fmt.Printf("k: %v\n", k)
		fmt.Printf("v: %v\n", v)
	}
}

func testJsonDecode() {
	f, _ := os.Open("a.json")
	defer f.Close()
	d := json.NewDecoder(f)
	var v map[string]interface{}
	d.Decode(&v)
	fmt.Printf("v: %v\n", v) // v: map[Age:20 Gender:male Name:tom Parents:[a b]]
}

func testJsonEncode() {
	type Person struct {
		Name   string
		Age    int
		gender string
		Parent []string
	}
	p := Person{
		Name:   "bob",
		Age:    18,
		gender: "male",
		Parent: []string{"a", "b"},
	}

	f, _ := os.OpenFile("b.json", os.O_CREATE|os.O_WRONLY, 0777)
	defer f.Close()
	e := json.NewEncoder(f)
	e.Encode(p)
}

func main() {
	// testJson()
	// testJsonUnmarshal()
	// testJsonUnmarshal2()
	// testJsonDecode()
	testJsonEncode()
}
