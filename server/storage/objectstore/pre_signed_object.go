package objectstore

import (
	"context"
	"fmt"
	"github.com/ArkamFahry/uploadnexus/server/envs"
	"github.com/ArkamFahry/uploadnexus/server/errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"time"
)

func (c *S3Client) CreatePresignedPutObject(ctx context.Context, bucketName string, objectName string, expiry int64) (string, error) {
	const Op errors.Op = "objectstore.CreatePresignedPutObject"

	bucket := envs.EnvStoreInstance.GetEnv().BucketName
	expiryDuration := time.Duration(expiry) * time.Second
	object := fmt.Sprintf("%s/%s", bucketName, objectName)

	url, err := c.s3PresignedClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(object),
	}, s3.WithPresignExpires(expiryDuration))
	if err != nil {
		return "", errors.NewError(Op, "failed to create presigned put url", err)
	}

	return url.URL, nil
}

func (c *S3Client) CreatePresignedGetObject(ctx context.Context, bucketName string, objectName string, expiry int64) (string, error) {
	const Op errors.Op = "objectstore.CreatePresignedPutObject"

	bucket := envs.EnvStoreInstance.GetEnv().BucketName
	expiryDuration := time.Duration(expiry) * time.Second
	object := fmt.Sprintf("%s/%s", bucketName, objectName)

	url, err := c.s3PresignedClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(object),
	}, s3.WithPresignExpires(expiryDuration))
	if err != nil {
		return "", errors.NewError(Op, "failed to create presigned get url", err)
	}

	return url.URL, nil
}
