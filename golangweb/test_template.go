/* package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

func test1() {
	name := "xtbo97"
	templateStr := "hello, {{.}}"
	t := template.New("test")
	t2, err := t.Parse(templateStr)
	if err != nil {
		log.Fatal(err)
	}
	t2.Execute(os.Stdout, name)
}

type Person struct {
	Name string
	Age  int
}

func test2() {
	person := Person{Name: "bob", Age: 25}
	templateStr := "Name:{{.Name}}, Age:{{.Age}}"
	t, err := template.New("test").Parse(templateStr)
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(os.Stdout, person)
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("test4.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, "hello world")
}

func tmpl(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("test6.html")
	if err != nil {
		panic(err)
	}
	s := []string{"a", "b", "c", "d"}
	t.Execute(w, s)
}

func tmpl2(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html", "header.html", "footer.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func testHTML() {
	server := http.Server{
		Addr: "127.0.0.1:8081",
	}
	http.HandleFunc("/hello", tmpl2)
	server.ListenAndServe()
}

func main() {
	// test1()
	// test2()
	testHTML()
}
 */