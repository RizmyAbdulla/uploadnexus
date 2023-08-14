package objectstore

import (
	"context"
	"fmt"
	"github.com/ArkamFahry/uploadnexus/server/envs"
	"github.com/ArkamFahry/uploadnexus/server/errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (c *S3Client) DeleteObject(ctx context.Context, bucketName string, objectName string) error {
	const Op errors.Op = "objectstore.DeleteObjectByID"

	bucket := envs.EnvStoreInstance.GetEnv().BucketName
	object := fmt.Sprintf("%s/%s", bucketName, objectName)

	_, err := c.s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(object),
	})
	if err != nil {
		return errors.NewError(Op, "failed to delete object", err)
	}

	return nil
}

func (c *S3Client) CopyObject(ctx context.Context, sourceBucketName string, sourceObjectName string, destinationBucketName string, destinationObjectName string) error {
	const Op errors.Op = "objectstore.CopyObject"

	bucket := envs.EnvStoreInstance.GetEnv().BucketName
	source := fmt.Sprintf("%s/%s", sourceBucketName, sourceObjectName)
	destination := fmt.Sprintf("%s/%s", destinationBucketName, destinationObjectName)

	_, err := c.s3Client.CopyObject(ctx, &s3.CopyObjectInput{
		Bucket:     aws.String(bucket),
		CopySource: aws.String(source),
		Key:        aws.String(destination),
	})
	if err != nil {
		return errors.NewError(Op, "failed to copy object", err)
	}

	return nil
}

func (c *S3Client) MoveObject(ctx context.Context, sourceBucketName string, sourceObjectName string, destinationBucketName string, destinationObjectName string) error {
	const Op errors.Op = "objectstore.MoveObject"

	bucket := envs.EnvStoreInstance.GetEnv().BucketName
	source := fmt.Sprintf("%s/%s", sourceBucketName, sourceObjectName)
	destination := fmt.Sprintf("%s/%s", destinationBucketName, destinationObjectName)

	_, err := c.s3Client.CopyObject(ctx, &s3.CopyObjectInput{
		Bucket:     aws.String(bucket),
		CopySource: aws.String(source),
		Key:        aws.String(destination),
	})
	if err != nil {
		return errors.NewError(Op, "failed to copy object", err)
	}

	_, err = c.s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(source),
	})
	if err != nil {
		return errors.NewError(Op, "failed to delete object", err)
	}

	return nil
}
