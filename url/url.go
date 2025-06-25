package main

import (
	"fmt"
	"net/url"
)

func main() {
	myUrl := "www.google.com"
	fmt.Printf("%T\n", myUrl)

	parsedURL, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println("Cannot convert")
	}
	fmt.Printf("%T\n", parsedURL)
	fmt.Println("Scheme: ", parsedURL.Scheme)
	fmt.Println("Host: ", parsedURL.Host)
	fmt.Println("Path: ", parsedURL.Path)
	fmt.Println("RawQuery:", parsedURL.RawQuery)
}
