package services

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/errors"
	"github.com/ArkamFahry/uploadnexus/server/models"
)

type IPreSignedUrlService interface {
	CreatePresignedUrl(ctx context.Context, bucketId string) (*models.GeneralResponse, *errors.HttpError)
}
