package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API key from the Authorization header.
// Example:
// Authorization: ApiKey {insert API key here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication header found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid authentication header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("invalid authentication method")
	}

	return vals[1], nil
}
