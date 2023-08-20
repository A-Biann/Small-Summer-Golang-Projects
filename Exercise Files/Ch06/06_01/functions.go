package main

import "fmt"

func main() {
	doSomeThing()
	fmt.Println(addValues(1, 2))
}

func doSomeThing() {
	fmt.Println("A new function")
}

func addValues(value1 int, value2 int) int {
	return value1 + value2
}
