package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("%s\n\n\n", t)

	now := time.Now()
	fmt.Printf("%s\n\n", now)

	fmt.Println(t.Month())
}
