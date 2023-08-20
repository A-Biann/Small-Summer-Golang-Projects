package main

import (
	"fmt"
)

func main() {
	states := make(map[string]string)
	fmt.Println(states)

	states["WA"] = "washington"
	states["OR"] = "Oregon"
	fmt.Println(states)

}