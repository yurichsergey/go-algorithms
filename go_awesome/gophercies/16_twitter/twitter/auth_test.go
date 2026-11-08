package twitter

import (
	"net/http/httptest"
	"http"
)

func TestAuth(t *testing.T) {

	server :=httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}))
	defer server.Close()
	
}

