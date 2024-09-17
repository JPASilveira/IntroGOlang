package main

import "math"

// Interfaces
type Form interface {
	Area() float64
}

// Estruturas
type Circle struct {
	radius float64
}

type Rectangle struct {
	Width, Height float64
}

type Triangle struct {
	base, height float64
}

// Métodos
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (t Triangle) Area() float64 {
	return (t.base * t.height) / 2
}
