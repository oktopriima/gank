/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 04/01/23, 15:40
 * Copyright (c) 2023
 */

package helpers

import (
	"fmt"
	"net/http"
	"strings"
)

func HeaderExtractor(key string, r *http.Request) (string, error) {
	header := r.Header.Get(key)
	if header == "" {
		return "", fmt.Errorf("empty header. %s", http.StatusText(http.StatusForbidden))
	}

	switch key {
	case "Authorization":
		authHeaderParts := strings.Fields(header)
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			return "", fmt.Errorf("invalid token. %s", http.StatusText(http.StatusForbidden))
		}
		return authHeaderParts[1], nil
	default:
		return header, nil
	}
}
