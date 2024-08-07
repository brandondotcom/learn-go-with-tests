package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}

// CheckWebsitesWithMaxConcurrency does the same thing as CheckWebsites, but limits the number of websites that are
// checked at any given time using a separate **buffered channel** whose length represents the max concurrent requests.
func CheckWebsitesWithMaxConcurrency(wc WebsiteChecker, urls []string, maxConns int) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	requestPool := make(chan bool, maxConns)

	for _, url := range urls {
		go func(u string) {
			requestPool <- true // Sends block when the buffer is full

			resultChannel <- result{u, wc(u)}

			<-requestPool // Receives block only when the buffer is empty
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
