package main

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

// todo: Create helper function to recursively traverse the HTML doc and find the title tag.
// This will take in var of type *html.Node and return a string
// for c := n.FirstChild; c != nil; c = c.NextSibling {
// 	// Recursively call extractTitle on child nodes
// 	if title := extractTitle(c); title != "" {
// 		return title
// 	}
// }
// return ""

// todo: Main function
// Get URL or list of URLS
// Create buffered channel to hold the results, size = to number of URLS
// For each URL, launch a worker goroutine
// Collect all the results from the channels
