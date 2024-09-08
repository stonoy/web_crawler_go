package main

import (
	"net/url"
	"strings"
)

func normalizeUrl(rawUrl string) (string, error) {
	// parse the rawurl to url struct
	url, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	// get reqired info from the url struct
	fullPath := url.Host + url.Path

	fullPath = strings.ToLower(fullPath)
	fullPath = strings.TrimSuffix(fullPath, "/")

	// return url-> hostname and path
	return fullPath, nil
}
