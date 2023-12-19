package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(header http.Header) (string, error) {
	val := header.Get("Authorization")
	if val == "" {
		return "", errors.New("no authorization found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed authorization header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed authorization first part")
	}
	return vals[1], nil
}