package utils

import (
	"fmt"
	"regexp"
)

func ValidateMimeTypes(mimeTypes []string) (bool, string) {
	mimeTypePattern := regexp.MustCompile(`^[a-zA-Z0-9-]+/[a-zA-Z0-9-+]+$`)

	if len(mimeTypes) == 0 {
		return true, ""
	}

	if len(mimeTypes) == 1 && mimeTypes[0] == "*" {
		return true, ""
	}

	for _, mimeType := range mimeTypes {
		if !mimeTypePattern.MatchString(mimeType) {
			return false, fmt.Sprintf("'%s' is invalid mime type", mimeType)
		}
	}

	return true, ""
}
