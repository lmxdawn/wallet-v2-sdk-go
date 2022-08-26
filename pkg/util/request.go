package util

import (
	"net/http"
	"strings"
)

func GetUrl(r *http.Request, path string) string {
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}

	return strings.Join([]string{scheme, r.Host, path}, "")
}
