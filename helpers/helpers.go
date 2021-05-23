package helpers

import "strings"

func IsImage(name string) bool {
	n := strings.ToLower(name)
	if strings.HasSuffix(n, ".jpg") {
		return true
	}

	if strings.HasSuffix(n, ".png") {
		return true
	}

	if strings.HasSuffix(n, ".gif") {
		return true
	}

	return false
}
