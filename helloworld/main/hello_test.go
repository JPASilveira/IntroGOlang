package main

import (
	"IntroGOlang"
	"testing"
)

func TestHello(t *testing.T) {
	verify := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result was %s, expected %s", result, expected)
		}
	}

	t.Run("hello to people", func(t *testing.T) {
		result := IntroGOlang.hello("João")
		expected := "Hello, João"
		verify(t, result, expected)
	})

	t.Run("hello to people Spanish", func(t *testing.T) {
		result := IntroGOlang.hello("João", "Spanish")
		expected := "Hola, João"
		verify(t, result, expected)
	})

	t.Run("hello to people Russian", func(t *testing.T) {
		result := IntroGOlang.hello("João", "Russian")
		expected := "привет, João"
		verify(t, result, expected)
	})

	t.Run("empty case hello world", func(t *testing.T) {
		result := IntroGOlang.hello("")
		expected := "Hello, world"
		verify(t, result, expected)
	})
}
