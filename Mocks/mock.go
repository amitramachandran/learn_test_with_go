package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type SpySleeper struct{}

type Sleeper interface {
	Sleep()
}

func (s SpySleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(b io.Writer, s Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(b, i)
		//time.Sleep(1 * time.Second)
		s.Sleep()
	}
	fmt.Fprintf(b, "Go!")

}

func main() {
	s := SpySleeper{}
	Countdown(os.Stdout, s)
}
