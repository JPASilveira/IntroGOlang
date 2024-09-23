package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type SleeperSpy struct {
	Calls int
}

func (s *SleeperSpy) Sleep() {
	s.Calls++
}

func Cont(exit io.Writer, sleeper Sleeper) {
	const text = "Lets go!"
	const repeat = 3

	for i := repeat; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(exit, i)
	}
	sleeper.Sleep()
	fmt.Fprint(exit, text)
}

func main() {
	sleeper := &DefaultSleeper{}
	Cont(os.Stdout, sleeper)

}
