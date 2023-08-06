package models

type PreSignedObject struct {
	BucketName string `json:"bucket_name"`
	ObjectName string `json:"object_name"`
	Url        string `json:"url"`
	HttpMethod string `json:"http_method"`
	Expiry     int64  `json:"expiry"`
}
