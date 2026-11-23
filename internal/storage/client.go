package storage

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	awscredentials "github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/transfermanager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
)

var Instance, _ = NewClient()

type Client struct {
	client  *s3.Client
	presign *s3.PresignClient
}

func NewClient() (*Client, error) {
	credentialsProvider := awsconfig.WithCredentialsProvider(
		awscredentials.NewStaticCredentialsProvider(
			os.Getenv("STORAGE_ACCESS_KEY_ID"),
			os.Getenv("STORAGE_SECRET_ACCESS_KEY"),
			"",
		),
	)

	region := awsconfig.WithRegion(os.Getenv("STORAGE_REGION"))
	cfg, err := awsconfig.LoadDefaultConfig(context.TODO(), credentialsProvider, region)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		if endpoint := os.Getenv("STORAGE_ENDPOINT"); endpoint != "" {
			o.BaseEndpoint = aws.String(endpoint)
			o.UsePathStyle = true
		}
	})

	return &Client{client: client, presign: s3.NewPresignClient(client)}, nil
}

func (s *Client) GetObjectURL(ctx context.Context, bucket string, objectKey string) (string, error) {
	req, err := s.presign.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
	}, s3.WithPresignExpires(15*time.Minute))

	if err != nil {
		return "", err
	}
	return req.URL, nil
}

func (s *Client) UploadObject(ctx context.Context, bucket string, objectKey string, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Object upload Failed. bucket=%s, objectKey=%s, filePath=%s. %w", bucket, objectKey, filePath, err)
	}
	defer file.Close()
	_, err = s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
		Body:   file,
	})

	if err != nil {
		var apiErr smithy.APIError
		if errors.As(err, &apiErr) && apiErr.ErrorCode() == "EntityTooLarge" {
			err = fmt.Errorf("Object too large to upload. Needs multipart upload. bucket=%s, objectKey=%s, filePath=%s. %w", bucket, objectKey, filePath, err)
		} else {
			err = fmt.Errorf("Object Upload Failed. bucket=%s, objectKey=%s, filePath=%s. %w", bucket, objectKey, filePath, err)
		}
		return err
	}

	waiter := s3.NewObjectExistsWaiter(s.client)
	err = waiter.Wait(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
	}, 2*time.Minute)
	if err != nil {
		err = fmt.Errorf("Wait for file exists failed. bucket=%s, objectKey=%s, filePath=%s. %w", bucket, objectKey, filePath, err)
	}
	return err
}

func (s *Client) UploadLargeObject(ctx context.Context, bucket string, objectKey string, largeObject []byte) error {
	largeBuffer := bytes.NewBuffer(largeObject)
	var partMiBs int64 = 10

	uploader := transfermanager.New(s.client, func(o *transfermanager.Options) {
		o.PartSizeBytes = partMiBs * 1024 * 1024
	})

	_, err := uploader.UploadObject(ctx, &transfermanager.UploadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
		Body:   largeBuffer,
	})
	if err != nil {
		err = fmt.Errorf("Large object upload failed. bucket=%s, objectKey=%s. %w", bucket, objectKey, err)

	}
	return err
}

func (s *Client) GetObjectAsBytes(ctx context.Context, bucket string, objectKey string) ([]byte, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
	}
	result, err := s.client.GetObject(ctx, input)
	if err != nil {
		var noKey *types.NoSuchKey
		if errors.As(err, &noKey) {
			err = fmt.Errorf("Can't get object from bucket. No such key exists. bucket=%s, objectKey=%s.", bucket, objectKey)
		} else {
			err = fmt.Errorf("Can't get object from bucket. bucket=%s, objectKey=%s. %w", bucket, objectKey, err)
		}
		return nil, err
	}
	defer result.Body.Close()
	body, err := io.ReadAll(result.Body)
	return body, err
}
