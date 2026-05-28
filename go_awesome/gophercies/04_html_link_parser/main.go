package main

import (
	"fmt"
	"log"
	"os"

	"github.com/yurichsergey/go-algorithms/go_awesome/gophercies/04_html_link_parser/link"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <html_file>")
	}
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	links, err := link.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	for _, url := range links {
		fmt.Printf("url = \"%s\", %s\n", url.Href, url.Text)
	}
}
