package repository

import (
	"aws-wallet/database/config"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

func CreateBucket(bucketName string) (resp *s3.CreateBucketOutput, err error) {
	resp, err = config.S3session.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
			LocationConstraint: aws.String(os.Getenv("AWS_REGION")),
		},
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
				return nil, fmt.Errorf("bucket name is already in use")

			case s3.ErrCodeBucketAlreadyOwnedByYou:
				return nil, fmt.Errorf("bucket exists and is owned by you")

			default:
				return nil, err
			}
		}
	}
	return resp, nil
}

func UploadObject(bucketName string, fileName string, file multipart.File)  (err error) {
	fmt.Println("Uploading:", fileName)
	_, err = config.S3session.PutObject(&s3.PutObjectInput{
		Body:   file,
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		ACL:    aws.String(s3.BucketCannedACLPublicRead),
	})

	return err
}

func GetAllObject(bucketName string) (resp *s3.ListObjectsV2Output, err error){
	resp, err = config.S3session.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	return resp, err
}
