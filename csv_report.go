package main

import (
	"encoding/csv"
	"os"
	"strings"
)

func writeCSVRport(pages map[string]PageData, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	err = writer.Write([]string{"page_url", "h1", "first_paragraph", "outgoing_link_urls", "image_urls"})
	if err != nil {
		return err
	}

	for normalizedURL, page := range pages {
		err = writer.Write([]string{
			normalizedURL,
			page.H1,
			page.FirstParagraph,
			strings.Join(page.OutgoingLinks, ";"),
			strings.Join(page.ImageURLs, ";"),
		})
		if err != nil {
			return err
		}
	}

	return nil
}
