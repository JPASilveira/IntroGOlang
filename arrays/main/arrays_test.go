package main

import (
	"testing"
)

func TestArrays(t *testing.T) {
	t.Run("sum arrays", func(t *testing.T) {
		resulted := sumArray([10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
		expected := 6
		if resulted != expected {
			t.Errorf("got %v, want %v", resulted, expected)
		}
	})

	t.Run("sum multi arrays", func(t *testing.T) {
		resulted := sumSlice(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		expected := 15
		if resulted != expected {
			t.Errorf("got %v, want %v", resulted, expected)
		}
	})
}

func BenchmarkArrays(b *testing.B) {
	b.Run("sum arrays", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sumArray([10]int{156, 2452, 345, 447, 665, 654, 78999, 8343, 9786, 134530})
		}
	})
	b.Run("sum multi arrays", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sumSlice(1156, 2452, 345, 447, 665, 654, 78999, 8343, 9786, 134530)
		}
	})
}
