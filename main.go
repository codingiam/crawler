package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	argc := len(args)

	if argc == 1 {
		baseURL := args[0]

		fmt.Println("starting crawl of:", baseURL)

		pages := make(map[string]int)

		crawlPage(baseURL, baseURL, pages)

		for normalizedURL, count := range pages {
			fmt.Printf("%d - %s\n", count, normalizedURL)
		}

		os.Exit(0)
	}

	if argc < 1 {
		fmt.Println("no website provided")
	} else if argc > 1 {
		fmt.Println("too many arguments provided")
	}

	os.Exit(1)
}
