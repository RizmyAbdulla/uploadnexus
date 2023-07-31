package utils

import (
	"github.com/google/uuid"
	"time"
)

func GetTime() int64 {
	return time.Now().Unix()
}

func GetUuid() string {
	return uuid.NewString()
}
