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

		const max = 3
		cfg, err := configure(baseURL, max)
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
	} else if argc > 1 {
		fmt.Println("too many arguments provided")
	}

	os.Exit(1)
}
