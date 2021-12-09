package services

import (
	"aws-wallet/config"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/gin-gonic/gin"
)

func UploadFile(ctx *gin.Context) (res Response, status int) {
	file, handler, fileErr := ctx.Request.FormFile("myFile")

	if fileErr != nil {
		return Response{Success: false, Message: fileErr.Error(), Data: nil}, 503
	}
	defer file.Close()

	// upload the file
	fmt.Println("Uploading:", handler.Filename)
	resp, err := config.S3session.PutObject(&s3.PutObjectInput{
	  Body: file,
	  Bucket: aws.String("roshankumar-buckets-4e91bd6f-5c81-41fd-93ff-aec9412d8953"),
	  Key: aws.String(handler.Filename),
	  ACL: aws.String(s3.BucketCannedACLPublicRead),
	})
  
	if err != nil {
	  fmt.Println(err)
	}
  
	return Response{Success: false, Message: "file uploaded successfully", Data: resp}, 503
}