package main

import (
	"bytes"
	"testing"
)

type SpyMockSleeper struct {
	call int
}

func (m *SpyMockSleeper) Sleep() {
	m.call += 1
}

func TestCountdown(t *testing.T) {

	assertstrings := func(got, want string) {
		if got != want {
			t.Errorf("got %s but expected %s", got, want)
		}
	}

	assertint := func(got, want int) {
		if got != want {
			t.Errorf("got %d but expected %d", got, want)
		}
	}

	t.Run("to test the countdown method", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spy_sleeper := &SpyMockSleeper{}
		Countdown(buffer, spy_sleeper)
		got := buffer.String()
		want := `3
2
1
Go!`
		assertstrings(got, want)
	})

	t.Run("number of times the sleep has been called", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spy_sleeper := &SpyMockSleeper{}
		Countdown(buffer, spy_sleeper)
		got := spy_sleeper.call
		want := 3
		assertint(got, want)
	})
}
