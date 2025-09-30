package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getH1FromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	elm := doc.Find("h1").First().Text()

	return strings.TrimSpace(elm)
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	elm := doc.Find("main").First()
	if elm.Length() > 0 {
		elm = elm.Find("p").First()
	} else {
		elm = doc.Find("p").First()
	}

	return strings.TrimSpace(elm.Text())
}
