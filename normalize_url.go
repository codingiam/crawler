package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	u.RawQuery = ""
	u.Fragment = ""
	u.Scheme = ""

	str = strings.TrimSuffix(strings.ToLower(u.String()), "/")

	return str[2:], nil
}
