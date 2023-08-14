package services

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/constants"
	"github.com/ArkamFahry/uploadnexus/server/envs"
	"github.com/ArkamFahry/uploadnexus/server/errors"
	"github.com/ArkamFahry/uploadnexus/server/models"
	"github.com/ArkamFahry/uploadnexus/server/storage/database/clients"
	"github.com/ArkamFahry/uploadnexus/server/storage/entities"
	"github.com/ArkamFahry/uploadnexus/server/storage/objectstore"
	"github.com/ArkamFahry/uploadnexus/server/utils"
	"net/url"
)

type IObjectService interface {
	CreatePreSignedPutObject(ctx context.Context, bucketName string, body []byte) (*models.PreSignedObjectResponse, *errors.HttpError)
	CreatePreSignedGetObject(ctx context.Context, bucketName string, objectName string) (*models.PreSignedObjectResponse, *errors.HttpError)
	DeleteObject(ctx context.Context, bucketName string, objectName string) (*models.ObjectGeneralResponse, *errors.HttpError)
}

type ObjectService struct {
	objectStoreClient objectstore.StoreClient
	databaseClient    clients.DatabaseClient
	modelValidator    utils.IModelValidator
}

var _ IObjectService = (*ObjectService)(nil)

func NewObjectService(objectStoreClient objectstore.StoreClient, databaseClient clients.DatabaseClient, modelValidator utils.IModelValidator) *ObjectService {
	return &ObjectService{
		objectStoreClient: objectStoreClient,
		databaseClient:    databaseClient,
		modelValidator:    modelValidator,
	}
}

func (s *ObjectService) CreatePreSignedPutObject(ctx context.Context, bucketName string, body []byte) (*models.PreSignedObjectResponse, *errors.HttpError) {
	var preSignedObjectCreate models.PresignedObjectCreate

	if bucketName == "" {
		return nil, errors.NewBadRequestError("bucket name cannot be empty")
	}

	bucketName, err := url.QueryUnescape(bucketName)
	if err != nil {
		return nil, errors.NewBadRequestError("invalid bucket name")
	}

	exists, err := s.databaseClient.BucketExistsByName(ctx, bucketName)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to check if bucket exists")
	}
	if !exists {
		return nil, errors.NewNotFoundError("bucket with name '" + bucketName + "' does not exist")
	}

	err = utils.ParseRequestBody(body, &preSignedObjectCreate)
	if err != nil {
		return nil, errors.NewBadRequestError("invalid request body")
	}

	validate, err := s.modelValidator.ValidateModel(preSignedObjectCreate)
	if err != nil {
		return nil, errors.NewBadRequestError(validate)
	}

	exists, err = s.databaseClient.ObjectExistsByBucketNameAndObjectName(ctx, bucketName, preSignedObjectCreate.ObjectName)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to check if object exists")
	}
	if exists {
		return nil, errors.NewBadRequestError("object with name '" + preSignedObjectCreate.ObjectName + "' already exists in bucket '" + bucketName + "")
	}

	bucket, err := s.databaseClient.GetBucketByName(ctx, bucketName)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to get bucket")
	}

	expiry := envs.EnvStoreInstance.GetEnv().PresignedPutObjectExpiration

	presignedUrl, err := s.objectStoreClient.CreatePresignedPutObject(ctx, bucketName, preSignedObjectCreate.ObjectName, expiry)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to create presigned put object")
	}

	err = s.databaseClient.CreateObject(ctx, entities.Object{
		Id:           utils.GetUUID(),
		Name:         preSignedObjectCreate.ObjectName,
		Bucket:       bucket.Id,
		MimeType:     preSignedObjectCreate.MimeType,
		Size:         preSignedObjectCreate.Size,
		Metadata:     preSignedObjectCreate.MetaData,
		UploadStatus: models.UploadStatusPending,
		CreatedAt:    utils.GetTimeUnix(),
	})
	if err != nil {
		return nil, errors.NewInternalServerError("unable to create object")
	}

	preSignedPutObject := models.PreSignedObject{
		BucketName: bucketName,
		ObjectName: preSignedObjectCreate.ObjectName,
		Url:        presignedUrl,
		HttpMethod: constants.PUT,
		Expiry:     expiry,
	}

	return &models.PreSignedObjectResponse{
		Code:            200,
		Message:         "pre-signed put object retrieved successfully",
		PreSignedObject: preSignedPutObject,
	}, nil
}

