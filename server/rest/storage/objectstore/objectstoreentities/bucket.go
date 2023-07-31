package objectstoreentities

type Bucket struct {
	Name               string         `json:"name"`
	Region             string         `json:"region"`
	ObjectLocking      bool           `json:"object_locking"`
	LocationConstraint string         `json:"location_constraint"`
	Retention          int64          `json:"retention"`
	IsPublic           bool           `json:"is_public"`
	Tags               map[string]any `json:"tags"`
	ForceCreate        bool           `json:"force_create"`
}
