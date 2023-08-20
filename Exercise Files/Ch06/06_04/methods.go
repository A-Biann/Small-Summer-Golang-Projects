package main

import "fmt"

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func main() {
	poodle := Dog{"Poodle", 34, "Woof"}
	fmt.Println(poodle)
	poodle.Speak()
}

func (d Dog) Speak() {
	fmt.Println(d.Sound)
}
