package retries

import (
	"context"
	"time"
)

const defaultMaxRetries = 3

func Retry(ctx context.Context, fn func() error, customMaxRetries ...int) error {
	var maxRetries int
	if len(customMaxRetries) > 0 {
		maxRetries = customMaxRetries[0]
	} else {
		maxRetries = defaultMaxRetries
	}
	retryInterval := time.Second

	for retry := 0; retry <= maxRetries; retry++ {
		if err := fn(); err != nil {
			if retry == maxRetries {
				return err
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(retryInterval * time.Duration(retry)):
			}
			continue
		}
		return nil
	}
	return nil
}
