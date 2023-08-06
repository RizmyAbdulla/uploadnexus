package utils

import (
	"reflect"
	"strings"
	"testing"
)

func TestModelValidatorValidateModel(t *testing.T) {
	type TestModel struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}

	tests := []struct {
		name     string
		model    TestModel
		expected []map[string]interface{}
	}{
		{
			name:     "Valid model",
			model:    TestModel{Name: "John", Email: "john@example.com"},
			expected: nil,
		},
		{
			name:  "Invalid model",
			model: TestModel{Name: "", Email: "invalid_email"},
			expected: []map[string]interface{}{
				{
					"field":   "name",
					"message": "this field is required",
				},
				{
					"field":   "email",
					"message": "this field is email",
				},
			},
		},
	}

	modelValidator := NewModelValidator()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			messages, err := modelValidator.ValidateModel(tt.model)

			if err != nil {
				if tt.expected == nil {
					t.Errorf("Expected no error, but got: %v", err)
				}
			} else {
				if tt.expected != nil {
					t.Errorf("Expected error, but got nil")
				}
			}

			if !reflect.DeepEqual(messages, tt.expected) {
				t.Errorf("Expected messages: %v, but got: %v", tt.expected, messages)
			}

			if tt.expected != nil {
				for i, expectedMessage := range tt.expected {
					actualMessage := messages[i]
					checkMessageSubstring(t, strings.ToLower(expectedMessage["message"].(string)), strings.ToLower(actualMessage["message"].(string)))
				}
			}
		})
	}
}

func checkMessageSubstring(t *testing.T, expected, actual string) {
	if !strings.Contains(actual, expected) {
		t.Errorf("Expected message substring: %s, but got: %s", expected, actual)
	}
}
