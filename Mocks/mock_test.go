package main

import (
	"bytes"
	"testing"

	"golang.org/x/exp/slices"
)

type SpyCountOperations struct {
	calls []string
}

func (m *SpyCountOperations) Sleep() {
	m.calls = append(m.calls, "sleep")
}

func (m *SpyCountOperations) Write(p []byte) (n int, e error) {
	m.calls = append(m.calls, "write")
	return
}

func TestCountdown(t *testing.T) {

	assertstrings := func(got, want string) {
		if got != want {
			t.Errorf("got %s but expected %s", got, want)
		}
	}

	t.Run("to test the countdown method", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spy_sleeper := &SpyCountOperations{}
		Countdown(buffer, spy_sleeper)
		got := buffer.String()
		want := `3
2
1
Go!`
		assertstrings(got, want)
	})

	t.Run("number of times the sleep has been called", func(t *testing.T) {
		//buffer := &bytes.Buffer{}
		spy_operation_counter := &SpyCountOperations{}
		Countdown(spy_operation_counter, spy_operation_counter)
		got := spy_operation_counter.calls
		want := []string{`write sleep write sleep write sleep write`}
		if slices.Equal(got, want) {
			t.Errorf("operations are not in correct sequence got %s, wanted %s", got, want)
		}
	})
}
