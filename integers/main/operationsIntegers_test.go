package main

import (
	"testing"
)

func TestInt(t *testing.T) {
	verify := func(t *testing.T, result, expected int) {
		t.Helper()
		if result != expected {
			t.Errorf("result: %d, expected: %d", result, expected)
		}
	}

	t.Run("sum integers", func(t *testing.T) {
		resulted := sum(1, 2, 3)
		expected := 6
		verify(t, resulted, expected)
	})

	t.Run("subtract integers", func(t *testing.T) {
		resulted := sub(10, 2, 5)
		expected := 3
		verify(t, resulted, expected)
	})

	t.Run("multiply integers", func(t *testing.T) {
		resulted := multi(2, 2)
		expected := 4
		verify(t, resulted, expected)
	})

	t.Run("divide integers", func(t *testing.T) {
		resulted := div(2, 2)
		expected := 1
		verify(t, resulted, expected)
	})

	t.Run("remainder integers", func(t *testing.T) {
		resulted := rem(2, 2)
		expected := 0
		verify(t, resulted, expected)
	})
}
