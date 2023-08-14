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

	if err := utils.ParseRequestBody(body, &bucketCreate); err != nil {
		return nil, errors.NewBadRequestError("invalid request body")
	}

	if validate, err := s.modelValidator.ValidateModel(bucketCreate); err != nil {
		return nil, errors.NewBadRequestError(validate)
	}

	if exists, err := s.databaseClient.BucketExistsByName(ctx, bucketCreate.Name); err != nil {
		return nil, errors.NewInternalServerError("unable to check if bucket exists by name")
	} else if exists {
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

	if err := s.databaseClient.CreateBucket(ctx, bucket); err != nil {
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

	if err := utils.ParseRequestBody(body, &bucketUpdate); err != nil {
		return nil, errors.NewBadRequestError("invalid request body")
	}

	if exists, err := s.databaseClient.BucketExistsByID(ctx, id); err != nil {
		return nil, errors.NewInternalServerError("unable to check if bucket exists")
	} else if !exists {
		return nil, errors.NewNotFoundError("bucket with the id '" + id + "' does not exist")
	}

	if validate, err := s.modelValidator.ValidateModel(bucketUpdate); err != nil {
		return nil, errors.NewBadRequestError(validate)
	}

	oldBucket, err := s.databaseClient.GetBucketById(ctx, id)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to get bucket")
	}

	if bucketUpdate.Name != "" {
		if exists, err := s.databaseClient.BucketExistsByName(ctx, bucketUpdate.Name); err != nil {
			return nil, errors.NewInternalServerError("unable to check if bucket exists")
		} else if exists {
			return nil, errors.NewBadRequestError("bucket with the name '" + bucketUpdate.Name + "' already exists")
		}

		if err := s.objectStoreClient.RenameBucket(ctx, oldBucket.Name, bucketUpdate.Name); err != nil {
			return nil, errors.NewInternalServerError("unable to rename bucket in object store")
		}

		if err := s.databaseClient.UpdateObjectBucketName(ctx, oldBucket.Name, bucketUpdate.Name); err != nil {
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

	if err := s.databaseClient.UpdateBucket(ctx, newBucket); err != nil {
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

	if exists, err := s.databaseClient.BucketExistsByID(ctx, id); err != nil {
		return nil, errors.NewInternalServerError("unable to check if bucket exists")
	} else if !exists {
		return nil, errors.NewNotFoundError("bucket with the id '" + id + "' does not exist")
	}

	bucket, err := s.databaseClient.GetBucketById(ctx, id)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to get bucket")
	}

	if err := s.objectStoreClient.DeleteBucket(ctx, bucket.Name); err != nil {
		return nil, errors.NewInternalServerError("unable to delete bucket from object store")
	}

	if err := s.databaseClient.DeleteObjectsByBucketID(ctx, bucket.Name); err != nil {
		return nil, errors.NewInternalServerError("unable to delete objects from database")
	}

	if err := s.databaseClient.DeleteBucketByID(ctx, bucket.Id); err != nil {
		return nil, errors.NewInternalServerError("unable to delete bucket from database")
	}

	return &models.BucketGeneralResponse{
		Code:    constants.StatusOK,
		Message: "bucket deleted",
	}, nil
}

func (s *BucketService) GetBucketById(ctx context.Context, id string) (*models.BucketResponse, *errors.HttpError) {
	if exists, err := s.databaseClient.BucketExistsByID(ctx, id); err != nil {
		return nil, errors.NewInternalServerError("unable to check if bucket exists")
	} else if !exists {
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
	var bucketList models.BucketListResponse

	buckets, err := s.databaseClient.ListBuckets(ctx)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to get buckets")
	}

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
	} else {
		return nil, errors.NewNotFoundError("no buckets found")
	}

	return &bucketList, nil
}

func (s *BucketService) EmptyBucket(ctx context.Context, id string) (*models.BucketGeneralResponse, *errors.HttpError) {
	if exists, err := s.databaseClient.BucketExistsByID(ctx, id); err != nil {
		return nil, errors.NewInternalServerError("unable to check if bucket exists")
	} else if !exists {
		return nil, errors.NewNotFoundError("bucket with the id '" + id + "' does not exist")
	}

	bucket, err := s.databaseClient.GetBucketById(ctx, id)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to get bucket")
	}

	if err := s.objectStoreClient.EmptyBucket(ctx, bucket.Name); err != nil {
		return nil, errors.NewInternalServerError("unable to empty bucket from object store")
	}

	if err := s.databaseClient.DeleteObjectsByBucketID(ctx, bucket.Id); err != nil {
		return nil, errors.NewInternalServerError("unable to delete objects from database")
	}

	return &models.BucketGeneralResponse{
		Code:    constants.StatusOK,
		Message: "bucket emptied successfully",
	}, nil
}
