package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Person1 struct {
	// 反引号
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Gender  string   `xml:"gender"`
}

func xmlMarshal() {
	person := Person1{
		Name:   "bob",
		Age:    21,
		Gender: "male",
	}
	// b, _ := xml.Marshal(person)

	b, _ := xml.MarshalIndent(person, " ", "  ")

	fmt.Printf("%v\n", string(b)) // string(b): <person><name>bob</name><age>21</age><gender>male</gender></person>
}

func xmlUnmarshal() {
	s := `
	<person>
	<name>bob</name>
	<age>21</age>
	<gender>male</gender>
  </person>
	`
	b := []byte(s)
	var per Person1
	xml.Unmarshal(b, &per)
	fmt.Printf("per: %v\n", per) // per: {{ person} bob 21 male}
}

func read() {
	b, _ := ioutil.ReadFile("a.xml")
	var per Person1
	xml.Unmarshal(b, &per)
	fmt.Printf("per: %v\n", per) // per: {{ person} bob 21 male}
}

func write() {
	p := Person1{
		Name:   "bob",
		Age:    22,
		Gender: "male",
	}
	f, _ := os.OpenFile("a.xml", os.O_CREATE|os.O_WRONLY, 0777)
	defer f.Close()
	e := xml.NewEncoder(f)
	e.Encode(p)
}

func main() {
	// xmlMarshal()
	// xmlUnmarshal()
	// read()
	write()
}
