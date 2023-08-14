package services

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/constants"
	"github.com/ArkamFahry/uploadnexus/server/errors"
	"github.com/ArkamFahry/uploadnexus/server/models"
	"github.com/ArkamFahry/uploadnexus/server/storage/database/clients"
	"github.com/ArkamFahry/uploadnexus/server/storage/entities"
	"github.com/ArkamFahry/uploadnexus/server/storage/objectstore"
	"github.com/ArkamFahry/uploadnexus/server/utils"
	"net/url"
)

type IBucketService interface {
	CreateBucket(ctx context.Context, body []byte) (*models.BucketResponse, *errors.HttpError)
	UpdateBucket(ctx context.Context, id string, body []byte) (*models.BucketResponse, *errors.HttpError)
	DeleteBucket(ctx context.Context, id string) (*models.BucketGeneralResponse, *errors.HttpError)
	GetBucketById(ctx context.Context, id string) (*models.BucketResponse, *errors.HttpError)
	ListBuckets(ctx context.Context) (*models.BucketListResponse, *errors.HttpError)
	EmptyBucket(ctx context.Context, id string) (*models.BucketGeneralResponse, *errors.HttpError)
}

type BucketService struct {
	objectStoreClient objectstore.StoreClient
	databaseClient    clients.DatabaseClient
	modelValidator    utils.IModelValidator
}

var _ IBucketService = (*BucketService)(nil)

func NewBucketService(objectStoreClient objectstore.StoreClient, databaseClient clients.DatabaseClient, modelValidator utils.IModelValidator) *BucketService {
	return &BucketService{
		objectStoreClient: objectStoreClient,
		databaseClient:    databaseClient,
		modelValidator:    modelValidator,
	}
}

func (s *BucketService) CreateBucket(ctx context.Context, body []byte) (*models.BucketResponse, *errors.HttpError) {
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

	return &models.BucketResponse{
		Code:    constants.StatusOK,
		Message: "bucket created",
		Bucket: models.Bucket{
			Id:                bucket.Id,
			Name:              bucket.Name,
			Description:       bucket.Description,
			AllowedMimeTypes:  bucket.AllowedMimeTypes,
			AllowedObjectSize: bucket.AllowedObjectSize,
			IsPublic:          bucket.IsPublic,
			CreatedAt:         bucket.CreatedAt,
			UpdatedAt:         bucket.UpdatedAt,
		},
	}, nil
}

func (s *BucketService) UpdateBucket(ctx context.Context, id string, body []byte) (*models.BucketResponse, *errors.HttpError) {
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

		err = s.objectStoreClient.RenameBucket(ctx, oldBucket.Name, bucketUpdate.Name)
		if err != nil {
			return nil, errors.NewInternalServerError("unable to rename bucket in object store")
		}

		err = s.databaseClient.UpdateObjectBucketName(ctx, oldBucket.Name, bucketUpdate.Name)
		if err != nil {
			return nil, errors.NewInternalServerError("unable to update bucket name in database")
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

	updateAt := utils.GetTimeUnix()
	oldBucket.UpdatedAt = &updateAt

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

	return &models.BucketResponse{
		Code:    constants.StatusOK,
		Message: "bucket updated",
		Bucket: models.Bucket{
			Id:                newBucket.Id,
			Name:              newBucket.Name,
			Description:       newBucket.Description,
			AllowedMimeTypes:  newBucket.AllowedMimeTypes,
			AllowedObjectSize: newBucket.AllowedObjectSize,
			IsPublic:          newBucket.IsPublic,
			CreatedAt:         newBucket.CreatedAt,
			UpdatedAt:         newBucket.UpdatedAt,
		},
	}, nil
}

func (s *BucketService) DeleteBucket(ctx context.Context, id string) (*models.BucketGeneralResponse, *errors.HttpError) {
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

	bucket, err := s.databaseClient.GetBucketById(ctx, id)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to get bucket")
	}

	err = s.objectStoreClient.DeleteBucket(ctx, bucket.Name)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to delete bucket from object store")
	}

	err = s.databaseClient.DeleteObjectsByBucketId(ctx, bucket.Name)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to delete objects from database")
	}

	err = s.databaseClient.DeleteBucketById(ctx, bucket.Id)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to delete bucket from database")
	}

	return &models.BucketGeneralResponse{
		Code:    constants.StatusOK,
		Message: "bucket deleted",
	}, nil
}

func (s *BucketService) GetBucketById(ctx context.Context, id string) (*models.BucketResponse, *errors.HttpError) {
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

	return &models.BucketResponse{
		Code:    constants.StatusOK,
		Message: "bucket retrieved successfully",
		Bucket: models.Bucket{
			Id:                bucket.Id,
			Name:              bucket.Name,
			Description:       bucket.Description,
			AllowedMimeTypes:  bucket.AllowedMimeTypes,
			AllowedObjectSize: bucket.AllowedObjectSize,
			IsPublic:          bucket.IsPublic,
			CreatedAt:         bucket.CreatedAt,
			UpdatedAt:         bucket.UpdatedAt,
		},
	}, nil
}

func (s *BucketService) ListBuckets(ctx context.Context) (*models.BucketListResponse, *errors.HttpError) {
	buckets, err := s.databaseClient.ListBuckets(ctx)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to get buckets")
	}

	var bucketList models.BucketListResponse

	if buckets != nil {
		for _, bucket := range *buckets {
			bucketList.Bucket = append(bucketList.Bucket, models.Bucket{
				Id:                bucket.Id,
				Name:              bucket.Name,
				Description:       bucket.Description,
				AllowedMimeTypes:  bucket.AllowedMimeTypes,
				AllowedObjectSize: bucket.AllowedObjectSize,
				IsPublic:          bucket.IsPublic,
				CreatedAt:         bucket.CreatedAt,
				UpdatedAt:         bucket.UpdatedAt,
			})
		}
		bucketList.Code = constants.StatusOK
		bucketList.Message = "buckets retrieved successfully"
	}

	return &bucketList, nil
}

func (s *BucketService) EmptyBucket(ctx context.Context, id string) (*models.BucketGeneralResponse, *errors.HttpError) {
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

	err = s.objectStoreClient.EmptyBucket(ctx, bucket.Name)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to empty bucket from object store")
	}

	err = s.databaseClient.DeleteObjectsByBucketId(ctx, bucket.Id)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to delete objects from database")
	}

	return &models.BucketGeneralResponse{
		Code:    constants.StatusOK,
		Message: "bucket emptied successfully",
	}, nil
}
