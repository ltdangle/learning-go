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

func scrapeLinks(urlStr string) {
	parsedUrlStr, _ := url.Parse(urlStr)
	pageHost := parsedUrlStr.Host

	links := make(map[string]map[string]int)

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
		//fmt.Printf("Link href %d: %s\n", i, linkHref)
		urlObj, _ := url.Parse(linkHref)

		// For relative links, set host to pageHost
		if urlObj.Host == "" {
			urlObj.Host = pageHost
		}

		// Check if the nested map for the key "one" has been initialized.
		if links[urlObj.Host] == nil {
			// Initialize the nested map for the key "one".
			links[urlObj.Host] = make(map[string]int)
		}

		// Update links map.
		links[urlObj.Host][urlObj.Path]++
	})

	// Pretty print our found links.
	s, _ := json.MarshalIndent(links, "", "\t")
	fmt.Print(string(s))

}

func main() {
	scrapeLinks(os.Args[1])
}
