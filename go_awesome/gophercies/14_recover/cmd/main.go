package main

import (
	"14_recover/middleware"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/panic/", panicDemo)
	mux.HandleFunc("/panic-after/", panicAfterDemo)
	mux.HandleFunc("/stream/", streamDemo)
	mux.HandleFunc("/panic-stream/", panicStreamDemo)
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

func streamDemo(w http.ResponseWriter, r *http.Request) {
	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming not supported", http.StatusInternalServerError)
		return
	}
	for i := 1; i <= 3; i++ {
		_, err := fmt.Fprintf(w, "chunk %d\n", i)
		if err != nil {
			return
		}
		f.Flush()
		time.Sleep(time.Second)
	}
}

func panicStreamDemo(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "chunk 1")
	if err != nil {
		return
	}
	w.(http.Flusher).Flush()
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
  <li><a href="/stream/">/stream/</a> (streams 3 chunks, no panic)</li>
  <li><a href="/panic-stream/">/panic-stream/</a> (flushes, then panics)</li>
</ul>`)
	if err != nil {
		return
	}
}
