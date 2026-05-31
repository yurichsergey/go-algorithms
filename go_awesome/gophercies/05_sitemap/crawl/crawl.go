package crawl

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"link_parser/link"
)

func Crawl(startURL string) []string {
	parsedStartURL, err := url.Parse(startURL)
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

	var visitedHrefs []string
	for visitedURL := range visited {
		fmt.Println(visitedURL)
		visitedHrefs = append(visitedHrefs, visitedURL)
	}
	return visitedHrefs
}

func downloadAndAnalyzeURL(currentHref string, parsedStartURL *url.URL, queue []string) []string {
	currentLinks, errExtract := extractLinksFromUrl(currentHref)
	if errExtract != nil {
		log.Printf("Error extracting links from URL \"%s\": %s", currentHref, errExtract)
		return queue
	}
	for _, currentLink := range currentLinks {
		fmt.Printf("Processing URL \"%s\"\n", currentLink.Href)
		parsedCurrentURL, errParse := url.Parse(currentLink.Href)
		if errParse != nil {
			log.Printf("Error parsing URL \"%s\": %s", currentLink.Href, errParse)
		}
		if parsedCurrentURL.Host != parsedStartURL.Host && parsedCurrentURL.Host != "" {
			continue
		}
		urlToAdd := parsedStartURL.ResolveReference(parsedCurrentURL).String()
		fmt.Printf("Adding URL \"%s\" to the queue\n", urlToAdd)
		queue = append(queue, urlToAdd)
	}
	return queue
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
