package crawl

import (
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"
)

// newTestServer creates a local HTTP server with predefined pages.
// pages maps path -> HTML content.
func newTestServer(pages map[string]string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if content, ok := pages[r.URL.Path]; ok {
			w.Header().Set("Content-Type", "text/html")
			_, err := w.Write([]byte(content))
			if err != nil {
				return
			}
		} else {
			http.NotFound(w, r)
		}
	}))
}

func TestCrawlSinglePage(t *testing.T) {
	srv := newTestServer(map[string]string{
		"/": `<html><body><p>No links here</p></body></html>`,
	})
	defer srv.Close()

	pages := Crawl(srv.URL+"/", 10)
	if len(pages) != 1 {
		t.Errorf("expected 1 page, got %d: %v", len(pages), pages)
	}
}

func TestCrawlFollowsSameDomainLinks(t *testing.T) {
	srv := newTestServer(map[string]string{
		"/": `<html><body>
			<a href="/about">About</a>
			<a href="/contact">Contact</a>
		</body></html>`,
		"/about":   `<html><body><a href="/">Home</a></body></html>`,
		"/contact": `<html><body><a href="/">Home</a></body></html>`,
	})
	defer srv.Close()

	pages := Crawl(srv.URL+"/", 10)
	sort.Strings(pages)

	expected := []string{srv.URL + "/", srv.URL + "/about", srv.URL + "/contact"}
	sort.Strings(expected)

	if len(pages) != len(expected) {
		t.Errorf("expected %d pages, got %d: %v", len(expected), len(pages), pages)
	}
}

func TestCrawlIgnoresExternalLinks(t *testing.T) {
	srv := newTestServer(map[string]string{
		"/": `<html><body>
			<a href="/about">About</a>
			<a href="https://external.com/page">External</a>
		</body></html>`,
		"/about": `<html><body>About page</body></html>`,
	})
	defer srv.Close()

	pages := Crawl(srv.URL+"/", 10)
	for _, p := range pages {
		if p == "https://external.com/page" {
			t.Error("external link should not be included in sitemap")
		}
	}
	if len(pages) != 2 {
		t.Errorf("expected 2 pages, got %d: %v", len(pages), pages)
	}
}

func TestCrawlRespectsMaxDepth(t *testing.T) {
	srv := newTestServer(map[string]string{
		"/":  `<html><body><a href="/a">A</a></body></html>`,
		"/a": `<html><body><a href="/b">B</a></body></html>`,
		"/b": `<html><body><a href="/c">C</a></body></html>`,
		"/c": `<html><body>End</body></html>`,
	})
	defer srv.Close()

	// depth 2: should reach / → /a → /b, but not /c
	pages := Crawl(srv.URL+"/", 2)
	for _, p := range pages {
		if p == srv.URL+"/c" {
			t.Error("/c should not be visited at depth 2")
		}
	}
	if len(pages) != 3 {
		t.Errorf("expected 3 pages at depth 2, got %d: %v", len(pages), pages)
	}
}

func TestCrawlHandlesCycles(t *testing.T) {
	srv := newTestServer(map[string]string{
		"/a": `<html><body><a href="/b">B</a></body></html>`,
		"/b": `<html><body><a href="/a">A</a></body></html>`,
	})
	defer srv.Close()

	// should terminate, not loop forever
	pages := Crawl(srv.URL+"/a", 10)
	if len(pages) != 2 {
		t.Errorf("expected 2 pages, got %d: %v", len(pages), pages)
	}
}
