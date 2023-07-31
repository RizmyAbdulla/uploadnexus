package databaseentities

type Application struct {
	Id          string  `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
	CreatedAt   int64   `json:"created_at" db:"created_at"`
	UpdatedAt   *int64  `json:"updated_at" db:"updated_at"`
}

func (Application) CollectionName() string {
	return "applications"
}
