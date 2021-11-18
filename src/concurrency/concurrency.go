package concurrency

type Result struct {
	string
	bool
}
type Checker func (url string) bool

func checkWebsites(checker Checker, urls []string) map[string]bool {
	resultChannel := make(chan Result)
	urlWithResult := make(map[string]bool)
	for _, url := range urls {
		go func(u string) {
			resultChannel <- Result{u, checker(u)}
		}(url)
	}

	for i := 0;i < len(urls);i++ {
		result := <- resultChannel
		urlWithResult[result.string] = result.bool
	}
	return urlWithResult
}
