package twitter

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	httpClient  *http.Client
	baseURL     string
	bearerToken string
}

type response struct {
	Data []struct {
		Username string `json:"username"`
	} `json:"data"`
	Meta struct {
		NextToken string `json:"next_token"`
	} `json:"meta"`
}

func NewClient(httpClient *http.Client, baseURL string, bearerToken string) *Client {
	return &Client{httpClient: httpClient, baseURL: baseURL, bearerToken: bearerToken}
}

func (c *Client) Retweeters(ctx context.Context, tweetID string) ([]string, error) {
	token := ""
	usernames := make([]string, 0)
	for {
		result, err := c.fetchPage(ctx, tweetID, token)
		if err != nil {
			return []string{}, err
		}

		for _, d := range result.Data {
			usernames = append(usernames, d.Username)
		}

		token = result.Meta.NextToken
		if token == "" {
			break
		}
	}

	return usernames, nil
}

func (c *Client) fetchPage(ctx context.Context, tweetID string, token string) (response, error) {
	endpoint := c.baseURL + "/2/tweets/" + tweetID + "/retweeted_by"
	q := url.Values{}
	q.Set("max_results", "100")
	if token != "" {
		q.Set("pagination_token", token)
	}
	endpoint += "?" + q.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return response{}, fmt.Errorf("error during request creation for %s; err: %w", endpoint, err)
	}
	req.Header.Add("Authorization", "Bearer "+c.bearerToken)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return response{}, fmt.Errorf("error during sending request to the %s; err: %w", req.URL, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("error during closing response body from %s; err: %v", req.URL, err)
		}
	}(resp.Body)

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		body, _ := io.ReadAll(resp.Body)
		return response{}, fmt.Errorf(
			"the response status code is not 2** from %s, we got: %d; body: %s",
			req.URL,
			resp.StatusCode,
			body,
		)
	}

	var result response
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return response{}, fmt.Errorf("error during response decoding from %s; %w", req.URL, err)
	}
	return result, nil
}
