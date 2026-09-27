package server

import (
	"13_quiet_hn/cache"
	"13_quiet_hn/reader"
	"html/template"
	"net/http"
	"time"
)

func Handler(numStories int, tpl *template.Template) http.HandlerFunc {
	cached := cache.Cache{}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		//stories, err := reader.GetStories(numStories)
		stories, err := cached.GetStories(numStories)
		if err != nil {
			http.Error(w, "Failed to load top stories", http.StatusInternalServerError)
			return
		}

		data := reader.TemplateData{
			Stories: stories,
			Time:    time.Now().Sub(start),
		}
		err = tpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func InitServer(numStories int) {
	tpl := template.Must(template.ParseFiles("./server/view/index.gohtml"))

	http.HandleFunc("/", Handler(numStories, tpl))
}
