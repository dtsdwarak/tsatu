package util

import (
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

// Tweaked from: https://stackoverflow.com/a/42718113/1170664
func GetHTTPContent(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return []byte{}, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}

func UnshortenURL(url string) (string, error) {

	sanitizedURL, err := SanitizeURL(url)
	if err != nil {
		return "", fmt.Errorf("Sanitize URL Error: %v", err)
	}

	resp, err := http.Get(sanitizedURL)
	if err != nil {
		return "", fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	unshortenedURL := resp.Request.URL.String()
	if unshortenedURL != sanitizedURL {

		unshortenedURL, err = UnshortenURL(unshortenedURL)
		if err != nil {
			return "", errors.Wrapf(err, "Unable to unshorten URL")
		}
		return unshortenedURL, nil
	}

	return unshortenedURL, nil
}
