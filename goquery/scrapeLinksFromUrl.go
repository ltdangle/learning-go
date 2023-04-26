package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type PageLink struct {
	Count   int  `json:"count"`
	Visited bool `json:"visited"`
}

type HostLinks struct {
	Host string              `json:"host"`
	Urls map[string]PageLink `json:"urls"`
}

type Page struct {
	Links map[string]HostLinks
}

func scrapeLinksFromUrl(siteMapper *SiteMapper, urlStr string) Page {
	fmt.Println("Scraping " + urlStr)
	parsedUrlStr, _ := url.Parse(urlStr)
	pageHost := parsedUrlStr.Host

	pageLinks := make(map[string]HostLinks)

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

		if urlObj.Path == "" {
			urlObj.Path = "/"
		}

		if !strings.HasPrefix(urlObj.Path, "/") {
			urlObj.Path = "/" + urlObj.Path
		}

		// Check if the nested map has been initialized.
		if pageLinks[urlObj.Host].Urls == nil {
			// Initialize the nested map.
			pageLinks[urlObj.Host] = HostLinks{
				Host: urlObj.Host,
				Urls: make(map[string]PageLink),
			}
		}

		// Update links map.
		link := pageLinks[urlObj.Host].Urls[urlObj.Path]
		link.Count++
		pageLinks[urlObj.Host].Urls[urlObj.Path] = link

		// Update site map.
		if urlObj.Host == siteMapper.Host {
			siteMapper.addPathToSiteMap(urlObj.Host, urlObj.Path)
		}

	})

	// Create a Page struct with the pageLinks
	page := Page{
		Links: pageLinks,
	}

	return page
}
