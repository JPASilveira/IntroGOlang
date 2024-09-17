package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

func main() {
	fmt.Println(Perimeter(Rectangle{Width: 10, Height: 5}))
}
