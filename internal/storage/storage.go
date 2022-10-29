// ref:
// http://www.inanzzz.com/index.php/post/egpk/a-simple-aws-s3-example-with-golang-suing-localstack
// https://github.com/johncalvinroberts/location-pingee/blob/main/src/storage/storage.service.ts
package storage

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	DELIMITER = "::"
)

type StorageService struct {
	timeout    time.Duration
	client     *s3.S3
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
}

func (svc *StorageService) Write(ctx context.Context, bucket, key string, body io.Reader) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, svc.timeout)
	defer cancel()

	res, err := svc.uploader.UploadWithContext(ctx, &s3manager.UploadInput{
		Body:   body,
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return "", fmt.Errorf("failed to write to storage: %w", err)
	}

	return res.Location, nil
}

func (svc *StorageService) Read(ctx context.Context, bucket, key string, body io.WriterAt) error {
	if _, err := svc.downloader.DownloadWithContext(ctx, body, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}); err != nil {
		return fmt.Errorf("failed to read from storage: %w", err)
	}

	return nil
}

func (svc *StorageService) Delete(ctx context.Context, bucket, key string) error {
	if _, err := svc.client.DeleteObjectWithContext(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}); err != nil {
		return fmt.Errorf("failed to delete from storage: %w", err)
	}

	if err := svc.client.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}); err != nil {
		return fmt.Errorf("wait: %w", err)
	}

	return nil
}

// Should take a cursor + max keys
// returns list of keys
func (svc *StorageService) List(ctx context.Context, bucket string, cursor *string, maxKeys int, prefix *string) ([]string, error) {
	ctx, cancel := context.WithTimeout(ctx, svc.timeout)
	defer cancel()
	res, err := svc.client.ListObjectsV2WithContext(ctx, &s3.ListObjectsV2Input{
		Bucket:     aws.String(bucket),
		MaxKeys:    aws.Int64(int64(maxKeys)),
		StartAfter: cursor,
		Prefix:     prefix,
	})
	if err != nil {
		return nil, fmt.Errorf("list: %w", err)
	}
	keys := make([]string, len(res.Contents))
	for _, item := range res.Contents {
		keys = append(keys, *item.Key)
	}
	return keys, nil
}

func ComposeKey(comps ...string) string {
	key := ""
	for _, s := range comps {
		key = fmt.Sprintf("%s%s%s", key, DELIMITER, s)
	}
	return key
}

func DecomposeKey(key string) []string {
	return strings.Split(key, DELIMITER)
}

func InitStorageService(session *session.Session, timeout int) *StorageService {
	return &StorageService{
		timeout:    time.Duration(timeout) * time.Millisecond,
		client:     s3.New(session),
		uploader:   s3manager.NewUploader(session),
		downloader: s3manager.NewDownloader(session),
	}
}
