package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	content := "Hello from Go!"

	file, err := os.Create("./try.txt")
	checkError(err)
	defer file.Close()

	ln, err := io.WriteString(file, content)
	checkError(err)

	fmt.Printf("All done with %v char\n", ln)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
