package crawl

import (
	queue_lib "05_sitemap/queue"
	"io"
	"log"
	"net/http"
	"net/url"

	"link_parser/link"
)

type queueItem struct {
	url   string
	depth int
}

func Crawl(startURL string, maxDepth int) []string {
	parsedStartURL, err := url.Parse(startURL)
	if err != nil {
		log.Fatal(err)
	}

	queue := queue_lib.Queue[queueItem]{}
	queue.Enqueue(queueItem{url: startURL, depth: 0})

	visited := make(map[string]struct{})
	for !queue.IsEmpty() {
		current := queue.Dequeue()
		if _, ok := visited[current.url]; ok || current.depth > maxDepth {
			continue
		}
		downloadAndAnalyzeURL(current.url, parsedStartURL, &queue, current.depth)
		visited[current.url] = struct{}{}
	}

	log.Printf("Visited URLs = %d\n", len(visited))

	var visitedHrefs []string
	for visitedURL := range visited {
		log.Println(visitedURL)
		visitedHrefs = append(visitedHrefs, visitedURL)
	}
	return visitedHrefs
}

func downloadAndAnalyzeURL(currentHref string, parsedStartURL *url.URL, queue *queue_lib.Queue[queueItem], currentDepth int) {
	currentLinks, errExtract := extractLinksFromUrl(currentHref)
	if errExtract != nil {
		log.Printf("Error extracting links from URL \"%s\": %s", currentHref, errExtract)
		return
	}
	for _, currentLink := range currentLinks {
		log.Printf("Processing URL \"%s\"\n", currentLink.Href)
		parsedCurrentURL, errParse := url.Parse(currentLink.Href)
		if errParse != nil {
			log.Printf("Error parsing URL \"%s\": %s", currentLink.Href, errParse)
		}
		if parsedCurrentURL.Host != parsedStartURL.Host && parsedCurrentURL.Host != "" {
			continue
		}
		urlToAdd := parsedStartURL.ResolveReference(parsedCurrentURL).String()
		log.Printf("Adding URL \"%s\" to the queue\n", urlToAdd)
		queue.Enqueue(queueItem{url: urlToAdd, depth: currentDepth + 1})
	}
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
