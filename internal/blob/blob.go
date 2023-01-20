package blob

import (
	"encoding/json"
	"mime/multipart"
	"strings"
	"time"

	"github.com/johncalvinroberts/cryp/internal/errors"
	"github.com/johncalvinroberts/cryp/internal/storage"
	"github.com/johncalvinroberts/cryp/internal/utils"
	"github.com/rs/xid"
)

type BlobService struct {
	storageSrv                            *storage.StorageService
	blobBucketName, blobPointerBucketName string
	emailMaskSecret                       string
}

func (svc *BlobService) UploadFile(file multipart.File, email string) (*Blob, error) {
	guid := xid.New()
	id := guid.String()
	key := storage.ComposeKey(id, utils.EncryptMessage(svc.emailMaskSecret, email))
	location, err := svc.storageSrv.Write(svc.blobBucketName, key, file)
	if err != nil {
		return nil, errors.ErrDataCreationFailure
	}
	blob, err := svc.AddBlobPointer(location, key, email)
	if err != nil {
		return nil, errors.ErrDataCreationFailure
	}
	return blob, nil
}

func (svc *BlobService) AddBlobPointer(url, key, email string) (*Blob, error) {
	var (
		now               = time.Now().Unix()
		blobToAdd         = &Blob{Url: url, CreatedAt: now, UpdatedAt: now, Key: key}
		blobPointers, err = svc.FindOrCreateBlobPointers(email)
	)
	if err != nil {
		return nil, err
	}
	// TODO: need to lock the s3 object to prevent concurrent writes to the same object resulting in data loss
	blobPointers.Blobs = append(blobPointers.Blobs, *blobToAdd)
	blobPointers.Count++
	encodedPointers, err := json.Marshal(blobPointers)
	if err != nil {
		return nil, err
	}
	_, err = svc.storageSrv.Write(svc.blobPointerBucketName, email, strings.NewReader(string(encodedPointers)))
	if err != nil {
		return nil, err
	}
	return blobToAdd, nil
}

func (svc *BlobService) FindOrCreateBlobPointers(email string) (*BlobPointers, error) {
	var (
		blobPointers = &BlobPointers{}
		exists, err  = svc.storageSrv.Exists(svc.blobPointerBucketName, email)
	)
	if err != nil {
		return nil, err
	}
	if exists {
		var existingJSONPointers string
		// read directly and copy to blobPointers
		existingJSONPointers, err = svc.storageSrv.ReadToString(svc.blobPointerBucketName, email)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(existingJSONPointers), blobPointers)
	} else {
		var emptyPointersJSON []byte
		// if doesn't exist, write the empty value to s3
		emptyPointersJSON, err = json.Marshal(blobPointers)
		if err != nil {
			return nil, err
		}
		// write to db
		_, err = svc.storageSrv.Write(svc.blobPointerBucketName, email, strings.NewReader(string(emptyPointersJSON)))
	}
	return blobPointers, err
}

func (svc *BlobService) ListBlobs(email string) (*BlobPointers, error) {
	// TODO: pagination?
	pointersStr, err := svc.storageSrv.ReadToString(svc.blobPointerBucketName, email)
	if err != nil {
		return nil, errors.ErrDataAccessFailure
	}
	pointers := &BlobPointers{}
	err = json.Unmarshal([]byte(pointersStr), pointers)
	return pointers, err
}

func (svc *BlobService) DestroyBlob(email, key string) error {
	var (
		keyComponents  = storage.DecomposeKey(key)
		decryptedEmail = utils.DecryptMessage(svc.emailMaskSecret, keyComponents[1])
		blobPointers   = &BlobPointers{}
	)
	// check if it belongs to bearer
	if decryptedEmail != email {
		return errors.ErrForbidden
	}
	// first, remove it from the blob pointers
	existingJSONPointers, err := svc.storageSrv.ReadToString(svc.blobPointerBucketName, email)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(existingJSONPointers), blobPointers)
	if err != nil {
		return err
	}
	var nextBlobs []Blob

	for _, blob := range blobPointers.Blobs {
		if blob.Key != key {
			nextBlobs = append(nextBlobs, blob)
		}
	}
	blobPointers.Blobs = nextBlobs
	blobPointers.Count--
	encodedPointers, err := json.Marshal(blobPointers)
	if err != nil {
		return err
	}
	_, err = svc.storageSrv.Write(svc.blobPointerBucketName, email, strings.NewReader(string(encodedPointers)))
	if err != nil {
		return err
	}
	// remove the blob, finally
	err = svc.storageSrv.Delete(svc.blobBucketName, key)
	return err

}

func InitBlobService(storageSrv *storage.StorageService, blobBucketName, blobPointerBucketName, emailMaskSecret string) *BlobService {
	return &BlobService{
		storageSrv:            storageSrv,
		blobBucketName:        blobBucketName,
		blobPointerBucketName: blobPointerBucketName,
		emailMaskSecret:       emailMaskSecret,
	}
}
