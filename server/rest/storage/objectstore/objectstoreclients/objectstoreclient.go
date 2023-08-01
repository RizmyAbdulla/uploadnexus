package objectstoreclients

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/rest/storage/objectstore/objectstoreentities"
)

type ObjectStoreClient interface {
	CreatePresignedPutUrl(ctx context.Context, presignedUrl objectstoreentities.PresignedUrl) (string, error)
	CratedPresignedGetUrl(ctx context.Context, presignedUrl objectstoreentities.PresignedUrl) (string, error)

	DeleteObject(ctx context.Context, object objectstoreentities.Object) error
}
