package s3

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"strings"
	"time"

	"thmanyah/internal/conf"
	"thmanyah/internal/modules/cms/biz"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type s3Client struct {
	minioClient *minio.Client
	config      *conf.S3
}

func NewS3Client(ctx context.Context, c *conf.Data) (biz.S3Client, error) {
	minioClient, err := minio.New(c.S3.Host, &minio.Options{
		Creds:  credentials.NewStaticV4(c.S3.AccessKey, c.S3.SecretKey, ""),
		Region: c.S3.Region,
	})
	if err != nil {
		return nil, err
	}

	for _, bucket := range c.S3.InitialBuckets {
		exists, err := minioClient.BucketExists(ctx, bucket)
		if exists || err != nil {
			continue
		}

		policy := fmt.Sprintf(`{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Effect": "Allow",
				"Principal": {"AWS": "*"},
				"Action": ["s3:GetObject"],
				"Resource": ["arn:aws:s3:::%s/*"]
			}
		]
	}`, bucket)

		err = minioClient.MakeBucket(ctx, bucket, minio.MakeBucketOptions{
			Region: c.S3.Region,
		})
		if err != nil {
			return nil, err
		}

		_ = minioClient.SetBucketPolicy(ctx, bucket, policy)
	}

	return &s3Client{
		minioClient: minioClient,
		config:      c.S3,
	}, nil
}

func (c *s3Client) GetObject(ctx context.Context, bucket, key string) (io.Reader, error) {
	resp, err := c.minioClient.GetObject(ctx, bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *s3Client) PutObject(ctx context.Context, bucket, key string, file multipart.File) error {
	_, err := c.minioClient.PutObject(ctx, bucket, key, file, -1, minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (c *s3Client) DeleteObject(ctx context.Context, bucket, key string) error {
	err := c.minioClient.RemoveObject(ctx, bucket, key, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (c *s3Client) GetObjectSignedURL(ctx context.Context, bucket, key string) (string, error) {
	resp, err := c.minioClient.PresignedGetObject(ctx, bucket, key, time.Hour*24, url.Values{})
	if err != nil {
		return "", err
	}

	return resp.String(), nil
}

func (c *s3Client) GetObjectPublicURL(ctx context.Context, bucket, key string) string {
	if !strings.HasPrefix("/", key) {
		key = "/" + key
	}

	return fmt.Sprintf("%s/%s/%s", c.config.Host, bucket, key)
}
