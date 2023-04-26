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

type HostLinks struct {
	Host string           `json:"domain"`
	URLs map[string]*Link `json:"urls"`
}

func scrapeLinksFromUrl(urlStr string) {
	parsedUrlStr, _ := url.Parse(urlStr)
	pageHost := parsedUrlStr.Host

	var pageLinks []*HostLinks

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

		hostLinkMap := findHostLinks(pageLinks, urlObj.Host)
		if hostLinkMap == nil {
			hostLinkMap = &HostLinks{
				Host: urlObj.Host,
				URLs: make(map[string]*Link),
			}
			pageLinks = append(pageLinks, hostLinkMap)
		}

		link, ok := hostLinkMap.URLs[urlObj.Path]
		if !ok {
			link = &Link{Count: 0, Visited: false}
			hostLinkMap.URLs[urlObj.Path] = link
		}
		link.Count++
	})

	// Pretty print our found links.
	s, _ := json.MarshalIndent(pageLinks, "", "\t")
	fmt.Print(string(s))
}

func findHostLinks(hostLinks []*HostLinks, domain string) *HostLinks {
	for _, linkMap := range hostLinks {
		if linkMap.Host == domain {
			return linkMap
		}
	}
	return nil
}

func main() {
	scrapeLinksFromUrl(os.Args[1])
}
