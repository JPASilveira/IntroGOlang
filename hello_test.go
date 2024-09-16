package main

import "testing"

func TestHello(t *testing.T) {
	verify := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result was %s, expected %s", result, expected)
		}
	}

	t.Run("hello to people", func(t *testing.T) {
		result := hello("João")
		expected := "Hello, João"
		verify(t, result, expected)
	})

	t.Run("empty case hello world", func(t *testing.T) {
		result := hello("")
		expected := "Hello, world"
		verify(t, result, expected)
	})
}
