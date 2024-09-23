package main

import (
	"fmt"
	"io"
	"net/http"
)

func Hello(myWritter io.Writer, text string) {
	fmt.Fprintf(myWritter, "%s", text)
}

func HandlerHello(w http.ResponseWriter, r *http.Request) {
	Hello(w, "hello world")
}

func main() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(HandlerHello))
	if err != nil {
		fmt.Println(err)
	}
}
