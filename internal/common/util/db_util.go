package util

import (
	"fmt"
	"go.etcd.io/bbolt"
	"ws/internal/common/apperror"
)

func GetBucket(tx *bbolt.Tx, bucketName string) (*bbolt.Bucket, error) {
	b := tx.Bucket([]byte(bucketName))
	if b == nil {
		return nil, HandleError(fmt.Sprintf("bucket not found: %v", bucketName), apperror.NotFoundBucket)
	}
	return b, nil
}
