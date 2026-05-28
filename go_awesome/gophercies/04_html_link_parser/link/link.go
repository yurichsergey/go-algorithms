package link

import (
	"io"
	"strings"

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

		text := getText(n)
		if len(text) == 0 {
			text = href
		}

		if len(href) > 0 {
			links = append(links, Link{Href: href, Text: text})
		}
		return links
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, walk(c)...)
	}
	return links
}

func getText(n *html.Node) string {
	if n.Type == html.TextNode {
		return strings.Join(strings.Fields(n.Data), " ")
	}
	if n.Type == html.CommentNode {
		return ""
	}
	var text []string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if t := getText(c); len(t) > 0 {
			text = append(text, t)
		}
	}
	return strings.Join(strings.Fields(strings.Join(text, " ")), " ")
}
