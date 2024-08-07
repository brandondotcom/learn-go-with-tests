package racer

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func BenchmarkSelectTest(b *testing.B) {
	b.Run("with all channels ready for receive", func(b *testing.B) {
		results := map[string]int{"a": 0, "b": 0, "c": 0, "d": 0}
		for i := 0; i < b.N; i++ {
			chanA, chanB, chanC, chanD := make(chan any, 1), make(chan any, 1), make(chan any, 1), make(chan any, 1)

			chanA <- struct{}{}
			chanB <- struct{}{}
			chanC <- struct{}{}
			chanD <- struct{}{}

			got := SelectTest(chanA, chanB, chanC, chanD)

			results[got]++
		}

		//results(n=1): map[a:0 b:0 c:1 d:0 default:0]
		//results(n=100): map[a:30 b:32 c:21 d:17 default:0]
		//results(n=10000): map[a:2470 b:2471 c:2540 d:2519 default:0]
		//results(n=43410): map[a:11062 b:10856 c:10663 d:10829 default:0]
		//results(n=74500): map[a:18490 b:18613 c:18785 d:18612 default:0]

		fmt.Printf("results(n=%d): %v\n", b.N, results)
	})

	b.Run("with some channels ready for receive", func(b *testing.B) {
		results := map[string]int{"a": 0, "b": 0, "c": 0, "d": 0}
		for i := 0; i < b.N; i++ {
			chanA, chanB, chanC, chanD := make(chan any, 1), make(chan any, 1), make(chan any, 1), make(chan any, 1)

			chanA <- struct{}{}
			chanB <- struct{}{}

			got := SelectTest(chanA, chanB, chanC, chanD)

			results[got]++
		}

		fmt.Printf("results(n=%d): %v\n", b.N, results)
	})
	b.Run("with a default statement", func(b *testing.B) {
		results := map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "default": 0}
		for i := 0; i < b.N; i++ {
			chanA, chanB, chanC, chanD := make(chan any, 1), make(chan any, 1), make(chan any, 1), make(chan any, 1)

			got := SelectWithDefaultTest(chanA, chanB, chanC, chanD)

			results[got]++
		}

		fmt.Printf("results(n=%d): %v\n", b.N, results)
	})
}

func BenchmarkSelectWithDefaultTest(b *testing.B) {
	b.Run("with all channels ready for receive", func(b *testing.B) {
		results := map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "default": 0}
		for i := 0; i < b.N; i++ {
			chanA, chanB, chanC, chanD := make(chan any, 1), make(chan any, 1), make(chan any, 1), make(chan any, 1)

			chanA <- struct{}{}
			chanB <- struct{}{}
			chanC <- struct{}{}
			chanD <- struct{}{}

			got := SelectWithDefaultTest(chanA, chanB, chanC, chanD)

			results[got]++
		}

		fmt.Printf("results(n=%d): %v\n", b.N, results)
	})
}

func TestRacer(t *testing.T) {

	t.Run("compares speeds of servers", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("times out after 10s", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
