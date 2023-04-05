package main

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	channelResults := make(chan result)

	for _, url := range urls {
		go func(u string) {
			channelResults <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-channelResults
		results[result.string] = result.bool
	}

	return results
}
