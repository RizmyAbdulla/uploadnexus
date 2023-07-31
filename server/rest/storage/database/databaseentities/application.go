package databaseentities

type Application struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	CreatedAt   int64   `json:"created_at"`
	UpdatedAt   *int64  `json:"updated_at"`
}

func (Application) CollectionName() string {
	return "applications"
}