func (s *ObjectService) CreatePreSignedGetObject(ctx context.Context, bucketName string, objectName string) (*models.PreSignedObjectResponse, *errors.HttpError) {
	if bucketName == "" {
		return nil, errors.NewBadRequestError("bucket name cannot be empty")
	}

	bucketName, err := url.QueryUnescape(bucketName)
	if err != nil {
		return nil, errors.NewBadRequestError("invalid bucket name")
	}

	if objectName == "" {
		return nil, errors.NewBadRequestError("object name cannot be empty")
	}

	objectName, err = url.QueryUnescape(objectName)
	if err != nil {
		return nil, errors.NewBadRequestError("invalid object name")
	}

	if exists, err := s.databaseClient.BucketExistsByName(ctx, bucketName); err != nil {
		return nil, errors.NewInternalServerError("unable to check if bucket exists")
	} else if !exists {
		return nil, errors.NewNotFoundError("bucket with name '" + bucketName + "' does not exist")
	}

	if exists, err := s.databaseClient.ObjectExistsByBucketNameAndObjectName(ctx, bucketName, objectName); err != nil {
		return nil, errors.NewInternalServerError("unable to check if object exists")
	} else if !exists {
		return nil, errors.NewNotFoundError("object with name '" + objectName + "' does not exist in bucket '" + bucketName + "'")
	}

	expiry := envs.EnvStoreInstance.GetEnv().PresignedGetObjectExpiration

	presignedUrl, err := s.objectStoreClient.CreatePresignedGetObject(ctx, bucketName, objectName, expiry)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to create presigned get object")
	}

	preSignedPutObject := models.PreSignedObject{
		BucketName: bucketName,
		ObjectName: objectName,
		Url:        presignedUrl,
		HttpMethod: constants.GET,
		Expiry:     expiry,
	}

	return &models.PreSignedObjectResponse{
		Code:            200,
		Message:         "pre-signed get object retrieved successfully",
		PreSignedObject: preSignedPutObject,
	}, nil
}

func (s *ObjectService) DeleteObject(ctx context.Context, bucketName string, objectName string) (*models.ObjectGeneralResponse, *errors.HttpError) {
	if bucketName == "" {
		return nil, errors.NewBadRequestError("bucket name cannot be empty")
	}

	bucketName, err := url.QueryUnescape(bucketName)
	if err != nil {
		return nil, errors.NewBadRequestError("invalid bucket name")
	}

	if objectName == "" {
		return nil, errors.NewBadRequestError("object name cannot be empty")
	}

	objectName, err = url.QueryUnescape(objectName)
	if err != nil {
		return nil, errors.NewBadRequestError("invalid object name")
	}

	if exists, err := s.databaseClient.BucketExistsByName(ctx, bucketName); err != nil {
		return nil, errors.NewInternalServerError("unable to check if bucket exists")
	} else if !exists {
		return nil, errors.NewNotFoundError("bucket with name '" + bucketName + "' does not exist")
	}

	if exists, err := s.databaseClient.ObjectExistsByBucketNameAndObjectName(ctx, bucketName, objectName); err != nil {
		return nil, errors.NewInternalServerError("unable to check if object exists")
	} else if !exists {
		return nil, errors.NewNotFoundError("object with name '" + objectName + "' does not exist in bucket '" + bucketName + "'")
	}

	object, err := s.databaseClient.GetObjectByBucketNameAndObjectName(ctx, bucketName, objectName)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to get object")
	}

	if err := s.databaseClient.DeleteObjectByID(ctx, object.Id); err != nil {
		return nil, errors.NewInternalServerError("unable to delete object")
	}

	if err = s.objectStoreClient.DeleteObject(ctx, bucketName, objectName); err != nil {
		return nil, errors.NewInternalServerError("unable to delete object")
	}

	return &models.ObjectGeneralResponse{
		Code:    200,
		Message: "object deleted successfully",
	}, nil
}
