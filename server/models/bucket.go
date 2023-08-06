package models

const BucketCollection = "buckets"

// Bucket model for buckets collection in database
type Bucket struct {
	Id                string   `json:"id" db:"id"`
	Name              string   `json:"name" db:"name"`
	Description       *string  `json:"description" db:"description"`
	AllowedMimeTypes  []string `json:"allowed_mime_types" db:"allowed_mime_types"`
	AllowedObjectSize int64    `json:"allowed_object_size" db:"allowed_object_size"`
	IsPublic          bool     `json:"is_public" db:"is_public"`
	CreatedAt         int64    `json:"created_at" db:"created_at"`
	UpdatedAt         int64    `json:"updated_at" db:"updated_at"`
}

type BucketCreate struct {
	Name              string   `json:"name" validate:"required"`
	Description       *string  `json:"description,omitempty"`
	AllowedMimeTypes  []string `json:"allowed_mime_types,omitempty"`
	AllowedObjectSize int64    `json:"allowed_object_size,omitempty"`
	IsPublic          bool     `json:"is_public,omitempty"`
}
