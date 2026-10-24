package middleware

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime/debug"
)

type bufferedWriter struct {
	http.ResponseWriter // embedded real writer
	status              int
	body                bytes.Buffer
	committed           bool
}

func (bw *bufferedWriter) WriteHeader(status int) {
	if bw.committed {
		return
	}
	bw.status = status
}

func (bw *bufferedWriter) Write(b []byte) (int, error) {
	if bw.committed {
		return bw.ResponseWriter.Write(b)
	}
	return bw.body.Write(b)
}

func (bw *bufferedWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := bw.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, http.ErrNotSupported
	}
	conn, rw, err := h.Hijack()
	if err == nil {
		bw.committed = true
	}
	return conn, rw, err
}

func (bw *bufferedWriter) Flush() {
	bw.commit()
	if f, ok := bw.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
}

func (bw *bufferedWriter) commit() {
	if bw.committed {
		return
	}
	if bw.status != 0 {
		bw.ResponseWriter.WriteHeader(bw.status)
	}
	_, _ = bw.ResponseWriter.Write(bw.body.Bytes())
	bw.body.Reset()
	bw.committed = true
}

var _ http.Flusher = (*bufferedWriter)(nil)
var _ http.Hijacker = (*bufferedWriter)(nil)

func Recover(next http.Handler, dev bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bw := &bufferedWriter{ResponseWriter: w}

		defer func() {
			if err := recover(); err != nil {
				stack := debug.Stack()
				log.Printf("panic: %v\n%s", err, stack)
				if bw.committed {
					return
				}
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

		bw.commit()
	})
}
