package twitter

import (
	"context"
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuth(t *testing.T) {
	tests := []struct {
		name      string
		status    int
		body      string
		wantToken string
		wantErr   bool
	}{
		{
			name:      "successful token retrieval",
			status:    http.StatusOK,
			body:      `{"token_type":"bearer","access_token":"test-token"}`,
			wantToken: "test-token",
			wantErr:   false,
		},
		{
			name:      "failed token retrieval (empty token)",
			status:    http.StatusOK,
			body:      `{"token_type":"bearer","access_token":""}`,
			wantToken: "",
			wantErr:   true,
		},
		{
			name:      "failed token retrieval (broken json)",
			status:    http.StatusOK,
			body:      `{"token_type":"bearer","access_token":"test-token"`,
			wantToken: "",
			wantErr:   true,
		},
		{
			name:      "unexpected status code",
			status:    http.StatusBadRequest,
			body:      `{"error":"invalid_request"}`,
			wantToken: "",
			wantErr:   true,
		},
		{
			name:      "unexpected token type",
			status:    http.StatusOK,
			body:      `{"token_type":"invalid","access_token":"test-token"}`,
			wantToken: "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

				r.ParseForm()
				if r.Form.Get("grant_type") != "client_credentials" {
					t.Errorf("Expected grant_type to be client_credentials, got %s", r.Form.Get("grant_type"))
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.status)
				_, _ = w.Write([]byte(tt.body))
			}))
			defer server.Close()

			got, err := appOnlyToken(
				context.Background(),
				server.Client(),
				server.URL+"/oauth2/token",
				"test-key",
				"test-secret",
			)
			if tt.wantErr == false && err != nil {
				t.Fatalf("appOnlyToken() error = %v", err)
			}
			if tt.wantErr == true && err == nil {
				t.Fatalf("appOnlyToken() expected error but got none")
			}
			if got != tt.wantToken {
				t.Errorf("appOnlyToken() got = %v, want %v", got, tt.wantToken)
			}
		})
	}
}
