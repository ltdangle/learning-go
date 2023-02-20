package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

type SiteMap map[string][]string

func main() {
	// Set the starting URL for the crawl
    startURL := "https://pkg.go.dev/golang.org/x/net/html"

	// Create a new sitemap
	siteMap := make(SiteMap)

	// Crawl the site
	crawl(startURL, startURL, siteMap)

	// Print the sitemap
	for page, links := range siteMap {
		fmt.Printf("%s\n", page)
		for _, link := range links {
			fmt.Printf("  -> %s\n", link)
		}
	}
}

func crawl(urlStr, rootURL string, siteMap SiteMap) {
	// Parse the URL
	u, err := url.Parse(urlStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing URL %s: %v\n", urlStr, err)
		return
	}

	// Check if the URL has already been crawled
	if _, ok := siteMap[urlStr]; ok {
		return
	}

	// Fetch the page content
	resp, err := http.Get(urlStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching URL %s: %v\n", urlStr, err)
		return
	}
	defer resp.Body.Close()

	// Check if the response was successful
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Error fetching URL %s: %s\n", urlStr, resp.Status)
		return
	}

	// Read the page content
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading response body for URL %s: %v\n", urlStr, err)
		return
	}

	// Extract links from the page
	links := extractLinks(string(body), u.Host)

	// Add the page and links to the sitemap
	siteMap[urlStr] = links

	// Recursively crawl the linked pages
	for _, link := range links {
		if link == rootURL {
			continue
		}
		crawl(link, rootURL, siteMap)
	}
}

func extractLinks(pageHTML, host string) []string {
	var links []string
	z := html.NewTokenizer(strings.NewReader(pageHTML))
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			// End of the document
			return links
		case tt == html.StartTagToken:
			t := z.Token()
			if t.Data == "a" {
				for _, attr := range t.Attr {
					if attr.Key == "href" {
						u, err := url.Parse(attr.Val)
						if err == nil {
							if u.IsAbs() && u.Host == host {
								links = append(links, u.String())
								fmt.Println("Found link: ", u.String())
							} else if !u.IsAbs() {
								links = append(links, "https://"+host+filepath.Join("/", u.Path))
								fmt.Println("Found link: ", "https://"+host+filepath.Join("/", u.Path))
							}
						}
					}
				}
			}
		}
	}
}
