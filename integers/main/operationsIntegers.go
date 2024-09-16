package main

import "fmt"

func sum(numbers ...int) int {
	var total int
	for _, number := range numbers {
		total += number
	}
	return total
}

func sub(base int, numbers ...int) int {
	for _, number := range numbers {
		base -= number
	}
	return base
}

func multi(base int, numbers ...int) int {
	for _, number := range numbers {
		base *= number
	}
	return base
}

func div(base int, numbers ...int) int {
	for _, number := range numbers {
		base /= number
	}
	return base
}

func rem(base int, numbers ...int) int {
	for _, number := range numbers {
		base %= number
	}
	return base
}

func main() {
	fmt.Println(div(2, 2))
}
