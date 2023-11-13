package utils

import "net/url"

func IsValidURL(urlStr string) bool {
	_, err := url.ParseRequestURI(urlStr)

	if err != nil {
		return false
	}

	u, err := url.Parse(urlStr)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
