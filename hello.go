package main

import "fmt"

func hello(name string) string {
	if name == "" {
		return "Hello, world"
	}

	return "Hello, " + name
}

func sum(a, b float32) float32 {
	return a + b
}

func main() {
	fmt.Println(hello(""))
}
