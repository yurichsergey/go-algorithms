package twitter

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
)

func TestRetweets(t *testing.T) {
	tests := []struct{
		name string
		status int
		body string
		want []string
		wantErr bool
	}{
		{
			name: "test successful retweeters retrieval",
			status: http.StatusOK,
			body: `{"data":[{"username":"user1"},{"username":"user2"}]}`,
			want: []string{"user1", "user2"},
			wantErr: false,
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