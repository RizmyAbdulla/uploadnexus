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
	const Op errors.Op = "objectstore.DeleteObject"

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
