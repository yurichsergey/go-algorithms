package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"link_parser/link"
)

func main() {

	urlAttr := flag.String("url", "", "URL to generate sitemap for")
	flag.Parse()
	url := *urlAttr

	if url == "" {
		log.Fatal("url flag is required")
	}
	fmt.Printf("url = \"%s\"\n", url)

	links, errParse := extractLinksFromUrl(url)
	if errParse != nil {
		log.Fatal(errParse)
	}
	visited := make(map[string]struct{})
	queue := make([]string, 0)
	for _, url := range links {
		queue = append(queue, url.Href)
		//fmt.Printf("%s - %s\n", url.Href, url.Text)
	}

	for len(queue) > 0 {
		current := queue[0]
		if _, ok := visited[current]; ok {
			continue
		}
		links, err := extractLinksFromUrl(current)
		if err != nil {
			log.Fatal(err)
		}
		for _, url := range links {
			queue = append(queue, url.Href)
		}
		visited[current] = struct{}{}
	}
}

func extractLinksFromUrl(url string) ([]link.Link, error) {
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

	return link.Parse(resp.Body)
}
