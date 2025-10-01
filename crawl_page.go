package main

import (
	"fmt"
	"log"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	html, err := getHTML(rawBaseURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(html)
}
