package link

import (
	"io"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	return walk(doc), nil
}

func walk(n *html.Node) []Link {
	var links []Link
	if n.Type == html.ElementNode && n.Data == "a" {
		href := ""
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				href = attr.Val
			}
		}

		if len(href) > 0 && href[0] == '#' {
			href = ""
		}

		text := ""
		// @todo when some elements there
		if n.FirstChild == n.LastChild && n.FirstChild.Type == html.TextNode {
			text = n.FirstChild.Data
		}

		if len(href) > 0 && len(text) > 0 {
			links = append(links, Link{Href: href, Text: text})
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, walk(c)...)
	}
	return links
}
