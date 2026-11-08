package twitter

import (
	"encoding/base64"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuth(t *testing.T) {

	server :=httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if strings.HasPrefix(r.URL.Path, "/oauth2/token") == false {
			t.Errorf("Expected request to /oauth2/token, got %s", r.URL.Path)
		}
		if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded;charset=UTF-8" {
			t.Errorf("Expected Content-Type header to be set, got %s", r.Header.Get("Content-Type"))
		}
		if strings.HasPrefix(r.Header.Get("Authorization"), "Basic ") == false {
			t.Errorf("Expected Authorization header to start with Basic, got %s", r.Header.Get("Authorization"))
		}
		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Basic ")
		base64Decoded, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			t.Errorf("Failed to decode base64 token: %v", err)
		}
		if string(base64Decoded) != "test-key:test-secret" {
			t.Errorf("Expected base64 decoded token to be 'test-key:test-secret', got %s", string(base64Decoded))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"token_type":"bearer","access_token":"test-token"}`))
	}))
	defer server.Close()

	got, err := appOnlyToken(
		context.Background(),
		server.Client(),
		server.URL + "/oauth2/token",
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

