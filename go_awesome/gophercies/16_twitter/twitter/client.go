package twitter

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Client struct {
	httpClient  *http.Client
	baseURL     string
	bearerToken string
}

func NewClient(httpClient *http.Client, baseURL string, bearerToken string) *Client {
	return &Client{httpClient: httpClient, baseURL: baseURL, bearerToken: bearerToken}
}

func (c *Client) Retweeters(ctx context.Context, tweetID string) ([]string, error) {
	URLPath := c.baseURL + "/2/tweets/" + tweetID + "/retweeted_by"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URLPath, nil)
	if err != nil {
		return nil, fmt.Errorf("error during request creation for %s; err: %w", URLPath, err)
	}
	req.Header.Add("Authorization", "Bearer "+c.bearerToken)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error during sending request to the %s; err: %w", req.URL, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("error during closing response body from %s; err: %v", req.URL, err)
		}
	}(resp.Body)

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf(
			"the response status code is not 2** from %s, we got: %d; body: %s",
			req.URL,
			resp.StatusCode,
			body,
		)
	}

	var result struct {
		Data []struct {
			Username string `json:"username"`
		} `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("error during response decoding from %s; %w", req.URL, err)
	}

	usernames := make([]string, 0, len(result.Data))
	for _, d := range result.Data {
		usernames = append(usernames, d.Username)
	}

	return usernames, nil
}
