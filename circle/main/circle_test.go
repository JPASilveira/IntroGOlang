package main

import "testing"

func TestCircle(t *testing.T) {
	verify := func(t *testing.T, result, expected float64) {
		t.Helper()
		if result != expected {
			t.Errorf("result %.2f != expected %.2f", result, expected)
		}
	}

	t.Run("area circle", func(t *testing.T) {
		circle := Circle{10}
		result := circle.Area()
		expect := 314.1592653589793
		verify(t, result, expect)
	})
}

func BenchmarkCircle(b *testing.B) {
	circle := Circle{10}
	for i := 0; i < b.N; i++ {
		circle.Area()
	}
}
