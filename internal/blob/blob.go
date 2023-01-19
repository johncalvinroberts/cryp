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

	_, err = svc.AddBlobPointer(location, email)
	if err != nil {
		return "", errors.ErrDataCreationFailure
	}
	return location, nil
}

func (svc *BlobService) AddBlobPointer(keyToAdd, email string) (string, error) {
	var (
		pointersStr string
		pointers    = &BlobPointers{}
		exists, err = svc.storageSrv.Exists(svc.blobPointerBucketName, email)
	)
	if err != nil {
		return "", err
	}
	if exists {
		res, err := svc.storageSrv.ReadToString(svc.blobPointerBucketName, email)
		if err != nil {
			return "", err
		}
		pointersStr = res
	} else {
		res, err := svc.CreateBlobPointer(email)
		if err != nil {
			return "", err
		}
		pointersStr = res
	}
	err = json.Unmarshal([]byte(pointersStr), pointers)
	if err != nil {
		return "", err
	}
	// TODO: need to lock the s3 object to prevent concurrent writes to the same object resulting in data loss
	pointers.Blobs = append(pointers.Blobs, keyToAdd)
	pointers.Count++
	nextPointers, err := json.Marshal(pointers)
	if err != nil {
		return "", err
	}
	_, err = svc.storageSrv.Write(svc.blobPointerBucketName, keyToAdd, strings.NewReader(string(nextPointers)))
	if err != nil {
		return "", err
	}
	return "", nil
}

func (svc *BlobService) CreateBlobPointer(email string) (string, error) {
	pointers := &BlobPointers{}
	nextPointers, err := json.Marshal(pointers)
	if err != nil {
		return "", nil
	}
	_, err = svc.storageSrv.Write(svc.blobPointerBucketName, email, strings.NewReader(string(nextPointers)))
	if err != nil {
		return "", nil
	}
	return string(nextPointers), err
}

func (svc *BlobService) ListBlobs(email string) (*BlobPointers, error) {
	// TODO: pagination?
	pointersStr, err := svc.storageSrv.ReadToString(svc.blobPointerBucketName, email)
	if err != nil {
		return nil, err
	}
	pointers := &BlobPointers{}
	err = json.Unmarshal([]byte(pointersStr), pointers)
	return pointers, err
}

func InitBlobService(storageSrv *storage.StorageService, blobBucketName, blobPointerBucketName string) *BlobService {
	return &BlobService{
		storageSrv:            storageSrv,
		blobBucketName:        blobBucketName,
		blobPointerBucketName: blobPointerBucketName,
	}
}
