package services

import (
	"aws-wallet/http/repository"
	"aws-wallet/http/utils"

	"github.com/gin-gonic/gin"
)

func UploadFile(ctx *gin.Context) (res Response, status int) {

	//getting the bucket name from the headers
	username := ctx.Request.Header.Get("username")
    userId := ctx.Request.Header.Get("userId")

    form, err := ctx.MultipartForm()
    if err != nil {
        return Response{Success: false, Message: err.Error(), Data: nil}, 415
    }
    files := form.File["myFile"]

    for _, file := range files {
        f, fErr := file.Open()
        if fErr != nil {
            return Response{Success: false, Message: err.Error(), Data: nil}, 415
        }
        defer f.Close()

        bucketName := utils.CreateBucketName(username, userId)
        err = repository.UploadObject(bucketName, file.Filename, f)

        if err != nil {
            return Response{Success: false, Message: err.Error(), Data: nil}, 400
        }

    }
  
	return Response{Success: false, Message: "file uploaded successfully", Data: nil}, 200
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