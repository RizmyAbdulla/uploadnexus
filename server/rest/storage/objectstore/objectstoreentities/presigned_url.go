package objectstoreentities

import "time"

type CreatedPresignedUrl struct {
	BucketName string `json:"bucket_name"`
	ObjectName string `json:"object_name"`
	Expiry     int64  `json:"expiry"`
}

func (c *CreatedPresignedUrl) GetExpiry() time.Duration {
	return time.Duration(c.Expiry)
}
