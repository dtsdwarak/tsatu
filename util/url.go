package util

import (
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

func SanitizeURL(inputURL string) (string, error) {
	strippedURL := strings.Split(strings.Split(strings.Split(strings.Split(strings.Split(strings.TrimPrefix(inputURL, "https://www.google.com/url?q="), "ref=")[0], "ref%3D")[0], "&amp")[0], "?")[0], "%3F")[0]
	parsedURL, err := url.Parse(strippedURL)
	if err != nil {
		return "", errors.Wrapf(err, "Error parsing URL - %s", inputURL)
	}
	return parsedURL.Scheme + "://" + parsedURL.Host + parsedURL.Path, nil
}
