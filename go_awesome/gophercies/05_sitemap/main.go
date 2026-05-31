package main

import (
	"05_sitemap/crawl"
	"flag"
	"fmt"
	"log"
)

func main() {

	urlAttr := flag.String("startURL", "", "URL to generate sitemap for")
	flag.Parse()
	startURL := *urlAttr

	if startURL == "" {
		log.Fatal("startURL flag is required")
	}
	fmt.Printf("startURL = \"%s\"\n", startURL)

	crawl.Crawl(startURL)

	//fmt.
}
