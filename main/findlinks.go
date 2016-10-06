package main


import (
	"log"
	"os"

	"github.com/barakb/web-crawler/links"
	"runtime"
	"time"
)

var tokens = make(chan struct{}, runtime.NumCPU())

func crawl(url string) []string {
	//log.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	log.Printf("Concurrency level is %d\n", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	start := time.Now()

	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
	elapsed := time.Since(start)

	log.Printf("%d links processed in %s\n", len(seen), elapsed)
}