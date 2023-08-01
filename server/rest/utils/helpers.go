package utils

import (
	"github.com/google/uuid"
	"time"
)

func GetTimeUnix() int64 {
	return time.Now().Unix()
}

func GetUUID() string {
	return uuid.New().String()
}
