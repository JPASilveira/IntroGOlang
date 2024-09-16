package main

import "testing"

func TestOla(t *testing.T) {
	result := hello("João")
	expected := "Hello, João"
	if result != expected {
		t.Errorf("Input: '%s', '%s'", result, expected)
	}
}
