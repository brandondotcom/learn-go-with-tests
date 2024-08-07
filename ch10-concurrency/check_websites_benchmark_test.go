package concurrency

import (
	"testing"
	"time"
)

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	b.ReportAllocs()

	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

func BenchmarkCheckWebsitesWithMaxConcurrency(b *testing.B) {
	b.Run("with one request at a time", func(b *testing.B) {
		b.ReportAllocs()

		urls := make([]string, 100)
		for i := 0; i < len(urls); i++ {
			urls[i] = "a url"
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			CheckWebsitesWithMaxConcurrency(slowStubWebsiteChecker, urls, 1)
		}
	})

	b.Run("with all requests at once", func(b *testing.B) {
		b.ReportAllocs()

		urls := make([]string, 100)
		for i := 0; i < len(urls); i++ {
			urls[i] = "a url"
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			CheckWebsitesWithMaxConcurrency(slowStubWebsiteChecker, urls, len(urls))
		}
	})
}
