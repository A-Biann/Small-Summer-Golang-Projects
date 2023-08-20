package main

import (
	"fmt"
	"net/http"
)

type Hello struct {
}

func (h Hello) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "<h1>Hello from server</h1>")
	panic("implement me")
}

func main() {
	var h Hello
	err := http.ListenAndServe("localhost:4000", h)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
