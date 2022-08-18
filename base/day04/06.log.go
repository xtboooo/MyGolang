package main

import (
	"fmt"
	"log"
	"os"
)

var logger *log.Logger

func t1() {
	log.Print("my log")
	log.Printf("my log %d", 100)
	name := "tom"
	age := 20
	log.Println(name, " ", age)
}

func t2() {
	defer fmt.Println("panic 结束之后再执行...")
	log.Print("my log")
	log.Panic("my panic")
	fmt.Println("end...")
}

func t3() {
	defer fmt.Println("defer...")
	log.Print("my log")
	log.Fatal("fatal...") // os.exit(1)
	fmt.Println("end...")
}

func init() {
	/* 	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	   	log.SetPrefix("My Log: ")
	   	f, err := os.OpenFile("a.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
	   	// defer f.Close()
	   	if err != nil {
	   		log.Fatal(err)
	   	}
	   	log.SetOutput(f) */

	f, err := os.OpenFile("a.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
	// defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	logger = log.New(f, "My log: ", log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	// t1()
	// t2()
	// t3()
	i := log.Flags()
	fmt.Printf("i: %v\n", i)
	logger.Print("my log")
}
