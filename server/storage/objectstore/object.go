package objectstore

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (c *S3Client) DeleteObject(ctx context.Context, bucketName string, objectName string) error {
	const Op errors.Op = "objectstore.DeleteObject"

	_, err := c.s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectName),
	})
	if err != nil {
		return errors.NewError(Op, errors.Msg("failed to delete object"), err)
	}

	return nil
}
