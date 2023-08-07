package services

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/constants"
	"github.com/ArkamFahry/uploadnexus/server/errors"
	"github.com/ArkamFahry/uploadnexus/server/models"
	"github.com/ArkamFahry/uploadnexus/server/storage/database/clients"
	"github.com/ArkamFahry/uploadnexus/server/storage/entities"
	"github.com/ArkamFahry/uploadnexus/server/utils"
)

type IBucketService interface {
	CreateBucket(ctx context.Context, bucketCreate models.BucketCreate) (*models.GeneralResponse, *errors.HttpError)
	UpdateBucket(ctx context.Context, id string, bucketUpdate models.BucketCreate) (*models.GeneralResponse, *errors.HttpError)
	DeleteBucket(ctx context.Context, id string) (*models.GeneralResponse, *errors.HttpError)
	GetBucketById(ctx context.Context, id string) (*models.GeneralResponse, *errors.HttpError)
	GetBuckets(ctx context.Context) (*models.GeneralResponse, *errors.HttpError)
}

type BucketService struct {
	databaseClient clients.DatabaseClient
	modelValidator utils.IModelValidator
}

var _ IBucketService = (*BucketService)(nil)

func NewBucketService(databaseClient clients.DatabaseClient, modelValidator utils.IModelValidator) *BucketService {
	return &BucketService{
		databaseClient: databaseClient,
		modelValidator: modelValidator,
	}
}

func (s *BucketService) CreateBucket(ctx context.Context, bucketCreate models.BucketCreate) (*models.GeneralResponse, *errors.HttpError) {
	validate, err := s.modelValidator.ValidateModel(bucketCreate)
	if err != nil {
		return nil, errors.NewBadRequestError(validate)
	}

	exists, err := s.databaseClient.CheckIfBucketExistsByName(ctx, bucketCreate.Name)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to check if bucket exists")
	}
	if exists {
		return nil, errors.NewBadRequestError("bucket with the name '" + bucketCreate.Name + "' already exists")
	}

	if len(bucketCreate.AllowedMimeTypes) != 0 {
		isValid, invalidMimeType := utils.ValidateMimeTypes(bucketCreate.AllowedMimeTypes)
		if !isValid {
			return nil, errors.NewBadRequestError(invalidMimeType)
		}
	}

	if &bucketCreate.IsPublic == nil {
		bucketCreate.IsPublic = false
	}

	if len(bucketCreate.AllowedMimeTypes) == 0 {
		bucketCreate.AllowedMimeTypes = []string{"*"}
	}

	bucket := entities.Bucket{
		Id:                utils.GetUUID(),
		Name:              bucketCreate.Name,
		Description:       bucketCreate.Description,
		AllowedMimeTypes:  &bucketCreate.AllowedMimeTypes,
		AllowedObjectSize: &bucketCreate.AllowedObjectSize,
		IsPublic:          bucketCreate.IsPublic,
		CreatedAt:         utils.GetTimeUnix(),
	}

	err = s.databaseClient.CreateBucket(ctx, bucket)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to create bucket")
	}

	return models.NewGeneralResponse(constants.StatusCreated, "bucket created successfully", bucket), nil
}

func (s *BucketService) UpdateBucket(ctx context.Context, id string, bucketUpdate models.BucketCreate) (*models.GeneralResponse, *errors.HttpError) {
	exists, err := s.databaseClient.CheckIfBucketExistsById(ctx, id)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to check if bucket exists")
	}
	if !exists {
		return nil, errors.NewBadRequestError("bucket with the id '" + id + "' does not exist")
	}

	validate, err := s.modelValidator.ValidateModel(bucketUpdate)
	if err != nil {
		return nil, errors.NewBadRequestError(validate)
	}

	oldBucket, err := s.databaseClient.GetBucketById(ctx, id)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to get bucket by id")
	}

	if bucketUpdate.Name != "" {
		oldBucket.Name = bucketUpdate.Name
	}

	if bucketUpdate.Description != nil {
		oldBucket.Description = bucketUpdate.Description
	}

	if len(bucketUpdate.AllowedMimeTypes) != 0 {
		isValid, invalidMimeType := utils.ValidateMimeTypes(bucketUpdate.AllowedMimeTypes)
		if !isValid {
			return nil, errors.NewBadRequestError(invalidMimeType)
		}
		oldBucket.AllowedMimeTypes = &bucketUpdate.AllowedMimeTypes
	}

	if bucketUpdate.AllowedObjectSize != 0 {
		oldBucket.AllowedObjectSize = &bucketUpdate.AllowedObjectSize
	}

	if &bucketUpdate.IsPublic != nil {
		oldBucket.IsPublic = bucketUpdate.IsPublic
	}

	*oldBucket.UpdatedAt = utils.GetTimeUnix()

	newBucket := entities.Bucket{
		Id:                oldBucket.Id,
		Name:              oldBucket.Name,
		Description:       oldBucket.Description,
		AllowedMimeTypes:  oldBucket.AllowedMimeTypes,
		AllowedObjectSize: oldBucket.AllowedObjectSize,
		IsPublic:          oldBucket.IsPublic,
		CreatedAt:         oldBucket.CreatedAt,
		UpdatedAt:         oldBucket.UpdatedAt,
	}

	err = s.databaseClient.UpdateBucket(ctx, newBucket)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to update bucket")
	}

	return models.NewGeneralResponse(constants.StatusOK, "bucket updated successfully", newBucket), nil
}

func (s *BucketService) DeleteBucket(ctx context.Context, id string) (*models.GeneralResponse, *errors.HttpError) {
	exists, err := s.databaseClient.CheckIfBucketExistsById(ctx, id)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to check if bucket exists")
	}
	if !exists {
		return nil, errors.NewBadRequestError("bucket with the id '" + id + "' does not exist")
	}

	err = s.databaseClient.DeleteBucket(ctx, id)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to delete bucket")
	}

	return models.NewGeneralResponse(constants.StatusNoContent, "bucket deleted successfully", nil), nil
}

func (s *BucketService) GetBucketById(ctx context.Context, id string) (*models.GeneralResponse, *errors.HttpError) {
	exists, err := s.databaseClient.CheckIfBucketExistsById(ctx, id)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to check if bucket exists")
	}
	if !exists {
		return nil, errors.NewBadRequestError("bucket with the id '" + id + "' does not exist")
	}

	bucket, err := s.databaseClient.GetBucketById(ctx, id)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to get bucket")
	}

	return models.NewGeneralResponse(constants.StatusOK, "bucket retrieved successfully", bucket), nil
}

func (s *BucketService) GetBuckets(ctx context.Context) (*models.GeneralResponse, *errors.HttpError) {
	buckets, err := s.databaseClient.GetBuckets(ctx)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to get buckets")
	}

	return models.NewGeneralResponse(constants.StatusOK, "buckets retrieved successfully", buckets), nil
}
