package main

import "sync"

func main() {
	urls := []string{"https://www.google.com", "https://www.twitch.tv"}
	wg := sync.WaitGroup{}

	for _, link := range urls {
		wg.Add(1)
		go get(link, wg)
	}

	wg.Wait()

}
