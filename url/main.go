package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL")
	myUrl := "http://example.com:8080/path/to/resource?key1=value1&key2=value2"
	fmt.Printf("Type of URL: %T\n", myUrl)

	parsedURL, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println("Can't parse URL ", err)
		return
	}
	fmt.Printf("Type of URL: %T\n", parsedURL)

	fmt.Println("Scheme: ", parsedURL.Scheme)
	fmt.Println("Host: ", parsedURL.Host)
	fmt.Println("Path: ", parsedURL.Path)
	fmt.Println("RawQuery: ", parsedURL.RawQuery)

	// Modifying URL components
	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "username=iamshobhit"

	// Constructing a URL string from a URL object
	newURL := parsedURL.String()

	fmt.Println("New URL: ", newURL)
}
