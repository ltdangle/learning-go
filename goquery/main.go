package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
)

func main() {
	urlStr := os.Args[1]
	parsedUrlStr, _ := url.Parse(urlStr)
	host := parsedUrlStr.Host

	page := scrapeLinksFromUrl(urlStr)

	// Pretty print the Page struct
	p1, _ := json.MarshalIndent(page, "", "\t")
	fmt.Println("Page links")
	fmt.Print(string(p1))

	// Extract sitemap.
	siteMap := &SiteMap{}
	extractPageLinksToSitemap(siteMap, page, host)

	// Pretty print sitemap.
	p2, _ := json.MarshalIndent(siteMap, "", "\t")
	fmt.Println("")
	fmt.Println("Site links")
	fmt.Print(string(p2))
}
