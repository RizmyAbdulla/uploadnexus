package services

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/constants"
	"github.com/ArkamFahry/uploadnexus/server/errors"
	"github.com/ArkamFahry/uploadnexus/server/models"
	"github.com/ArkamFahry/uploadnexus/server/storage/database/clients"
	"github.com/ArkamFahry/uploadnexus/server/storage/entities"
	"github.com/ArkamFahry/uploadnexus/server/utils"
	"net/url"
)

type IBucketService interface {
	CreateBucket(ctx context.Context, body []byte) (*models.GeneralResponse, *errors.HttpError)
	UpdateBucket(ctx context.Context, id string, body []byte) (*models.GeneralResponse, *errors.HttpError)
	DeleteBucket(ctx context.Context, id string) (*models.GeneralResponse, *errors.HttpError)
	GetBucketById(ctx context.Context, id string) (*models.GeneralResponse, *errors.HttpError)
	ListBuckets(ctx context.Context) (*models.GeneralResponse, *errors.HttpError)
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

func (s *BucketService) CreateBucket(ctx context.Context, body []byte) (*models.GeneralResponse, *errors.HttpError) {
	var bucketCreate models.BucketCreate

	err := utils.ParseRequestBody(body, &bucketCreate)
	if err != nil {
		return nil, errors.NewBadRequestError("invalid request body")
	}

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
		isValid, err := utils.ValidateMimeTypes(bucketCreate.AllowedMimeTypes)
		if !isValid {
			return nil, errors.NewInvalidMediaTypeError(err.Error())
		}
	}

	if &bucketCreate.IsPublic == nil {
		bucketCreate.IsPublic = false
	}

	if len(bucketCreate.AllowedMimeTypes) == 0 {
		bucketCreate.AllowedMimeTypes = []string{"*"}
	}

	bucket := entities.Bucket{
		Id:                bucketCreate.Name,
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

func (s *BucketService) UpdateBucket(ctx context.Context, id string, body []byte) (*models.GeneralResponse, *errors.HttpError) {
	var bucketUpdate models.BucketUpdate

	if id == "" {
		return nil, errors.NewBadRequestError("id cannot be empty")
	}

	id, err := url.QueryUnescape(id)
	if err != nil {
		return nil, errors.NewBadRequestError("invalid id")
	}

	err = utils.ParseRequestBody(body, &bucketUpdate)
	if err != nil {
		return nil, errors.NewBadRequestError("invalid request body")
	}

	exists, err := s.databaseClient.CheckIfBucketExistsById(ctx, id)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to check if bucket exists")
	}
	if !exists {
		return nil, errors.NewNotFoundError("bucket with the id '" + id + "' does not exist")
	}

	validate, err := s.modelValidator.ValidateModel(bucketUpdate)
	if err != nil {
		return nil, errors.NewBadRequestError(validate)
	}

	oldBucket, err := s.databaseClient.GetBucketById(ctx, id)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to get bucket")
	}

	if bucketUpdate.Name != "" {
		exists, err := s.databaseClient.CheckIfBucketExistsByName(ctx, bucketUpdate.Name)
		if err != nil {
			return nil, errors.NewInternalServerError("unable to check if bucket exists")
		}
		if exists {
			return nil, errors.NewBadRequestError("bucket with the name '" + bucketUpdate.Name + "' already exists")
		}
		oldBucket.Name = bucketUpdate.Name
	}

	if bucketUpdate.Description != nil {
		oldBucket.Description = bucketUpdate.Description
	}

	if len(bucketUpdate.AllowedMimeTypes) != 0 {
		isValid, err := utils.ValidateMimeTypes(bucketUpdate.AllowedMimeTypes)
		if !isValid {
			return nil, errors.NewInvalidMediaTypeError(err.Error())
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
	if id == "" {
		return nil, errors.NewBadRequestError("id cannot be empty")
	}

	id, err := url.QueryUnescape(id)
	if err != nil {
		return nil, errors.NewBadRequestError("invalid id")
	}

	exists, err := s.databaseClient.CheckIfBucketExistsById(ctx, id)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to check if bucket exists")
	}
	if !exists {
		return nil, errors.NewNotFoundError("bucket with the id '" + id + "' does not exist")
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
		return nil, errors.NewNotFoundError("bucket with the id '" + id + "' does not exist")
	}

	bucket, err := s.databaseClient.GetBucketById(ctx, id)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to get bucket")
	}

	return models.NewGeneralResponse(constants.StatusOK, "bucket retrieved successfully", bucket), nil
}

func (s *BucketService) ListBuckets(ctx context.Context) (*models.GeneralResponse, *errors.HttpError) {
	buckets, err := s.databaseClient.ListBuckets(ctx)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to get buckets")
	}

	return models.NewGeneralResponse(constants.StatusOK, "buckets retrieved successfully", buckets), nil
}
