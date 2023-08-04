package entities

type Bucket struct {
	Id               string    `json:"id" db:"id"`
	Name             string    `json:"name" db:"name"`
	Description      *string   `json:"description" db:"description"`
	AllowedMimeTypes *[]string `json:"allowed_mime_types" db:"allowed_mime_types"`
	FileSizeLimit    *int64    `json:"file_size_limit" db:"file_size_limit"`
	IsPublic         bool      `json:"is_public" db:"is_public"`
	CreatedAt        int64     `json:"created_at" db:"created_at"`
	UpdatedAt        *int64    `json:"updated_at" db:"updated_at"`
}

func (Bucket) CollectionName() string {
	return "buckets"
}

type BucketDto struct {
	Name             string    `json:"name" validate:"required"`
	Description      *string   `json:"description,omitempty"`
	AllowedMimeTypes *[]string `json:"allowed_mime_types,omitempty"`
	FileSizeLimit    *int64    `json:"file_size_limit,omitempty"`
	IsPublic         *bool     `json:"is_public,omitempty"`
}
