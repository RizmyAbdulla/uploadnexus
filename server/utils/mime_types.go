package utils

import (
	"fmt"
	"mime"
	"regexp"
	"strings"
)

func GetMimeTypeByExtension(extension string) (string, error) {
	mimeType := mime.TypeByExtension(extension)
	if mimeType == "" {
		return "", fmt.Errorf("unknown extension: %s", extension)
	}
	return mimeType, nil
}

func ValidateMimeTypes(mimeTypes []string) (bool, error) {
	mimeTypePattern := regexp.MustCompile(`^[a-zA-Z0-9-]+/[a-zA-Z0-9-+]+$`)

	if len(mimeTypes) == 0 {
		return true, nil
	}

	for _, mimeType := range mimeTypes {
		if mimeType == "*" || strings.HasSuffix(mimeType, "/*") {
			continue // Wildcard allows any MIME type
		}
		if !mimeTypePattern.MatchString(mimeType) {
			return false, fmt.Errorf(`'%s' is an invalid mime type`, mimeType)
		}
	}

	return true, nil
}

func ValidateAllowedMimeTypes(mimeType string, allowedMimeTypes []string) (bool, error) {
	if len(allowedMimeTypes) == 0 {
		return true, nil
	}

	if len(allowedMimeTypes) == 1 && allowedMimeTypes[0] == "*" {
		return true, nil
	}

	mimeType = strings.ToLower(mimeType)

	for _, allowed := range allowedMimeTypes {
		allowed = strings.ToLower(allowed)
		if allowed == "*" || strings.HasSuffix(allowed, "/*") {
			prefix := strings.TrimSuffix(allowed, "/*")
			if len(prefix) == 0 || strings.HasPrefix(mimeType, prefix) {
				return true, nil
			}
		} else if mimeType == allowed {
			return true, nil
		}
	}

	return false, fmt.Errorf(`expected mime types '%s' but got '%s'`, strings.Join(allowedMimeTypes, ","), mimeType)
}
