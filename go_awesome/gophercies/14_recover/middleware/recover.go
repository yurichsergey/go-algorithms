package middleware

import (
	"bytes"
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

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bw := &bufferedWriter{ResponseWriter: w}
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %v\n%s", err, debug.Stack())
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
