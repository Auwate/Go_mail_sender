package utils

import (
	"bytes"
	"context"
	"errors"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
)

type AWS struct {
	BucketName string
	Ctx        context.Context
	Client     *s3.Client
}

var currConfig = AWS{
	BucketName: "",
	Ctx:        nil,
	Client:     nil,
}

func UploadFile(logFileLocation string) error {

	data, _ := os.ReadFile(logFileLocation)

	_, err := currConfig.Client.PutObject(currConfig.Ctx, &s3.PutObjectInput{
		Bucket: aws.String(currConfig.BucketName),
		Key:    aws.String("log.txt"),
		Body:   bytes.NewReader(data),
	})

	return err

}

func HandleBucketCreation(bucketName string, region string) error {

	currConfig.BucketName = bucketName
	currConfig.Ctx = context.TODO()

	config, err := config.LoadDefaultConfig(currConfig.Ctx, config.WithRegion(region))
	if err != nil {
		return err
	}

	currConfig.Client = s3.NewFromConfig(config)

	exists, err := bucketExists(currConfig.BucketName, currConfig.Ctx, currConfig.Client)
	if err != nil {
		return err
	} else if exists {
		log.Println("Bucket was already created.")
		return nil
	}

	if err := createS3Bucket(currConfig.BucketName, currConfig.Ctx, currConfig.Client); err != nil {
		return err
	}

	return nil

}

func bucketExists(bucketName string, ctx context.Context, client *s3.Client) (bool, error) {

	_, err := client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})

	exists := true

	if err != nil {

		var apiError smithy.APIError
		if errors.As(err, &apiError) {
			switch apiError.(type) {
			case *types.NotFound:
				exists = false
				err = nil
			default:
				exists = false
			}
		}

	}

	return exists, err

}

func createS3Bucket(bucketName string, ctx context.Context, client *s3.Client) error {

	_, err := client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})

	return err

}
