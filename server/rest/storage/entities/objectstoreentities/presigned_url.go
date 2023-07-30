package objectstoreentities

import "time"

type PresignedUrl struct {
	Url    string `json:"url"`
	Expiry int64  `json:"expiry"`
}

type CreatedPresignedUrl struct {
	BucketName string `json:"bucket_name"`
	ObjectName string `json:"object_name"`
	Expiry     int64  `json:"expiry"`
}

func (c *CreatedPresignedUrl) GetExpiry() time.Time {
	return time.Unix(c.Expiry, 0)
}
