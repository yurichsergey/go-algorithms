package twitter

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func appOnlyToken(
	ctx context.Context,
	client *http.Client,
	tokenURL,
	key,
	secret string,
) (string, error) {

	form := url.Values{}
	form.Set("grant_type", "client_credentials")

	credentials := url.QueryEscape(key) + ":" + url.QueryEscape(secret)
	encoded := base64.StdEncoding.EncodeToString([]byte(credentials))

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		tokenURL,
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Basic "+encoded)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("error during closing response body from %s; err: %v", req.URL, err)
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		return "", errors.New("unexpected status code: " + resp.Status)
	}

	var token struct {
		TokenType   string `json:"token_type"`
		AccessToken string `json:"access_token"`
	}

	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return "", err
	}

	if token.TokenType != "bearer" {
		return "", errors.New("unexpected token type: " + token.TokenType)
	}

	if token.AccessToken == "" {
		return "", errors.New("access token is empty")
	}

	return token.AccessToken, nil
}
