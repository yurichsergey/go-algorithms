package main

import (
	"14_recover/middleware"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/panic/", panicDemo)
	mux.HandleFunc("/panic-after/", panicAfterDemo)
	mux.HandleFunc("/", hello)
	dev := os.Getenv("APP_ENV") == "development"
	log.Fatal(http.ListenAndServe(":3000", middleware.Recover(mux, dev)))
}

func panicDemo(w http.ResponseWriter, r *http.Request) {
	funcThatPanics()
}

func panicAfterDemo(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "<h1>Hello!</h1>")
	if err != nil {
		return
	}
	funcThatPanics()
}

func funcThatPanics() {
	panic("Oh no!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, `<h1>Hello!</h1>
<ul>
  <li><a href="/">/</a></li>
  <li><a href="/panic/">/panic/</a></li>
  <li><a href="/panic-after/">/panic-after/</a></li>
</ul>`)
	if err != nil {
		return
	}
}
