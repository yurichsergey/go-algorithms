package main

import (
	"net/http"

	"github.com/yurichsergey/go-algorithms/go_awesome/gophercies/03_cyoa/server"
	"github.com/yurichsergey/go-algorithms/go_awesome/gophercies/03_cyoa/story"
)

func main() {

	storyStruct := story.FromFilePath("data/gopher.json")

	handler := server.NewRequestHandler(storyStruct)
	errHandler := http.ListenAndServe(":8080", handler)
	if errHandler != nil {
		return
	}

}
