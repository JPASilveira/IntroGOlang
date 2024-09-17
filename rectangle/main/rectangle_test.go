package main

import "testing"

func TestRectangle(t *testing.T) {
	verify := func(t *testing.T, result, expected float64) {
		t.Helper()
		if result != expected {
			t.Errorf("result should be %.2f, but %.2f", expected, result)
		}
	}

	t.Run("perimeter", func(t *testing.T) {
		result := Perimeter(Rectangle{10.0, 10.0})
		expected := 40.0
		verify(t, result, expected)
	})

	t.Run("area", func(t *testing.T) {
		result := Area(Rectangle{10.0, 10.0})
		expected := 100.0
		verify(t, result, expected)
	})
}
