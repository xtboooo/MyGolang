package main

import "fmt"

type Player interface {
	playMusic()
}

type Video interface {
	playVideo()
}

type Mobile1 struct {
}

func (m Mobile1) playMusic() {
	fmt.Println("playMusic...")
}

func (m Mobile1) playVideo() {
	fmt.Println("playVideo...")
}

type Pet1 interface {
	eat3()
}

type Dog1 struct {
}

type Cat struct {
}

func (dog Dog1) eat3() {
	fmt.Println("dog1 eat3...")
}

func (cat Cat) eat3() {
	fmt.Println("cat eat3...")
}

func main() {
	m := Mobile1{}

	m.playMusic()
	m.playVideo()

	/* 	dog1 := Dog1{}
	   	cat := Cat{}
	   	dog1.eat3()
	   	cat.eat3() */

	/* 	var pet Pet1
	   	pet = Dog1{} // dog1 eat3...
	   	pet.eat3()
	   	pet = Cat{} // cat eat3...
	   	pet.eat3() */

}
