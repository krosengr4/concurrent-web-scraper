package main

// --- PROJECT STEPS ---

// Step 1: Read Input: The program should accept a list of URLs, either from the command line or from a text file.

// Step 2: Create a Worker: Write a function that takes a URL as input. This function will be the "worker" goroutine.

// Step 3: Fetch and Parse: Inside the worker function, use http.Get to fetch the HTML content. Then, parse the HTML to find the page's title.

// Step 4: Use Channels: Create a channel to send the results back to the main goroutine. Your worker function will send a struct containing the URL and its title to this channel.

// Step 5: Main Goroutine: The main function will launch a goroutine for each URL and then wait for all of them to finish by receiving the results from the channel.

// Step 6: Print Results: Finally, print the URL and its corresponding title in a clean, formatted way.
