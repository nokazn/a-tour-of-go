package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type Cacher struct {
	mu      sync.Mutex
	visited map[string]bool
}

func New() *Cacher {
	return &Cacher{
		visited: make(map[string]bool),
	}
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, c *Cacher) {
	var wg sync.WaitGroup

	if depth <= 0 {
		return
	}

	c.mu.Lock()
	if c.visited[url] {
		c.mu.Unlock()
		return
	}

	c.visited[url] = true
	c.mu.Unlock()
	body, urls, err := fetcher.Fetch(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	// urls の fetch がすべて終わるまで待つ
	for _, u := range urls {
		wg.Add(1)
		go func(u string) {
			Crawl(u, depth-1, fetcher, c)
			wg.Done()
		}(u)
	}
	wg.Wait()
	return
}

func main() {
	c := New()
	Crawl("https://golang.org/", 4, fetcher, c)
}

type fakeResult struct {
	body string
	urls []string
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	time.Sleep(1 * time.Second)
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}

	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
