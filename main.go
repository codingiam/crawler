package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	argc := len(args)

	if argc == 3 {
		baseURL := args[0]

		maxConcurrency, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Error - maxConcurrency: %v", err)
			return
		}

		maxPages, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Printf("Error - maxPages: %v", err)
			return
		}

		cfg, err := configure(baseURL, maxConcurrency, maxPages)
		if err != nil {
			fmt.Printf("Error - configure: %v", err)
			return
		}

		fmt.Println("starting crawl of:", baseURL)

		cfg.wg.Add(1)
		go cfg.crawlPage(baseURL)
		cfg.wg.Wait()

		for normalizedURL, page := range cfg.pages {
			fmt.Printf("%v - %s\n", page.H1, normalizedURL)
		}

		os.Exit(0)
	}

	if argc < 1 {
		fmt.Println("no website provided")
	} else if argc > 3 {
		fmt.Println("too many arguments provided")
	} else {
		fmt.Println("too few arguments provided")
	}

	os.Exit(1)
}
