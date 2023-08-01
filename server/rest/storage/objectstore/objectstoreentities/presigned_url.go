package objectstoreentities

import "time"

type PresignedUrl struct {
	BucketName string `json:"bucket_name"`
	ObjectName string `json:"object_name"`
	Expiry     int64  `json:"expiry"`
}

func (c *PresignedUrl) GetExpiry() time.Duration {
	return time.Duration(c.Expiry)
}
