package main

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"strconv"
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

func scrapeLinksFromUrl(siteMapper *SiteMapper, urlStr string) error {

	parsedUrlStr, _ := url.Parse(urlStr)
	pageHost := parsedUrlStr.Host

	pageLinks := make(map[string]HostLinks)

	// Request the HTML page.
	res, err := http.Get(urlStr)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New(strconv.Itoa(res.StatusCode) + " - " + res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
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

		// Update page links.
		if pageLinks[urlObj.Host].Urls == nil {
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

	return nil
}
