package models

type PreSignedUrl struct {
	BucketName string `json:"bucket_name"`
	ObjectName string `json:"object_name"`
	Url        string `json:"url"`
	HttpMethod string `json:"http_method"`
	Expiry     int64  `json:"expiry"`
}

type PreSignedUrlCreate struct {
	BucketName string `json:"bucket_name" validate:"required"`
	ObjectName string `json:"object_name" validate:"required"`
}
