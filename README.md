# concurrent-web-scraper
The concurrent web scraper is a command line application that takes a list of URLs as input and concurrently scrapes some data from each HTML page.


## What I Will Learn From This Project

- HTTP Requests: You'll use Go's standard net/http package to make GET requests to a website and read the response body. This is a fundamental skill for any web-related project.

- Goroutines: Instead of scraping one website at a time, you'll launch a separate goroutine for each URL. This allows the program to fetch data from multiple websites simultaneously, significantly speeding up the process.

- Channels: You'll use channels to communicate between the goroutines. As each goroutine finishes scraping a page, it will send the result (the URL and the page title) back to the main goroutine, which will then print it to the console. This is the done channel pattern you've already explored!

- HTML Parsing: You'll need to parse the HTML to find the data tags that you wish to scrape. You can use Go's standard library or a third-partyt library like goquery.


