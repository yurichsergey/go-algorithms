package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"link_parser/link"
)

type ParsedURL struct {
	WholeURL string
	Scheme   string
	Host     string
	Path     string
}

func main() {

	urlAttr := flag.String("startURL", "", "URL to generate sitemap for")
	flag.Parse()
	startURL := *urlAttr

	if startURL == "" {
		log.Fatal("startURL flag is required")
	}
	fmt.Printf("startURL = \"%s\"\n", startURL)

	crawl(startURL)

	//fmt.
}

func crawl(startURL string) {
	parsedStartURL, err := parseURL(startURL)
	if err != nil {
		log.Fatal(err)
	}

	queue := make([]string, 0)
	queue = append(queue, startURL)

	visited := make(map[string]struct{})
	for len(queue) > 0 {
		currentHref := queue[0]
		queue = queue[1:]
		if _, ok := visited[currentHref]; ok {
			continue
		}
		queue = downloadAndAnalyzeURL(currentHref, parsedStartURL, queue)
		visited[currentHref] = struct{}{}
	}

	fmt.Printf("Visited URLs = %d\n", len(visited))
	for visitedURL := range visited {
		fmt.Println(visitedURL)
	}
}

func downloadAndAnalyzeURL(currentHref string, parsedStartURL *ParsedURL, queue []string) []string {
	currentLinks, errExtract := extractLinksFromUrl(currentHref)
	if errExtract != nil {
		log.Fatal(errExtract)
	}
	for _, currentLink := range currentLinks {
		fmt.Printf("Processing URL \"%s\"\n", currentLink.Href)
		parsedCurrentURL, errParse := parseURL(currentLink.Href)
		if errParse != nil {
			log.Fatal(errParse)
		}
		if parsedCurrentURL.Host != parsedStartURL.Host && parsedCurrentURL.Host != "" {
			continue
		}
		URLToAdd := normalizeURL(parsedStartURL, parsedCurrentURL)
		fmt.Printf("Adding URL \"%s\" to the queue\n", URLToAdd)
		queue = append(queue, URLToAdd)
	}
	return queue
}

func normalizeURL(baseURL *ParsedURL, customURL *ParsedURL) string {
	return baseURL.Scheme + "://" + baseURL.Host + customURL.Path
}

func parseURL(inputURL string) (*ParsedURL, error) {
	u, err := url.Parse(inputURL)
	if err != nil {
		return nil, err
	}
	return &ParsedURL{
		WholeURL: inputURL,
		Scheme:   u.Scheme,
		Host:     u.Host,
		Path:     u.Path,
	}, nil
}

func extractLinksFromUrl(url string) ([]link.Link, error) {
	log.Printf("Fetching links from URL \"%s\"", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	links, err := link.Parse(resp.Body)
	log.Printf("Parsed %d links", len(links))
	return links, err
}
