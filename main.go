package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type ScrapeResult struct {
	URL   string
	Title string
	Error error
}

// Takes URL and writes title to channel
func ScrapeWorker(url string, results chan<- ScrapeResult) {
	log.Printf("Starting worker for: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		// Send error back and exit
		results <- ScrapeResult{URL: url, Error: fmt.Errorf("HTTP GET request failed: %w", err)}
		return
	}
	defer resp.Body.Close()

	// If the HTTP response status is not OK
	if resp.StatusCode != http.StatusOK {
		results <- ScrapeResult{URL: url, Error: fmt.Errorf("HTTP response status not 200: %d", resp.StatusCode)}
		return
	}

	// Parse the html response body
	doc, err := html.Parse(resp.Body)
	if err != nil {
		results <- ScrapeResult{URL: url, Error: fmt.Errorf("HTML parsing failed: %w", err)}
		return
	}

	// Extract the title from the body
	title := extractTitle(doc)

	// Return results
	results <- ScrapeResult{
		URL:   url,
		Title: title,
		Error: nil,
	}
	log.Printf("Finished worker for: %s", url)

}

// Helper function to extract the title from a HTML body
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

func main() {
	fmt.Println("\n\t\t======WELCOME TO THE HTML TITLE SCRAPER======")
	fmt.Println(strings.Repeat("_", 20))

	scanner := bufio.NewScanner(os.Stdin)

	// Get user HTML input
	fmt.Println("How many URLs do you wish to enter:")
	var userUrlQuantity int
	fmt.Scanln(&userUrlQuantity)

	htmlList := []string{}

	for i := 0; i < userUrlQuantity; i++ {
		fmt.Println("Enter an URL:")
		scanner.Scan()
		userHtml := scanner.Text()
		htmlList = append(htmlList, userHtml)

	}

	// Buffered Channel that will record the ScrapeResult
	results := make(chan ScrapeResult, len(htmlList))

	// Launch worker goroutine for each html
	for _, url := range htmlList {
		go ScrapeWorker(url, results)
	}

	fmt.Println("Started scraping. Waiting for all workers to finish...")

	// Collect the results from this channel
	for i := 0; i < len(htmlList); i++ {
		result := <-results //<--- This blocks until a result is available

		if result.Error != nil {
			fmt.Printf("ERROR! %s - %v\n", result.URL, result.Error)
		} else {
			cleanTitle := strings.TrimSpace(result.Title)

			fmt.Println(strings.Repeat("-", 80))
			fmt.Printf("Title for %s: %s\n", result.URL, cleanTitle)
			fmt.Println(strings.Repeat("-", 80))
		}
	}

	fmt.Println("All Done")
	fmt.Println(strings.Repeat("_", 20))

}
