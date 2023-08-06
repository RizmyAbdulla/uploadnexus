package models

type BucketCreate struct {
	Name              string   `json:"name" validate:"required"`
	Description       *string  `json:"description,omitempty"`
	AllowedMimeTypes  []string `json:"allowed_mime_types,omitempty"`
	AllowedObjectSize int64    `json:"allowed_object_size,omitempty"`
	IsPublic          bool     `json:"is_public,omitempty"`
}
