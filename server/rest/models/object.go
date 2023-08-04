package entities

const (
	UploadStatusPending   = "pending"
	UploadStatusCompleted = "completed"
	UploadStatusFailed    = "failed"
)

type Object struct {
	Id           string      `json:"id" db:"id"`
	Bucket       string      `json:"bucket" db:"bucket"`
	Name         string      `json:"name" db:"name"`
	MimeType     string      `json:"mime_type" db:"mime_type"`
	Size         int64       `json:"size" db:"size"`
	UploadStatus string      `json:"upload_status" db:"upload_status"`
	Metadata     interface{} `json:"metadata" db:"metadata"`
	CreatedAt    int64       `json:"created_at" db:"created_at"`
	UpdatedAt    *int64      `json:"updated_at" db:"updated_at"`
}

func (Object) CollectionName() string {
	return "objects"
}
