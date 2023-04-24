package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <URL> <SaveDirectory>")
		os.Exit(1)
	}

	startURL := os.Args[1]
	saveDir := os.Args[2]

	visited := make(map[string]bool)
	err := scrape(startURL, visited, saveDir)
	if err != nil {
		fmt.Println("Error scraping:", err)
	}

	fmt.Println("Sitemap:")
	for url := range visited {
		fmt.Println(url)
	}
}

func scrape(urlStr string, visited map[string]bool, saveDir string) error {
	if visited[urlStr] {
		return nil
	}
	visited[urlStr] = true

	fmt.Println("Crawling:", urlStr)

	resp, err := http.Get(urlStr)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	content := string(body)
	filePath, err := storePage(baseURL, content, saveDir)
	if err != nil {
		return err
	}

	fmt.Printf("Stored: %s -> %s\n", urlStr, filePath)

	links := extractLinks(baseURL, content)

	for _, link := range links {
		err := scrape(link, visited, saveDir)
		if err != nil {
			fmt.Printf("Error scraping link (%s): %s\n", link, err)
		}
	}

	return nil
}

func storePage(baseURL *url.URL, content string, saveDir string) (string, error) {
	mirrorRoot := filepath.Join(saveDir, baseURL.Host)

	title := extractTitle(content)
	title = sanitizeFilename(title)
	filePath := filepath.Join(mirrorRoot, title+".html")

	err := os.MkdirAll(mirrorRoot, os.ModePerm)
	if err != nil {
		return "", err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	io.Copy(file, strings.NewReader(content))

	return filePath, nil
}

func sanitizeFilename(filename string) string {
	invalidChars := regexp.MustCompile(`[<>:"/\\|?*]`)
	return invalidChars.ReplaceAllString(filename, "")
}

func extractTitle(content string) string {
	titlePattern := regexp.MustCompile(`(?i)<title>(.*?)<\/title>`)
	matches := titlePattern.FindStringSubmatch(content)
	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}
	return "untitled"
}

func extractLinks(baseURL *url.URL, content string) []string {
	linkPattern := regexp.MustCompile(`(?i)href="(http[s]?://.*?)"`)
	matches := linkPattern.FindAllStringSubmatch(content, -1)

	var links []string
	for _, match := range matches {
		linkURL, err := url.Parse(match[1])
		if err != nil {
			continue
		}

		if linkURL.Host == baseURL.Host {
			links = append(links, linkURL.String())
		}
	}

	return links
}
