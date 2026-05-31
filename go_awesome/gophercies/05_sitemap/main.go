package main

import (
	"05_sitemap/crawl"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
)

type XMLUrl struct {
	Loc string `xml:"loc"`
}

type XMLUrlSet struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNS   string   `xml:"xmlns,attr"`
	URLs    []XMLUrl `xml:"url"`
}

func (x *XMLUrlSet) toXML(links []string) (string, error) {
	for _, link := range links {
		x.URLs = append(x.URLs, XMLUrl{Loc: link})
	}

	output, err := xml.MarshalIndent(x, "", " ")
	return xml.Header + string(output), err
}

func main() {

	urlAttr := flag.String("startURL", "", "URL to generate sitemap for")
	flag.Parse()
	startURL := *urlAttr

	if startURL == "" {
		log.Fatal("startURL flag is required")
	}
	fmt.Printf("startURL = \"%s\"\n", startURL)

	links := crawl.Crawl(startURL)

	xmlUrlSet := XMLUrlSet{XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9"}
	xmlOutput, err := xmlUrlSet.toXML(links)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n\n=====\n", xmlOutput)
}
