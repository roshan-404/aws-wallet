package services

import (
	"aws-wallet/http/repository"
	"aws-wallet/http/utils"

	"github.com/gin-gonic/gin"
)

func UploadFile(ctx *gin.Context) (res Response, status int) {
	file, handler, fileErr := ctx.Request.FormFile("myFile")

	if fileErr != nil {
		return Response{Success: false, Message: fileErr.Error(), Data: nil}, 415
	}
	defer file.Close()

	//getting the bucket name from the headers 
	username := ctx.Request.Header.Get("username")
	userid := ctx.Request.Header.Get("userId")
	bucketName := utils.CreateBucketName(username, userid)

	
	// upload the file
	resp,err := repository.UploadObject(bucketName, handler.Filename, file)
  
	if err != nil {
		return Response{Success: false, Message: err.Error(), Data: nil}, 400
	}
  
	return Response{Success: false, Message: "file uploaded successfully", Data: resp}, 200
}

func GetAllItem(ctx *gin.Context) (res Response, status int) {
	//getting the bucket name from the headers 
	username := ctx.Request.Header.Get("username")
	userid := ctx.Request.Header.Get("userId")
	bucketName := utils.CreateBucketName(username, userid)

	//get all items
	resp, err := repository.GetAllObject(bucketName)

	if err != nil {
		return Response{Success: false, Message: err.Error(), Data: nil}, 400
	}
	return Response{Success: false, Message: "file uploaded successfully", Data: resp},  200
}