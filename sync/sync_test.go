package sync

import (
	"sync"
	"testing"
)

func assertCounter(t testing.TB, counter *Counter, want int) {
	t.Helper()
	got := counter.Value()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

// function returns a pointer of the counter object
func NewCounter() *Counter {
	return &Counter{}
}

func TestCounter(t *testing.T) {

	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Increment()
		counter.Increment()
		counter.Increment()

		assertCounter(t, counter, 3)
	})

	t.Run("test for concurrent counter", func(t *testing.T) {
		counter := NewCounter()
		end_count := 1000
		var wg sync.WaitGroup

		wg.Add(end_count)

		for i := 0; i < end_count; i++ {

			go func() {
				counter.Increment()
				wg.Done()
			}()
		}
		wg.Wait()
		assertCounter(t, counter, end_count)
	})
}
