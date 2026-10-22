package middleware

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recovered:", err)
				log.Println(err)
				debug.Stack()
				http.Error(w, "Something went wrong", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)

	})
}
