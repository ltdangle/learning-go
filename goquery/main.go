package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Link struct {
	Count   int  `json:"count"`
	Visited bool `json:"visited"`
}

type DomainMap struct {
	Domain string          `json:"domain"`
	URLs   map[string]Link `json:"urls"`
}

type Page struct {
	Links map[string]DomainMap
}

func scrapeLinksFromUrl(urlStr string) Page {
	parsedUrlStr, _ := url.Parse(urlStr)
	pageHost := parsedUrlStr.Host

	pageLinks := make(map[string]DomainMap)

	// Request the HTML page.
	res, err := http.Get(urlStr)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find all links
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the linkHref.
		linkHref, _ := s.Attr("href")
		urlObj, _ := url.Parse(linkHref)

		// For relative links, set host to pageHost
		if urlObj.Host == "" {
			urlObj.Host = pageHost
		}

		// Check if the nested map for the key "one" has been initialized.
		if pageLinks[urlObj.Host].URLs == nil {
			// Initialize the nested map for the key "one".
			pageLinks[urlObj.Host] = DomainMap{
				Domain: urlObj.Host,
				URLs:   make(map[string]Link),
			}
		}

		// Update links map.
		link := pageLinks[urlObj.Host].URLs[urlObj.Path]
		link.Count++
		pageLinks[urlObj.Host].URLs[urlObj.Path] = link
	})

	// Create a Page struct with the pageLinks
	page := Page{
		Links: pageLinks,
	}

	return page
}

func main() {
	page := scrapeLinksFromUrl(os.Args[1])

	// Pretty print the Page struct
	s, _ := json.MarshalIndent(page, "", "\t")
	fmt.Print(string(s))
}
