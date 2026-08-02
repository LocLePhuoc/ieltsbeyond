package storage

import (
	"context"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	awscredentials "github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var Instance, _ = NewClient()

type Client struct {
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

	return &Client{presign: s3.NewPresignClient(client)}, nil
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
