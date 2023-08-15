package utils

import (
	"fmt"
	"github.com/ArkamFahry/uploadnexus/server/constants"
	"github.com/google/uuid"
	"time"
)

func GetTimeUnix() int64 {
	return time.Now().Unix()
}

func GetUUID() string {
	return uuid.New().String()
}

func ValidatePaging(page int, pageSize int) (int, int, error) {
	if page < 0 {
		return 0, 0, fmt.Errorf("page cannot be less than 0")
	} else if page == 0 {
		page = constants.DefaultPage
	}

	if pageSize < 0 {
		return 0, 0, fmt.Errorf("page limit cannot be less than 0")
	} else if pageSize == 0 {
		pageSize = constants.DefaultPageLimit
	}

	return page, pageSize, nil
}

func TransformPaging(page int, pageSize int) (int, int, error) {
	limit := pageSize
	offset := page * limit

	return limit, offset, nil
}
