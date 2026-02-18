package main

import "fmt"

type WebCrawler struct {
}

func (w *WebCrawler) Crawl(url string) {
	fmt.Printf("Crawling URL: %s\n", url)
	fmt.Println("1. Check if the URL is valid.")
	fmt.Println("2. Fetch the content of the URL.")
	fmt.Println("3. Extract links from the content.")
	fmt.Println("4. Save the information to the database.")
}

func main() {
	crawler := &WebCrawler{}
	crawler.Crawl("https://example.com")
}
