package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

// todo: Create struct for ScrapeResult with URL Title and Error
type ScrapeResult struct {
	URL   string
	Title string
	Error error
}

// todo: Create function for ScrapeWorker that takes in url(string) and results(writes to channel)
// 1: Fetch the URL content (http.Get)
// 2: Parse the HTML body (html.Parse)
// 3: Extract the title (extractTitle func)
// 4: Send the successful result back
func ScrapeWorker(url string, results chan<- ScrapeResult) {
	log.Printf("Starting worker for: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		// Send error back and exit
		results <- ScrapeResult{URL: url, Error: fmt.Errorf("HTTP GET request failed: %w", err)}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		results <- ScrapeResult{URL: url, Error: fmt.Errorf("HTTP response status not 200: %d", resp.StatusCode)}
		return
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		results <- ScrapeResult{URL: url, Error: fmt.Errorf("HTML parsing failed: %w", err)}
		return
	}

	title := extractTitle(doc)

	results <- ScrapeResult{
		URL:   url,
		Title: title,
		Error: nil,
	}
	log.Printf("Finished worker for: %s", url)

}

// todo: Create helper function to recursively traverse the HTML doc and find the title tag.
// This will take in var of type *html.Node and return a string
//
//	for c := n.FirstChild; c != nil; c = c.NextSibling {
//		// Recursively call extractTitle on child nodes
//		if title := extractTitle(c); title != "" {
//			return title
//		}
//	}
//
// return ""
func extractTitle(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "title" {
		if n.FirstChild != nil {
			return n.FirstChild.Data
		}
		return ""
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// Recursively call extractTitle on all child nodes
		if title := extractTitle(c); title != "" {
			return title
		}
	}
	return ""
}

// todo: Main function
// Get URL or list of URLS
// Create buffered channel to hold the results, size = to number of URLS
// For each URL, launch a worker goroutine
// Collect all the results from the channels
