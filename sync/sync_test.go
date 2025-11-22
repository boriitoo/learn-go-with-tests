package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increment the counter three times leaves it at three", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCount(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wanted := 3000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wanted)

		for range wanted {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()

		assertCount(t, counter, wanted)
	})
}

func assertCount(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
