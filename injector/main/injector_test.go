package main

import (
	"bytes"
	"testing"
)

func TestInjector(t *testing.T) {
	t.Run("hello world", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Hello(&buffer, "hello world")

		result := buffer.String()
		expected := "hello world"

		if result != expected {
			t.Errorf("expected: %s, got: %s", expected, result)
		}
	})
}
