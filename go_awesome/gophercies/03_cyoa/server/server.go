package server

import (
	"net/http"
	"strings"

	"github.com/yurichsergey/go-algorithms/go_awesome/gophercies/03_cyoa/story"
)

func NewRequestHandler(s story.Story) http.Handler {
	return RequestHandler{Story: s}
}

type RequestHandler struct {
	Story story.Story
}

func (h RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	storyPart, ok := h.Story[strings.Trim(r.URL.Path, "/")]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		notFoundHTML := `<h1>Story not found</h1><p>Go to the start: <a href="/intro">Intro</a></p>`
		if _, err := w.Write([]byte(notFoundHTML)); err != nil {
			return
		}
		return
	}

	storyHTML := `<h1>` + storyPart.Title + `</h1>`
	for _, paragraph := range storyPart.Story {
		storyHTML += `<p>` + paragraph + `</p>`
	}

	if len(storyPart.Options) > 0 {
		storyHTML += `<p>Continue<ul>`
		for _, choice := range storyPart.Options {
			storyHTML += `<li><a href="` + choice.Arc + `">` + choice.Text + `</a></li>`
		}
		storyHTML += `</ul></p>`
	}

	if _, err := w.Write([]byte(storyHTML)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
