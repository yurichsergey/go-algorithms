package twitter

import (
	"context"
	"encoding/base64"
	"net/http"
	"net/url"
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
	return "", nil
}
