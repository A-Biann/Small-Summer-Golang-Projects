package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	dow := rand.Intn(6) + 1
	fmt.Println("Day", dow)

	x := -42
	result := ""
	switch {
	case x < 0:
		result = "less than 0"
	}
	fmt.Println(result)
}
