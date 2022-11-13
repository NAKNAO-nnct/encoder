package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

type AwsConfig struct {
	Region          string
	AccessKey       string
	SecretAccessKey string
}

type AwsClient struct {
	config *aws.Config
}

func NewAwsClient(c *AwsConfig) *AwsClient {
	return &AwsClient{
		&aws.Config{
			Region:      aws.String(c.Region),
			Credentials: credentials.NewStaticCredentials(c.AccessKey, c.SecretAccessKey, ""),
		},
	}
}

func (c *AwsClient) NewSession() *session.Session {
	return session.Must(session.NewSession(c.config))
}

func (c *AwsClient) S3DownloadFile(bucket, key, savepath string) error {
	downloader := s3manager.NewDownloader(c.NewSession())

	file, err := os.Create(savepath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		})

	if err != nil {
		return err
	}

	return nil
}
