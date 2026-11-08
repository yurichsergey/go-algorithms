package twitter

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuth(t *testing.T) {

	server :=httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"token_type":"bearer","access_token":"test-token"}`))
	}))
	defer server.Close()

	got, err := appOnlyToken(
		context.Background(),
		server.Client(),
		server.URL,
		"test-key",
		"test-secret",
	)
	if err != nil {
		t.Fatalf("appOnlyToken() error = %v", err)
	}
	if got != "test-token" {
		t.Errorf("appOnlyToken() got = %v, want %v", got, "test-token")
	}
}

