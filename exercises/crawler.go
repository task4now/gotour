package main

import (
    "fmt"
	"sync"
    "time"
)


type Fetcher interface {
    Fetch(url string) (body string, urls []string, err error)
}


type SafeMap struct {
    m map[string] string
    mux sync.Mutex
}


func (mp *SafeMap) Crawl(url string, depth int, fetcher Fetcher) {
    if depth <= 0 {
        return
    }

    body, urls, err := fetcher.Fetch(url)

    if err != nil {
        return
    }

    mp.mux.Lock()
    mp.m[url] = body
    mp.mux.Unlock()

    for _, u := range urls {
        go mp.Crawl(u, depth-1, fetcher)
    }

    return
}


func main() {
    mp := SafeMap{m: make(map[string] string)}

    go mp.Crawl("http://golang.org/", 4, fetcher)

    time.Sleep(50 * time.Millisecond)
    for k := range mp.m {
        fmt.Printf("%s: %s\n", k, mp.m[k])
    }
}


type fakeFetcher map[string]*fakeResult


type fakeResult struct {
    body string
    urls []string
}


func (f fakeFetcher) Fetch(url string) (string, []string, error) {
    if res, ok := f[url]; ok {
        return res.body, res.urls, nil
    }

    return "", nil, fmt.Errorf("not found: %s", url)
}


var fetcher = fakeFetcher{
    "http://golang.org/": &fakeResult{
        "The Go Programming Language",
        []string{
            "http://golang.org/pkg/",
            "http://golang.org/cmd/",
        },
    },
    "http://golang.org/pkg/": &fakeResult{
        "Packages",
        []string{
            "http://golang.org/",
            "http://golang.org/cmd/",
            "http://golang.org/pkg/fmt/",
            "http://golang.org/pkg/os/",
        },
    },
    "http://golang.org/pkg/fmt/": &fakeResult{
        "Package fmt",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
    "http://golang.org/pkg/os/": &fakeResult{
        "Package os",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
}
