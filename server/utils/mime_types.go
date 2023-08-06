package utils

import (
	"fmt"
	"mime"
	"regexp"
	"strings"
)

func GetMimeType(objectName string) (bool, string) {
	mimeType := mime.TypeByExtension(objectName)

	if mimeType == "" {
		return false, ""
	}

	return true, mimeType
}

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
			return false, fmt.Sprintf(`'%s' is invalid mime type`, mimeType)
		}
	}

	return true, ""
}

func ValidateAllowedMimeTypes(mimeType string, allowedMimeTypes []string) (bool, string) {
	if len(allowedMimeTypes) == 0 {
		return true, ""
	}

	if len(allowedMimeTypes) == 1 && allowedMimeTypes[0] == "*" {
		return true, ""
	}

	mimeType = strings.ToLower(mimeType)

	for _, allowed := range allowedMimeTypes {
		if mimeType == strings.ToLower(allowed) {
			return true, ""
		}
	}

	return false, fmt.Sprintf(`expected mime types '%s' but got '%s'`, strings.Join(allowedMimeTypes, ","), mimeType)
}
