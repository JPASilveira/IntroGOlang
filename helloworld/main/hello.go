package main

import "fmt"

const spanish = "Spanish"
const russian = "Russian"
const helloSpanish = "Hola, "
const helloRussian = "привет, "
const helloDefault = "Hello, "

func Hello(name string, language ...string) string {
	if name == "" {
		name = "World"
	}

	if len(language) > 0 {
		switch language[0] {
		case spanish:
			return helloSpanish + name
		case russian:
			return helloRussian + name
		}
	}
	return helloDefault + name
}

func main() {
	fmt.Println(Hello(""))
}
