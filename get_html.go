package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "Crawler/0.1")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return "", fmt.Errorf("got HTTP error: %s", resp.Status)
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("got non-HTML response: %s", contentType)
	}

	html, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(html), nil
}
