package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	argc := len(args)

	if argc == 1 {
		baseURL := args[0]

		fmt.Println("starting crawl of:", baseURL)

		html, err := getHTML(baseURL)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(html)

		os.Exit(0)
	}

	if argc < 1 {
		fmt.Println("no website provided")
	} else if argc > 1 {
		fmt.Println("too many arguments provided")
	}

	os.Exit(1)
}
