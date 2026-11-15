package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/2/tweets/123/retweeted_by", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"data": []map[string]string{
				{"username": "alice"},
				{"username": "bob"},
			},
		})
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}