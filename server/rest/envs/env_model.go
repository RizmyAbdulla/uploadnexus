package envs

type Env struct {
	AppName        string `json:"APP_NAME" mapstructure:"APP_NAME"`
	AppHost        string `json:"APP_HOST" mapstructure:"APP_HOST"`
	AppPort        string `json:"APP_PORT" mapstructure:"APP_PORT"`
	AppUrl         string `json:"APP_URL" mapstructure:"APP_URL"`
	AppIdType      string `json:"APP_ID_TYPE" mapstructure:"APP_ID_TYPE"`
	AppEnv         string `json:"APP_ENV" mapstructure:"APP_ENV"`
	AppSecret      string `json:"APP_SECRET" mapstructure:"APP_SECRET"`
	AppAdminSecret string `json:"APP_ADMIN_SECRET" mapstructure:"APP_ADMIN_SECRET"`

	DatabaseType        string `json:"DATABASE_TYPE" mapstructure:"DATABASE_TYPE"`
	DatabaseHost        string `json:"DATABASE_HOST" mapstructure:"DATABASE_HOST"`
	DatabasePort        string `json:"DATABASE_PORT" mapstructure:"DATABASE_PORT"`
	DatabaseSsl         string `json:"DATABASE_SSL" mapstructure:"DATABASE_SSL"`
	DatabaseUser        string `json:"DATABASE_USER" mapstructure:"DATABASE_USER"`
	DatabasePassword    string `json:"DATABASE_PASSWORD" mapstructure:"DATABASE_PASSWORD"`
	DatabaseName        string `json:"DATABASE_NAME" mapstructure:"DATABASE_NAME"`
	DatabaseAutoMigrate bool   `json:"DATABASE_AUTO_MIGRATE" mapstructure:"DATABASE_AUTO_MIGRATE"`

	ObjectStoreType      string `json:"OBJECT_STORE_TYPE" mapstructure:"OBJECT_STORE_TYPE"`
	ObjectStoreEndpoint  string `json:"OBJECT_STORE_ENDPOINT" mapstructure:"OBJECT_STORE_ENDPOINT"`
	ObjectStoreAccessKey string `json:"OBJECT_STORE_ACCESS_KEY" mapstructure:"OBJECT_STORE_ACCESS_KEY"`
	ObjectStoreSecretKey string `json:"OBJECT_STORE_SECRET_KEY" mapstructure:"OBJECT_STORE_SECRET_KEY"`
	ObjectStoreSsl       bool   `json:"OBJECT_STORE_SSL" mapstructure:"OBJECT_STORE_SSL"`
	ObjectStoreRegion    string `json:"OBJECT_STORE_REGION" mapstructure:"OBJECT_STORE_REGION"`

	Buckets []Bucket `json:"Buckets" mapstructure:"Buckets"`

	CacheType     string `json:"CACHE_TYPE" mapstructure:"CACHE_TYPE"`
	CacheHost     string `json:"CACHE_HOST" mapstructure:"CACHE_HOST"`
	CachePort     string `json:"CACHE_PORT" mapstructure:"CACHE_PORT"`
	CacheUser     string `json:"CACHE_USER" mapstructure:"CACHE_USER"`
	CachePassword string `json:"CACHE_PASSWORD" mapstructure:"CACHE_PASSWORD"`

	EventStoreType     string `json:"EVENT_STORE_TYPE" mapstructure:"EVENT_STORE_TYPE"`
	EventStoreHost     string `json:"EVENT_STORE_HOST" mapstructure:"EVENT_STORE_HOST"`
	EventStorePort     string `json:"EVENT_STORE_PORT" mapstructure:"EVENT_STORE_PORT"`
	EventStoreUser     string `json:"EVENT_STORE_USER" mapstructure:"EVENT_STORE_USER"`
	EventStorePassword string `json:"EVENT_STORE_PASSWORD" mapstructure:"EVENT_STORE_PASSWORD"`
}

type Bucket struct {
	BucketName        string `json:"BUCKET_NAME" mapstructure:"BUCKET_NAME"`
	BucketDescription string `json:"BUCKET_DESCRIPTION" mapstructure:"BUCKET_DESCRIPTION"`
}
