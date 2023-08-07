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
)

type IObjectService interface {
	CreatePreSignedPutObject(ctx context.Context, bucketName string, objectName string) (*models.GeneralResponse, *errors.HttpError)
	CreatePreSignedGetObject(ctx context.Context, bucketName string, objectName string) (*models.GeneralResponse, *errors.HttpError)
	DeleteObject(ctx context.Context, bucketName string, objectName string) (*models.GeneralResponse, *errors.HttpError)
}

type ObjectService struct {
	databaseClient    clients.DatabaseClient
	objectStoreClient objectstore.StoreClient
}

var _ IObjectService = (*ObjectService)(nil)

func NewObjectService(databaseClient clients.DatabaseClient, objectStoreClient objectstore.StoreClient) *ObjectService {
	return &ObjectService{
		databaseClient:    databaseClient,
		objectStoreClient: objectStoreClient,
	}
}

func (s *ObjectService) CreatePreSignedPutObject(ctx context.Context, bucketName string, objectName string) (*models.GeneralResponse, *errors.HttpError) {
	exists, err := s.databaseClient.CheckIfBucketExistsByName(ctx, bucketName)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to check if bucket exists")
	}
	if !exists {
		return nil, errors.NewNotFoundError("bucket with name '" + bucketName + "' does not exist")
	}

	exists, err = s.databaseClient.CheckIfObjectExistsByBucketNameAndObjectName(ctx, bucketName, objectName)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to check if object exists")
	}
	if exists {
		return nil, errors.NewBadRequestError("object with name '" + objectName + "' already exists in bucket '" + bucketName + "")
	}

	expiry := envs.EnvStoreInstance.GetEnv().PresignedPutObjectExpiration

	url, err := s.objectStoreClient.CreatePresignedPutObject(ctx, bucketName, objectName, expiry)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to create presigned put object")
	}

	err = s.databaseClient.CreateObject(ctx, entities.Object{
		Id:           utils.GetUUID(),
		Name:         objectName,
		Bucket:       bucketName,
		MimeType:     "",
		Size:         0,
		UploadStatus: models.UploadStatusPending,
	})
	if err != nil {
		return nil, errors.NewInternalServerError("unable to create object")
	}

	preSignedPutObject := models.PreSignedObject{
		BucketName: bucketName,
		ObjectName: objectName,
		Url:        url,
		HttpMethod: constants.PUT,
		Expiry:     expiry,
	}

	return models.NewGeneralResponse(constants.StatusOK, "pre-signed put object created successfully", preSignedPutObject), nil
}

func (s *ObjectService) CreatePreSignedGetObject(ctx context.Context, bucketName string, objectName string) (*models.GeneralResponse, *errors.HttpError) {
	exists, err := s.databaseClient.CheckIfBucketExistsByName(ctx, bucketName)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to check if bucket exists")
	}
	if !exists {
		return nil, errors.NewNotFoundError("bucket with name '" + bucketName + "' does not exist")
	}

	exists, err = s.databaseClient.CheckIfObjectExistsByBucketNameAndObjectName(ctx, bucketName, objectName)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to check if object exists")
	}
	if !exists {
		return nil, errors.NewNotFoundError("object with name '" + objectName + "' does not exist in bucket '" + bucketName + "'")
	}

	expiry := envs.EnvStoreInstance.GetEnv().PresignedPutObjectExpiration

	url, err := s.objectStoreClient.CreatePresignedGetObject(ctx, bucketName, objectName, expiry)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to create presigned get object")
	}

	preSignedPutObject := models.PreSignedObject{
		BucketName: bucketName,
		ObjectName: objectName,
		Url:        url,
		HttpMethod: constants.GET,
		Expiry:     expiry,
	}

	return models.NewGeneralResponse(constants.StatusOK, "pre-signed get object created successfully", preSignedPutObject), nil
}

func (s *ObjectService) DeleteObject(ctx context.Context, bucketName string, objectName string) (*models.GeneralResponse, *errors.HttpError) {
	exists, err := s.databaseClient.CheckIfBucketExistsByName(ctx, bucketName)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to check if bucket exists")
	}
	if !exists {
		return nil, errors.NewNotFoundError("bucket with name '" + bucketName + "' does not exist")
	}

	exists, err = s.databaseClient.CheckIfObjectExistsByBucketNameAndObjectName(ctx, bucketName, objectName)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to check if object exists")
	}
	if !exists {
		return nil, errors.NewNotFoundError("object with name '" + objectName + "' does not exist in bucket '" + bucketName + "'")
	}

	object, err := s.databaseClient.GetObjectByBucketNameAndObjectName(ctx, bucketName, objectName)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to get object")
	}

	err = s.databaseClient.DeleteObject(ctx, object.Id)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to delete object")
	}

	err = s.objectStoreClient.DeleteObject(ctx, bucketName, objectName)
	if err != nil {
		return nil, errors.NewInternalServerError("unable to delete object")
	}

	return models.NewGeneralResponse(constants.StatusOK, "object deleted successfully", nil), nil
}
