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
