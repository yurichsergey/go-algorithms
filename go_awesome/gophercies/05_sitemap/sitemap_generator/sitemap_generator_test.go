package sitemap_generator

import (
	"strings"
	"testing"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

func newURLSet() XMLUrlSet {
	return XMLUrlSet{XMLNS: xmlns}
}

func TestXMLHeader(t *testing.T) {
	s := newURLSet()
	output, err := s.ToXML([]string{})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(output, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>") {
		t.Errorf("expected XML header, got: %s", output[:50])
	}
}

func TestXMLContainsURLs(t *testing.T) {
	s := newURLSet()
	output, err := s.ToXML([]string{"https://example.com", "https://example.com/about"})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(output, "<loc>https://example.com</loc>") {
		t.Errorf("expected loc tag for example.com, got:\n%s", output)
	}
	if !strings.Contains(output, "<loc>https://example.com/about</loc>") {
		t.Errorf("expected loc tag for /about, got:\n%s", output)
	}
}

func TestXMLEmptyLinks(t *testing.T) {
	s := newURLSet()
	output, err := s.ToXML([]string{})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(output, "<urlset") {
		t.Errorf("expected urlset element, got:\n%s", output)
	}
	if strings.Contains(output, "<url>") {
		t.Errorf("expected no url elements for empty input, got:\n%s", output)
	}
}

func TestXMLContainsXMLNS(t *testing.T) {
	s := newURLSet()
	output, err := s.ToXML([]string{})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(output, xmlns) {
		t.Errorf("expected xmlns attribute, got:\n%s", output)
	}
}
