package main

import (
	"flag"
	"fmt"
	"log"

	"05_sitemap/crawl"
	"05_sitemap/sitemap_generator"
)

func main() {

	urlAttr := flag.String("startURL", "", "URL to generate sitemap for")
	flag.Parse()
	startURL := *urlAttr

	if startURL == "" {
		log.Fatal("startURL flag is required")
	}
	fmt.Printf("startURL = \"%s\"\n", startURL)

	links := crawl.Crawl(startURL)

	xmlUrlSet := sitemap_generator.XMLUrlSet{XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9"}
	xmlOutput, err := xmlUrlSet.ToXML(links)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n\n=====\n", xmlOutput)
}
