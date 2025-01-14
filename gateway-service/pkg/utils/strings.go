package utils

import (
	"net/url"
	"os"
	"strings"
)

func ChangeDomainFromMinio(original *string) *string {
	if original == nil || *original == "" {
		return nil
	}

	parsedURL, err := url.Parse(*original)
	if err != nil {
		return nil
	}

	cdnDomain := os.Getenv("CDN_PUBLIC_HOST")
	if cdnDomain == "" {
		return nil
	}

	cdnDomain = strings.TrimPrefix(cdnDomain, "http://")
	cdnDomain = strings.TrimPrefix(cdnDomain, "https://")

	parsedURL.Host = cdnDomain

	parsedURL.Scheme = "https"
	result := parsedURL.String()
	return &result
}
