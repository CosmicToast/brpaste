package http

import "strings"

var (
	browsers = []string{
		"Firefox/",
		"Chrome/",
		"Safari/",
		"OPR/",
		"Edge/",
		"Trident/",
	}
)

func isBrowser(ua string) bool {
	for _, el := range browsers {
		if strings.Contains(ua, el) {
			return true
		}
	}
	return false
}
