package counter

import (
	"sync"
	"testing"
)

func TestSharedInt(t *testing.T) {

	wantedCount := 1000
	x := &SharedInt{}

	var wg sync.WaitGroup
	wg.Add(wantedCount)

	for i := 0; i < wantedCount; i++ {
		go func() {
			x.Inc()
			wg.Done()
		}()
	}
	wg.Wait()

	assertCounter(t, x, wantedCount) // sanity check
}

func BenchmarkCounter(b *testing.B) {
	benchmarkCounter := func(b *testing.B, c Counter) {
		b.Helper()
		var wg sync.WaitGroup
		wg.Add(b.N)

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			go func() {
				c.Inc()
				wg.Done()
			}()
		}

		wg.Wait()
		b.StopTimer()

		assertCounter(b, c, b.N) // sanity check
	}
	b.Run("Inc() with a mutex", func(b *testing.B) {
		counter := NewCounter()
		benchmarkCounter(b, counter)
	})

	b.Run("Inc() with a buffered chan", func(b *testing.B) {
		counter := NewCounterWithChan()
		benchmarkCounter(b, counter)
	})

	b.Run("Inc() with an unbuffered chan", func(b *testing.B) {
		counter := NewCounterWithUnbufferedChan()
		benchmarkCounter(b, counter)
	})

	//b.Run("initialize with a mutex", func(b *testing.B) {
	//	b.ReportAllocs()
	//	for i := 0; i < b.N; i++ {
	//		NewCounter()
	//	}
	//})
	//
	//b.Run("initialize with a chan", func(b *testing.B) {
	//	b.ReportAllocs()
	//	for i := 0; i < b.N; i++ {
	//		NewCounterWithChan()
	//	}
	//})
}

func TestCounter(t *testing.T) {

	t.Run("incrementing the counter 3 times", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})

	t.Run("it runs when you use channel synchronization", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounterWithChan()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})

	t.Run("it runs when you use channel synchronization", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounterWithChan()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})

	t.Run("it runs when you use unbuffered channel synchronization", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounterWithUnbufferedChan()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, got Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
