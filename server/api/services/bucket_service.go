package services

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/constants"
	"github.com/ArkamFahry/uploadnexus/server/errors"
	"github.com/ArkamFahry/uploadnexus/server/models"
	"github.com/ArkamFahry/uploadnexus/server/storage/database/clients"
	"github.com/ArkamFahry/uploadnexus/server/utils"
)

type IBucketService interface {
	CreateBucket(ctx context.Context, bucketDto models.BucketCreate) (*models.GeneralResponse, *errors.HttpError)
	UpdateBucket(ctx context.Context, id string, bucketDto models.BucketCreate) (*models.GeneralResponse, *errors.HttpError)
	DeleteBucket(ctx context.Context, id string) (*models.GeneralResponse, *errors.HttpError)
	GetBucketById(ctx context.Context, id string) (*models.GeneralResponse, *errors.HttpError)
	GetBuckets(ctx context.Context) (*models.GeneralResponse, *errors.HttpError)
}

type BucketService struct {
	databaseClient clients.DatabaseClient
}

var _ IBucketService = (*BucketService)(nil)

func NewBucketService(databaseClient clients.DatabaseClient) *BucketService {
	return &BucketService{
		databaseClient: databaseClient,
	}
}

func (s *BucketService) CreateBucket(ctx context.Context, bucketDto models.BucketCreate) (*models.GeneralResponse, *errors.HttpError) {
	validate, err := utils.ValidateModel(bucketDto)
	if err != nil {
		return nil, errors.NewBadRequestError("invalid input", validate)
	}

	exists, err := s.databaseClient.CheckIfBucketExistsByName(ctx, bucketDto.Name)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to check if bucket exists by name", nil)
	}
	if exists {
		return nil, errors.NewBadRequestError("bucket with the name '"+bucketDto.Name+"' already exists", nil)
	}

	if len(bucketDto.AllowedMimeTypes) != 0 {
		isValid, invalidMimeType := utils.ValidateMimeTypes(bucketDto.AllowedMimeTypes)
		if !isValid {
			return nil, errors.NewBadRequestError("invalid mime types", invalidMimeType)
		}
	}

	if &bucketDto.IsPublic == nil {
		bucketDto.IsPublic = false
	}

	if len(bucketDto.AllowedMimeTypes) == 0 {
		bucketDto.AllowedMimeTypes = []string{"*"}
	}

	if &bucketDto.FileSizeLimit == nil {
		bucketDto.FileSizeLimit = 0
	}

	bucket := models.Bucket{
		Id:               utils.GetUUID(),
		Name:             bucketDto.Name,
		Description:      bucketDto.Description,
		AllowedMimeTypes: bucketDto.AllowedMimeTypes,
		FileSizeLimit:    bucketDto.FileSizeLimit,
		IsPublic:         bucketDto.IsPublic,
		CreatedAt:        utils.GetTimeUnix(),
	}

	err = s.databaseClient.CreateBucket(ctx, bucket)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to create bucket", nil)
	}

	return models.NewGeneralResponse(constants.StatusCreated, "bucket created successfully", bucket), nil
}

func (s *BucketService) UpdateBucket(ctx context.Context, id string, bucketDto models.BucketCreate) (*models.GeneralResponse, *errors.HttpError) {
	exists, err := s.databaseClient.CheckIfBucketExistsById(ctx, id)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to check if bucket exists by id", nil)
	}
	if !exists {
		return nil, errors.NewBadRequestError("bucket with the id '"+id+"' does not exist", nil)
	}

	validate, err := utils.ValidateModel(bucketDto)
	if err != nil {
		return nil, errors.NewBadRequestError("invalid input", validate)
	}

	oldBucket, err := s.databaseClient.GetBucketById(ctx, id)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to get bucket by id", nil)
	}

	if bucketDto.Name != "" {
		oldBucket.Name = bucketDto.Name
	}

	if bucketDto.Description != nil {
		oldBucket.Description = bucketDto.Description
	}

	if len(bucketDto.AllowedMimeTypes) != 0 {
		isValid, invalidMimeType := utils.ValidateMimeTypes(bucketDto.AllowedMimeTypes)
		if !isValid {
			return nil, errors.NewBadRequestError("invalid mime types", invalidMimeType)
		}
		oldBucket.AllowedMimeTypes = bucketDto.AllowedMimeTypes
	}

	if bucketDto.FileSizeLimit != 0 {
		oldBucket.FileSizeLimit = bucketDto.FileSizeLimit
	}

	if &bucketDto.IsPublic != nil {
		oldBucket.IsPublic = bucketDto.IsPublic
	}

	oldBucket.UpdatedAt = utils.GetTimeUnix()

	newBucket := models.Bucket{
		Id:               oldBucket.Id,
		Name:             oldBucket.Name,
		Description:      oldBucket.Description,
		AllowedMimeTypes: oldBucket.AllowedMimeTypes,
		FileSizeLimit:    oldBucket.FileSizeLimit,
		IsPublic:         oldBucket.IsPublic,
		CreatedAt:        oldBucket.CreatedAt,
		UpdatedAt:        oldBucket.UpdatedAt,
	}

	err = s.databaseClient.UpdateBucket(ctx, newBucket)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to update bucket", nil)
	}

	return models.NewGeneralResponse(constants.StatusOK, "bucket updated successfully", newBucket), nil
}

func (s *BucketService) DeleteBucket(ctx context.Context, id string) (*models.GeneralResponse, *errors.HttpError) {
	exists, err := s.databaseClient.CheckIfBucketExistsById(ctx, id)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to check if bucket exists by id", nil)
	}
	if !exists {
		return nil, errors.NewBadRequestError("bucket with the id '"+id+"' does not exist", nil)
	}

	err = s.databaseClient.DeleteBucket(ctx, id)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to delete bucket", nil)
	}

	return models.NewGeneralResponse(constants.StatusNoContent, "bucket deleted successfully", nil), nil
}

func (s *BucketService) GetBucketById(ctx context.Context, id string) (*models.GeneralResponse, *errors.HttpError) {
	exists, err := s.databaseClient.CheckIfBucketExistsById(ctx, id)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to check if bucket exists by id", nil)
	}
	if !exists {
		return nil, errors.NewBadRequestError("bucket with the id '"+id+"' does not exist", nil)
	}

	bucket, err := s.databaseClient.GetBucketById(ctx, id)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to get bucket by id", nil)
	}

	return models.NewGeneralResponse(constants.StatusOK, "bucket retrieved successfully", bucket), nil
}

func (s *BucketService) GetBuckets(ctx context.Context) (*models.GeneralResponse, *errors.HttpError) {
	buckets, err := s.databaseClient.GetBuckets(ctx)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to get buckets", nil)
	}

	return models.NewGeneralResponse(constants.StatusOK, "buckets retrieved successfully", buckets), nil
}
