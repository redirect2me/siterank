package main

import (
	"strings"

	"golang.org/x/net/idna"
)

func purifyDomain(rawDomain string) (string, error) {
	if !isASCII(rawDomain) {
		// Punycode the domain
		punycode, err := idna.ToASCII(rawDomain)
		if err != nil {
			return "", err
		}
		rawDomain = punycode
	}

	domain := strings.ToLower(rawDomain)

	return domain, nil
}

func isASCII(s string) bool {
	for _, r := range s {
		if r > 127 {
			return false
		}
	}
	return true
}
