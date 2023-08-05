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
	CreatePresignedPutObject(ctx context.Context, bucketName string, objectName string, expiry int64) (string, error)
	CratedPresignedGetObject(ctx context.Context, bucketName string, objectName string, expiry int64) (string, error)

	DeleteObject(ctx context.Context, bucketName string, objectName string) error
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
