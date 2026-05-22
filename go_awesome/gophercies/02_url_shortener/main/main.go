package main

import (
	"fmt"
	"log"
	"net/http"
	"url_shortener"
)

func main() {
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	mux := defaultServeMux()
	mapHandler := url_shortener.MapHandler(pathsToUrls, mux)

	yamlHandler, err := url_shortener.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Starting the server on :8080")
	log.Fatal(http.ListenAndServe(":8080", yamlHandler))
}

func defaultServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintln(w, "Hello, world!"); err != nil {
			log.Fatal(err)
			return
		}
	})
	return mux
}
