package blob

import (
	"encoding/json"
	"mime/multipart"
	"strings"

	"github.com/johncalvinroberts/cryp/internal/errors"
	"github.com/johncalvinroberts/cryp/internal/storage"
	"github.com/rs/xid"
)

type BlobService struct {
	storageSrv                            *storage.StorageService
	blobBucketName, blobPointerBucketName string
}

type BlobPointers struct {
	Blobs []string `json:"blobs"`
	Count int      `json:"count"`
}

func (svc *BlobService) UploadFile(file multipart.File, email string) (string, error) {
	guid := xid.New()
	id := guid.String()
	key := storage.ComposeKey(id, email)
	location, err := svc.storageSrv.Write(svc.blobBucketName, key, file)
	if err != nil {
		return "", errors.ErrDataCreationFailure
	}
	_, err = svc.AddBlobPointer(key, email)
	if err != nil {
		return "", errors.ErrDataCreationFailure
	}
	return location, nil
}

func (svc *BlobService) AddBlobPointer(key, email string) (string, error) {
	pointersStr, err := svc.storageSrv.ReadToString(svc.blobPointerBucketName, key)
	if err != nil {
		return "", err
	}
	// TODO: need to lock the s3 object to prevent concurrent writes to the same object resulting in data loss
	pointers := &BlobPointers{}
	json.Unmarshal([]byte(pointersStr), pointers)
	pointers.Blobs = append(pointers.Blobs, key)
	pointers.Count++
	nextPointers, err := json.Marshal(pointers)
	if err != nil {
		return "", err
	}
	_, err = svc.storageSrv.Write(svc.blobPointerBucketName, key, strings.NewReader(string(nextPointers)))
	if err != nil {
		return "", err
	}
	return "", nil
}

func InitBlobService(storageSrv *storage.StorageService, blobBucketName, blobPointerBucketName string) *BlobService {
	return &BlobService{
		storageSrv:            storageSrv,
		blobBucketName:        blobBucketName,
		blobPointerBucketName: blobPointerBucketName,
	}
}
