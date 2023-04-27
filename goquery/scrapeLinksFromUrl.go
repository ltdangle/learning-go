package main

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
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

	// Find all links
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

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

	// Save page html.
	html, _ := doc.Html()
	err = saveHtmlToFile(pageHost, parsedUrlStr.Path, html)
	if err != nil {
		return err
	}

	return nil
}

func saveHtmlToFile(host string, path string, html string) error {
	if path == "" || path == "/" {
		path = "/index"
	}
	if path[len(path)-1] == '/' {
		path += "/index"
	}

	filePath := "/tmp/" + host + path

	// Create the directory along with any necessary parents.
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		// File exists
		err = os.RemoveAll(filePath)
		if err != nil {
			log.Fatal(err.Error())
		}
		return err
	}

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	// Write the html to the file.
	err = ioutil.WriteFile(filePath, []byte(html), 0644)
	if err != nil {
		return err
	}
	fmt.Println("Saved to " + filePath)
	return nil
}
