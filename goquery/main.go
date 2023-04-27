package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
)

func main() {
	urlStr := os.Args[1]
	parsedUrlStr, _ := url.Parse(urlStr)
	host := parsedUrlStr.Host
	scheme := parsedUrlStr.Scheme
	siteMapper := NewSiteMapper(scheme, host)

	// Scrape start page.
	err := scrapeLinksFromUrl(siteMapper, urlStr)
	if err != nil {
		log.Fatal("Could not scrape " + urlStr)
	}

	// Crawl site.
	for {
		if siteMapper.uncrawledLinksRemain() {
			link := siteMapper.nextUncrawledLink()
			linkUrl := siteMapper.Scheme + "://" + siteMapper.Host + link.Path
			fmt.Println("Scraping " + linkUrl)
			err := scrapeLinksFromUrl(siteMapper, linkUrl)
			if err != nil {
				fmt.Println("Could not scrape " + linkUrl)
			}
			link.Status.Visited = true
		} else {
			break
		}
	}

	// Pretty print sitemap.
	p2, _ := json.MarshalIndent(siteMapper, "", "\t")
	fmt.Println("")
	fmt.Println("Site links")
	fmt.Print(string(p2))
}
