package twitter

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	baseURL string
	bearerToken string
}

func (c *Client) Retweeters(ctx context.Context, tweetID string) ([]string, error) {
	url := c.baseURL+"/2/tweets/"+tweetID+"/retweeted_by"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error duriung request creationto the %s; err: %w", url, err)
	}
	req.Header.Add("Authorization", "Bearer "+c.bearerToken)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error during sending request to the %s; err: %w", req.RequestURI, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("the response status code is not 2** from %s, we got: %d", req.RequestURI, resp.StatusCode)
	}

	var result struct {
		Data []struct {
			Username string `json:"username"`
		} `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("error during response decoding from %s; %w", req.RequestURI, err)
	}
	
	usernames := make([]string, 0, len(result.Data))
	for _, d := range result.Data {
		usernames = append(usernames, d.Username)
	}

	return usernames, nil
}