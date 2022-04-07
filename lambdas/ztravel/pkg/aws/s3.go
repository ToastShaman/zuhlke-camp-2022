package aws

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Storage interface {
	PersistJSON(ctx context.Context, key string, content string) error
}

type S3Storage struct {
	Bucket string
	Client *s3.Client
}

func (s *S3Storage) PersistJSON(ctx context.Context, key string, content string) (err error) {
	_, err = s.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.Bucket),
		Key:         aws.String(key),
		ContentType: aws.String("application/json"),
		Body:        strings.NewReader(content),
	})

	return
}

func NewS3Client(cfg aws.Config) *s3.Client {
	return s3.NewFromConfig(cfg)
}

func NewS3Storage(bucket string, client *s3.Client) *S3Storage {
	return &S3Storage{
		Bucket: bucket,
		Client: client,
	}
}
