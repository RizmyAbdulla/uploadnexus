package objectstore

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/envs"
	"github.com/ArkamFahry/uploadnexus/server/errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"strings"
)

func (c *S3Client) RenameBucket(ctx context.Context, oldBucketName string, newBucketName string) error {
	const Op errors.Op = "objectstore.RenameBucket"

	bucket := envs.EnvStoreInstance.GetEnv().BucketName

	objectList, err := c.s3Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(oldBucketName),
	})
	if err != nil {
		return errors.NewError(Op, "failed to list objects", err)
	}

	for _, object := range objectList.Contents {
		objectKey := aws.ToString(object.Key)
		newObjectKey := strings.Replace(objectKey, oldBucketName, newBucketName, 1)
		_, err := c.s3Client.CopyObject(ctx, &s3.CopyObjectInput{
			Bucket:     aws.String(bucket),
			CopySource: aws.String(bucket + "/" + objectKey),
			Key:        aws.String(newObjectKey),
		})
		if err != nil {
			return errors.NewError(Op, "failed to copy object", err)
		}
	}

	for _, object := range objectList.Contents {
		_, err = c.s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
			Bucket: aws.String(bucket),
			Key:    object.Key,
		})
	}

	if err != nil {
		return errors.NewError(Op, "failed to delete object", err)
	}

	return nil
}

func (c *S3Client) EmptyBucket(ctx context.Context, bucketName string) error {
	const Op errors.Op = "objectstore.EmptyBucket"

	bucket := envs.EnvStoreInstance.GetEnv().BucketName

	objectList, err := c.s3Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(bucketName),
	})
	if err != nil {
		return errors.NewError(Op, "failed to list objects", err)
	}

	for _, object := range objectList.Contents {
		_, err := c.s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
			Bucket: aws.String(bucket),
			Key:    object.Key,
		})
		if err != nil {
			return errors.NewError(Op, "failed to delete object", err)
		}
	}

	return nil
}

func (c *S3Client) DeleteBucket(ctx context.Context, bucketName string) error {
	const Op errors.Op = "objectstore.DeleteBucket"

	bucket := envs.EnvStoreInstance.GetEnv().BucketName

	objectList, err := c.s3Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(bucketName),
	})
	if err != nil {
		return errors.NewError(Op, "failed to list objects", err)
	}

	for _, object := range objectList.Contents {
		_, err = c.s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
			Bucket: aws.String(bucket),
			Key:    object.Key,
		})
		if err != nil {
			return errors.NewError(Op, "failed to delete object", err)
		}
	}

	_, err = c.s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(bucketName),
	})
	if err != nil {
		return errors.NewError(Op, "failed to delete bucket", err)
	}

	return nil
}
