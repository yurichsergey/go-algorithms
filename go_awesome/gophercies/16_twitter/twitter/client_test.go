package twitter

import (
	"context"
	"net/http"
	"net/http/httptest"
	"regexp"
	"slices"
	"strings"
	"testing"
)

func TestRetweets(t *testing.T) {
	tests := []struct {
		name            string
		status          int
		body            string
		want            []string
		wantErr         bool
		wantErrContains string
	}{
		{
			name:    "test successful retweeters retrieval",
			status:  http.StatusOK,
			body:    `{"data":[{"username":"user1"},{"username":"user2"}]}`,
			want:    []string{"user1", "user2"},
			wantErr: false,
		},
		{
			name:    "empty list",
			status:  http.StatusOK,
			body:    `{}`,
			want:    []string{},
			wantErr: false,
		},
		{
			name:    "malformed json",
			status:  http.StatusOK,
			body:    `{"data": [`,
			wantErr: true,
		},
		{
			name:            "unauthorized",
			status:          http.StatusUnauthorized,
			body:            `{"title":"Unauthorized"}`,
			wantErr:         true,
			wantErrContains: `"title":"Unauthorized"`,
		},
		{
			name:            "too many requests",
			status:          http.StatusTooManyRequests,
			body:            `{"title":"Too Many Requests"}`,
			wantErr:         true,
			wantErrContains: `"title":"Too Many Requests"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet {
					t.Errorf("Expected GET request, got %s", r.Method)
				}

				pathPattern := regexp.MustCompile(`^/2/tweets/\d+/retweeted_by$`)
				if !pathPattern.MatchString(r.URL.Path) {
					t.Errorf("Expected request to /2/tweets/{tweetID}/retweeted_by, got %s", r.URL.Path)
				}

				if strings.HasPrefix(r.Header.Get("Authorization"), "Bearer ") == false {
					t.Errorf("Expected Authorization header to start with Bearer, got %s", r.Header.Get("Authorization"))
				}

				authHeader := r.Header.Get("Authorization")
				token := strings.TrimPrefix(authHeader, "Bearer ")

				if token != "test-token" {
					t.Errorf("Expected Bearer token should be 'test-token', got %s", token)
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.status)
				_, _ = w.Write([]byte(tt.body))
			}))
			defer server.Close()

			client := NewClient(server.Client(), server.URL, "test-token")
			got, err := client.Retweeters(context.Background(), "123456")
			if tt.wantErr == false && err != nil {
				t.Fatalf("client.Retweeters() error = %v", err)
			}
			if tt.wantErr == true && err == nil {
				t.Fatalf("client.Retweeters() expected error but got none")
			}
			if tt.wantErrContains != "" && !strings.Contains(err.Error(), tt.wantErrContains) {
				t.Errorf("client.Retweeters() error = %q, want it to contain %q", err, tt.wantErrContains)
			}
			if tt.wantErr {
				return
			}
			if slices.Equal(got, tt.want) == false {
				t.Errorf("client.Retweeters() got = %v, want %v", got, tt.want)
			}

		})
	}
}
