package databaseentities

type Bucket struct {
	Id          string  `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
	IsPublic    bool    `json:"is_public" db:"is_public"`
	CreatedAt   int64   `json:"created_at" db:"created_at"`
	UpdatedAt   *int64  `json:"updated_at" db:"updated_at"`
}

func (Bucket) CollectionName() string {
	return "buckets"
}
