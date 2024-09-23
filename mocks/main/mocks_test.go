package main

import (
	"bytes"
	"testing"
)

func TestMocks(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeperSpy := &SleeperSpy{}

	Cont(buffer, sleeperSpy)

	result := buffer.String()
	expected := `3
2
1
Lets go!`

	if result != expected {
		t.Errorf("Expected\n%s\ngot\n%s", expected, result)
	}

	if sleeperSpy.Calls != 4 {
		t.Errorf("SleeperSpy.Calls should be called 4 times")
	}
}
