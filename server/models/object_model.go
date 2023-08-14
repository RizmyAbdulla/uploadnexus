package models

const (
	UploadStatusPending   = "pending"
	UploadStatusCompleted = "completed"
	UploadStatusFailed    = "failed"
)

type ObjectCreate struct {
	Bucket   string      `json:"bucket" db:"bucket"`
	Name     string      `json:"name" db:"name"`
	Metadata interface{} `json:"metadata" db:"metadata"`
}

type PresignedObjectCreate struct {
	ObjectName string      `json:"object_name" validate:"required"`
	MimeType   string      `json:"mime_type" validate:"required"`
	Size       int64       `json:"size" validate:"required"`
	MetaData   interface{} `json:"metadata"`
}

type PreSignedObject struct {
	BucketName string `json:"bucket_name"`
	ObjectName string `json:"object_name"`
	Url        string `json:"url"`
	HttpMethod string `json:"http_method"`
	Expiry     int64  `json:"expiry"`
}

type PreSignedObjectResponse struct {
	Code            int             `json:"code"`
	Message         string          `json:"message"`
	PreSignedObject PreSignedObject `json:"pre_signed_object"`
}

type ObjectGeneralResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
