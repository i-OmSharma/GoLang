package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("URL in Go")

	link := "https://example.com"
	fmt.Printf("Type of URL: %T\n", link)

	parsedURL, err := url.Parse(link)
	if err != nil {
		fmt.Println("can't parse url: ", parsedURL)
		return
	}
	fmt.Printf("Type of parsedURL: %T\n", parsedURL) 
	
	fmt.Println("Scheme: ", parsedURL.Scheme)
	fmt.Println("Host: ", parsedURL.Host)
	fmt.Println("Path: ", parsedURL.Path)
	fmt.Println("RawQuery: ", parsedURL.RawQuery)
	
	// Modify URL Components
	parsedURL.Path = "/newpath/"
	parsedURL.RawQuery = "username=hero"
	
	//Create URL
	newURL := parsedURL.String()

	fmt.Println("new URL: ", newURL)
}
