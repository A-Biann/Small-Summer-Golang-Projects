package main

import (
	"fmt"
	//	"strings"
)

func main() {
	str1 := "An implicit string"
	fmt.Printf("str1: %v: %T\n\n", str1, str1)
	str2 := "An explicit string"
	fmt.Printf("str2: %v: %T\n\n", str2, str2)
}
