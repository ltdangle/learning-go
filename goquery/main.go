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
	siteMapper := NewSiteMapper(host)

	page := scrapeLinksFromUrl(siteMapper, urlStr)

	// Pretty print the Page struct
	p1, _ := json.MarshalIndent(page, "", "\t")
	fmt.Println("Page links")
	fmt.Print(string(p1))

	// Extract sitemap.
	extractPageLinksToSitemap(siteMapper.SiteMap, page, host)

	// Pretty print sitemap.
	p2, _ := json.MarshalIndent(siteMapper, "", "\t")
	fmt.Println("")
	fmt.Println("Site links")
	fmt.Print(string(p2))
}
