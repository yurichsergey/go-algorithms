package middleware

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

type bufferedWriter struct {
	http.ResponseWriter // embedded real writer
	status              int
	body                bytes.Buffer
}

func (bw *bufferedWriter) WriteHeader(status int) {
	bw.status = status
}

func (bw *bufferedWriter) Write(b []byte) (int, error) {
	return bw.body.Write(b)
}

func Recover(next http.Handler, dev bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bw := &bufferedWriter{ResponseWriter: w}
		defer func() {
			if err := recover(); err != nil {
				stack := debug.Stack()
				log.Printf("panic: %v\n%s", err, stack)
				if dev {
					http.Error(
						w,
						fmt.Sprintf("Something went wrong:\n%v\n\n%s", err, stack),
						http.StatusInternalServerError,
					)
					return
				}
				http.Error(w, "Something went wrong", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(bw, r)

		if bw.status != 0 {
			w.WriteHeader(bw.status)
		}
		_, err := w.Write(bw.body.Bytes())
		if err != nil {
			return
		}
	})
}
