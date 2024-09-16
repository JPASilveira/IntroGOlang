package main

func sumArray(numbers [10]int) int {
	var total int
	for _, number := range numbers {
		total += number
	}
	return total
}

func sumSlice(numbers ...int) int {
	var total int
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {

}
