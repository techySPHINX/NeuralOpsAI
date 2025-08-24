package clients

import (
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinIOClient struct {
	client *minio.Client
	bucket string // Default bucket
}

func NewMinIOClient(endpoint, accessKeyID, secretAccessKey string, useSSL bool, bucket string) (*MinIOClient, error) {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create minio client: %w", err)
	}

	return &MinIOClient{client: minioClient, bucket: bucket}, nil
}

func (c *MinIOClient) PutObject(ctx context.Context, objectName string, reader io.Reader, objectSize int64, contentType string) (minio.UploadInfo, error) {
	return c.client.PutObject(ctx, c.bucket, objectName, reader, objectSize, minio.PutObjectOptions{ContentType: contentType})
}

func (c *MinIOClient) GetObject(ctx context.Context, objectName string) (*minio.Object, error) {
	return c.client.GetObject(ctx, c.bucket, objectName, minio.GetObjectOptions{})
}
