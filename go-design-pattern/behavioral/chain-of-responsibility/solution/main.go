package main

import (
	"fmt"
	"log"
)

type Context struct {
	url     string
	content string
	data    any
}

type Handler func(*Context) error

func CheckingURL(ctx *Context) error {
	fmt.Printf("Checking if the URL: %s is valid.\n", ctx.url)
	return nil
}

func FetchingContent(ctx *Context) error {
	fmt.Printf("Fetching the content of the URL: %s\n", ctx.url)
	ctx.content = "Sample content from " + ctx.url
	return nil
}

func ExtractingLinks(ctx *Context) error {
	fmt.Printf("Extracting links from the content of URL: %s\n", ctx.url)
	ctx.data = map[string]string{
		"link1": "https://example.com/link1",
		"link2": "https://example.com/link2",
	}
	return nil
}

func SavingToDatabase(ctx *Context) error {
	fmt.Printf("Saving the information to the database for URL: %s\n", ctx.url)
	return nil
}

// Linked list structure to hold the chain of handlers.
type HandlerNode struct {
	hdl  Handler
	next *HandlerNode
}

func (c *HandlerNode) Handle(url string) error {
	ctx := &Context{url: url}

	if c == nil || c.hdl == nil {
		return nil
	}

	if err := c.hdl(ctx); err != nil {
		return err
	}

	if c.next != nil {
		return c.next.Handle(url)
	}

	return nil
}

func NewCrawler(hdl ...Handler) *HandlerNode {
	var head *HandlerNode
	var current *HandlerNode

	for _, handler := range hdl {
		if head == nil {
			head = &HandlerNode{hdl: handler}
			current = head
		} else {
			current.next = &HandlerNode{hdl: handler}
			current = current.next
		}
	}

	return head
}

type WebCrawler struct {
	chain *HandlerNode
}

func (w *WebCrawler) Crawl(url string) {
	if err := w.chain.Handle(url); err != nil {
		log.Printf("Error crawling URL: %s, error: %v\n", url, err)
	}
}

func main() {
	crawler := &WebCrawler{
		chain: NewCrawler(
			CheckingURL,
			FetchingContent,
			ExtractingLinks,
			SavingToDatabase,
		),
	}
	crawler.Crawl("https://example.com")
}
