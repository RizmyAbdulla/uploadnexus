package utils

import (
	"strings"
	"testing"
)

func TestGetMimeTypeByExtension(t *testing.T) {
	tests := []struct {
		extension string
		expected  string
		errMsg    string
	}{
		{".jpg", "image/jpeg", ""},
		{".pdf", "application/pdf", ""},
		{".txt", "text/plain", ""},
		{".invalid", "", "unknown extension: .invalid"},
		{".png", "image/png", ""},
	}

	for _, test := range tests {
		mimeType, err := GetMimeTypeByExtension(test.extension)
		if err != nil {
			if test.errMsg == "" {
				t.Errorf("For extension: %s, unexpected error: %v", test.extension, err)
			} else if err.Error() != test.errMsg {
				t.Errorf("For extension: %s, expected error message: %s, got: %v", test.extension, test.errMsg, err)
			}
		} else if !strings.HasPrefix(mimeType, test.expected) {
			t.Errorf("For extension: %s, expected: %s, got: %s", test.extension, test.expected, mimeType)
		}
	}
}

func TestValidateMimeTypes(t *testing.T) {
	tests := []struct {
		mimeTypes []string
		expected  bool
		errMsg    string
	}{
		{[]string{"image/jpeg", "application/pdf"}, true, ""},
		{[]string{"audio/mpeg", "video/mp4"}, true, ""},
		{[]string{"image/png", "invalid-mime-type"}, false, "'invalid-mime-type' is an invalid mime type"},
		{[]string{"application/*"}, true, ""},
		{[]string{}, true, ""},
		{[]string{"application/pdf", "application/msword", "text/plain"}, true, ""},
		{[]string{"audio/mpeg", "video/*"}, true, ""},
		{[]string{"invalid-mime-type", "audio/mpeg"}, false, "'invalid-mime-type' is an invalid mime type"},
		{[]string{"*"}, true, ""},
	}

	for _, test := range tests {
		valid, err := ValidateMimeTypes(test.mimeTypes)
		if valid != test.expected {
			t.Errorf("For mimeTypes: %v, expected: %v, got: %v", test.mimeTypes, test.expected, valid)
		}
		if test.errMsg != "" && err.Error() != test.errMsg {
			t.Errorf("For mimeTypes: %v, expected error message: %v, got: %v", test.mimeTypes, test.errMsg, err.Error())
		}
	}
}

func TestValidateAllowedMimeTypes(t *testing.T) {
	tests := []struct {
		mimeType         string
		allowedMimeTypes []string
		expected         bool
		errMsg           string
	}{
		{"image/jpeg", []string{"image/jpeg", "application/pdf"}, true, ""},
		{"video/mp4", []string{"audio/mpeg", "video/mp4"}, true, ""},
		{"image/png", []string{"image/jpeg", "application/pdf"}, false, "expected mime types 'image/jpeg,application/pdf' but got 'image/png'"},
		{"application/pdf", []string{"*"}, true, ""},
		{"text/plain", []string{}, true, ""},
		{"application/msword", []string{"application/pdf", "application/msword", "text/plain"}, true, ""},
		{"audio/mpeg", []string{"audio/*"}, true, ""},
		{"invalid-mime-type", []string{"audio/mpeg"}, false, "expected mime types 'audio/mpeg' but got 'invalid-mime-type'"},
	}

	for _, test := range tests {
		valid, err := ValidateAllowedMimeTypes(test.mimeType, test.allowedMimeTypes)
		if valid != test.expected {
			t.Errorf("For mimeType: %v, allowedMimeTypes: %v, expected: %v, got: %v", test.mimeType, test.allowedMimeTypes, test.expected, valid)
		}
		if test.errMsg != "" && err.Error() != test.errMsg {
			t.Errorf("For mimeType: %v, allowedMimeTypes: %v, expected error message: %v, got: %v", test.mimeType, test.allowedMimeTypes, test.errMsg, err.Error())
		}
	}
}
