package utils

import (
	"regexp"
	"testing"
	"time"
)

func TestGetTimeUnix(t *testing.T) {
	startTime := time.Now()

	result := GetTimeUnix()

	endTime := time.Now()

	if result < startTime.Unix()-5 || result > endTime.Unix()+5 {
		t.Errorf("Expected Unix timestamp between %d and %d, but got: %d", startTime.Unix()-5, endTime.Unix()+5, result)
	}
}

func TestGetUUID(t *testing.T) {
	uuidPattern := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)

	result := GetUUID()

	if !uuidPattern.MatchString(result) {
		t.Errorf("Invalid UUID format: %s", result)
	}
}
