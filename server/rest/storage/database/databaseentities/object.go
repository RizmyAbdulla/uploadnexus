package databaseentities

const (
	UploadStatusPending   = "pending"
	UploadStatusCompleted = "completed"
	UploadStatusFailed    = "failed"
)

type Object struct {
	Id              string      `json:"id" db:"id"`
	BucketReference string      `json:"bucket_reference" db:"bucket_reference"`
	FileKey         string      `json:"file_key" db:"file_key"`
	FileName        string      `json:"file_name" db:"file_name"`
	FileType        string      `json:"file_type" db:"file_type"`
	FileSize        int64       `json:"file_size" db:"file_size"`
	UploadStatus    string      `json:"upload_status" db:"upload_status"`
	Metadata        interface{} `json:"metadata" db:"metadata"`
	CreatedAt       int64       `json:"created_at" db:"created_at"`
	UpdatedAt       *int64      `json:"updated_at" db:"updated_at"`
}

func (Object) CollectionName() string {
	return "objects"
}
