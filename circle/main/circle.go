package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func main() {
	c := Circle{radius: 5}
	fmt.Println(c.Area())
}
