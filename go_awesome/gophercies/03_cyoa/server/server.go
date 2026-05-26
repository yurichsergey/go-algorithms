package server

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/yurichsergey/go-algorithms/go_awesome/gophercies/03_cyoa/story"
)

const tpl = `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{.Title}}</title>
</head>
<body>
  <h1>{{.Title}}</h1>
  {{range .Story}}
	<p>{{.}}</p>
  {{end}}
  {{if .Options}}
    <ul>
    {{range .Options}}
      <li><a href="/{{.Arc}}">{{.Text}}</a></li>
    {{end}}
    </ul>
  {{end}}
</body>
</html>
`

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

	tmpl := template.Must(template.New("story").Parse(tpl))
	if err := tmpl.Execute(w, storyPart); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
