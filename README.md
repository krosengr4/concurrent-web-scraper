# concurrent-web-scraper
The concurrent web scraper is a command line application that takes a list of URLs as input and concurrently scrapes some data from each HTML page.


## What I Will Learn From This Project

- HTTP Requests: You'll use Go's standard net/http package to make GET requests to a website and read the response body. This is a fundamental skill for any web-related project.

- Goroutines: Instead of scraping one website at a time, you'll launch a separate goroutine for each URL. This allows the program to fetch data from multiple websites simultaneously, significantly speeding up the process.

- Channels: You'll use channels to communicate between the goroutines. As each goroutine finishes scraping a page, it will send the result (the URL and the page title) back to the main goroutine, which will then print it to the console. This is the done channel pattern you've already explored!

- HTML Parsing: You'll need to parse the HTML to find the data tags that you wish to scrape. You can use Go's standard library or a third-partyt library like goquery.


--- PROJECT STEPS ---

Step 1: Read Input: The program should accept a list of URLs, either from the command line or from a text file.

Step 2: Create a Worker: Write a function that takes a URL as input. This function will be the "worker" goroutine.

Step 3: Fetch and Parse: Inside the worker function, use http.Get to fetch the HTML content. Then, parse the HTML to find the page's title.

Step 4: Use Channels: Create a channel to send the results back to the main goroutine. Your worker function will send a struct containing the URL and its title to this channel.

Step 5: Main Goroutine: The main function will launch a goroutine for each URL and then wait for all of them to finish by receiving the results from the channel.

Step 6: Print Results: Finally, print the URL and its corresponding title in a clean, formatted way.