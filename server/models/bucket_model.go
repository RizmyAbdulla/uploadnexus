package models

type Bucket struct {
	Id                string    `json:"id"`
	Name              string    `json:"name"`
	Description       *string   `json:"description"`
	AllowedMimeTypes  *[]string `json:"allowed_mime_types"`
	AllowedObjectSize *int64    `json:"allowed_object_size"`
	IsPublic          bool      `json:"is_public"`
	CreatedAt         int64     `json:"created_at"`
	UpdatedAt         *int64    `json:"updated_at"`
}

type BucketResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Bucket  Bucket `json:"bucket"`
}

type BucketListResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Bucket  []Bucket `json:"buckets"`
}

type BucketGeneralResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BucketCreate struct {
	Name              string   `json:"name" validate:"required,min=3,max=36"`
	Description       *string  `json:"description,omitempty"`
	AllowedMimeTypes  []string `json:"allowed_mime_types,omitempty"`
	AllowedObjectSize int64    `json:"allowed_object_size,omitempty"`
	IsPublic          bool     `json:"is_public,omitempty"`
}

type BucketUpdate struct {
	Name              string   `json:"name,omitempty" validate:"min=3,max=36"`
	Description       *string  `json:"description,omitempty"`
	AllowedMimeTypes  []string `json:"allowed_mime_types,omitempty"`
	AllowedObjectSize int64    `json:"allowed_object_size,omitempty"`
	IsPublic          bool     `json:"is_public,omitempty"`
}
