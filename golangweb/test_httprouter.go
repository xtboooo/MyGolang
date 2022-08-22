/* package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello, %s\n", ps.ByName("name"))
}

func main() {
	r := httprouter.New()
	r.GET("/", Index)
	r.GET("/hello/:name", Hello)
	log.Fatal(http.ListenAndServe(":8080", r))
} */