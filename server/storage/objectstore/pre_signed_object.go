package objectstore

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"time"
)

func (c *S3Client) CreatePresignedPutObject(ctx context.Context, bucketName string, objectName string, expiry int64) (string, error) {
	const Op errors.Op = "objectstore.CreatePresignedPutObject"

	expiryDuration := time.Duration(expiry) * time.Second

	url, err := c.s3PresignedClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectName),
	}, s3.WithPresignExpires(expiryDuration))
	if err != nil {
		return "", errors.NewError(Op, "failed to create presigned put url", err)
	}

	return url.URL, nil
}

func (c *S3Client) CratedPresignedGetObject(ctx context.Context, bucketName string, objectName string, expiry int64) (string, error) {
	const Op errors.Op = "objectstore.CreatePresignedPutObject"

	expiryDuration := time.Duration(expiry) * time.Second

	url, err := c.s3PresignedClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectName),
	}, s3.WithPresignExpires(expiryDuration))
	if err != nil {
		return "", errors.NewError(Op, "failed to create presigned get url", err)
	}

	return url.URL, nil
}
