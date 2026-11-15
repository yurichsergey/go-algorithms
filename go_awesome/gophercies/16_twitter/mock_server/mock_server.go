package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Mock Twitter API server is running"))
		if err != nil {
			log.Printf("writing hello response %v", err)
		}
	})
	http.HandleFunc("/2/tweets/123/retweeted_by", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"data": []map[string]string{
				{"username": "alice"},
				{"username": "bob"},
			},
		})
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}