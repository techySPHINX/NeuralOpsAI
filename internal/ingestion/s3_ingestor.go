package ingestion

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type S3Ingestor struct {
	client *minio.Client
	bucket string
}

func NewS3Ingestor(endpoint, accessKeyID, secretAccessKey, bucketName string) (*S3Ingestor, error) {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false, // Set to true for HTTPS
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create minio client: %w", err)
	}

	// Check if the bucket exists, if not, create it
	found, err := minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		return nil, fmt.Errorf("failed to check bucket existence: %w", err)
	}
	if !found {
		log.Printf("Bucket %s not found, creating it...", bucketName)
		err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to create bucket %s: %w", bucketName, err)
		}
		log.Printf("Bucket %s created successfully.", bucketName)
	}

	return &S3Ingestor{
		client: minioClient,
		bucket: bucketName,
	}, nil
}

func (s *S3Ingestor) ListObjects() ([]string, error) {
	var objectNames []string
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	objectCh := s.client.ListObjects(ctx, s.bucket, minio.ListObjectsOptions{Recursive: true})
	for object := range objectCh {
		if object.Err != nil {
			return nil, fmt.Errorf("error listing objects: %w", object.Err)
		}
		objectNames = append(objectNames, object.Key)
	}
	return objectNames, nil
}

// Placeholder for reading object content
func (s *S3Ingestor) GetObjectContent(objectName string) ([]byte, error) {
	// In a real scenario, you would read the object content here.
	// For now, we'll just return a dummy message.
	log.Printf("Attempting to get content for object: %s", objectName)
	return []byte(fmt.Sprintf("Content of %s (placeholder)", objectName)), nil
}
