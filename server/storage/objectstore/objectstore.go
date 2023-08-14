package objectstore

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/envs"
	"github.com/ArkamFahry/uploadnexus/server/errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type StoreClient interface {
	// CreatePresignedPutObject to create presigned put url
	CreatePresignedPutObject(ctx context.Context, bucketName string, objectName string, expiry int64) (string, error)
	// CreatePresignedGetObject to create presigned get url
	CreatePresignedGetObject(ctx context.Context, bucketName string, objectName string, expiry int64) (string, error)

	// DeleteObject to delete object
	DeleteObject(ctx context.Context, bucketName string, objectName string) error
	// CopyObject to copy object from one location to another
	CopyObject(ctx context.Context, sourceBucketName string, sourceObjectName string, destinationBucketName string, destinationObjectName string) error
	// MoveObject to move object from one location to another
	MoveObject(ctx context.Context, sourceBucketName string, sourceObjectName string, destinationBucketName string, destinationObjectName string) error

	// RenameBucket to rename bucket
	RenameBucket(ctx context.Context, oldBucketName string, newBucketName string) error
	// EmptyBucket to empty bucket
	EmptyBucket(ctx context.Context, bucketName string) error
	// DeleteBucket to delete bucket
	DeleteBucket(ctx context.Context, bucketName string) error
}

type S3Client struct {
	s3Client          *s3.Client
	s3PresignedClient *s3.PresignClient
}

var Client StoreClient

func NewClient() (*S3Client, error) {
	const Op errors.Op = "objectstore.NewClient"
	var err error
	var s3Client *s3.Client

	endPoint := envs.EnvStoreInstance.GetEnv().ObjectStoreEndpoint
	accessKey := envs.EnvStoreInstance.GetEnv().ObjectStoreAccessKey
	secretKey := envs.EnvStoreInstance.GetEnv().ObjectStoreSecretKey
	defaultRegion := envs.EnvStoreInstance.GetEnv().ObjectStoreRegion

	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...any) (aws.Endpoint, error) {
		return aws.Endpoint{
			PartitionID:       "aws",
			URL:               endPoint,
			SigningRegion:     defaultRegion,
			HostnameImmutable: true,
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(defaultRegion),
		config.WithEndpointResolverWithOptions(resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	)

	s3Client = s3.NewFromConfig(cfg)

	if err != nil {
		return nil, errors.NewError(Op, "error creating S3 client", err)
	}

	return &S3Client{
		s3Client:          s3Client,
		s3PresignedClient: s3.NewPresignClient(s3Client),
	}, nil
}

func InitObjectStore() error {
	const Op errors.Op = "objectstore.InitObjectStore"
	var err error

	Client, err = NewClient()

	if err != nil {
		return errors.NewError(Op, "failed to create object store provider", err)
	}

	return nil
}

func GetObjectStoreClient() StoreClient {
	return Client
}
