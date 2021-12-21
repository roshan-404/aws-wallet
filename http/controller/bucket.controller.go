package controller

import (
	"aws-wallet/http/services"

	"github.com/gin-gonic/gin"
)

// Upload controller
// @Summary authenticated users can upload files.
// @Description You need to signedIn and give a Token in headers then "upload item" will execute.
// @Tags File Upload
// @Accept  json
// @Produce  json
// @Router /upload [post]
func UploadItem(ctx *gin.Context) {
	res, status := services.UploadFile(ctx)
	if !res.Success {
		ctx.JSON(status, res)
		return
	}

	ctx.JSON(status, res)
}

// Upload controller
// @Summary authenticated users can see all there files.
// @Description You need to signedIn and give a Token in headers then "Get All Items" will execute.
// @Tags File Upload
// @Accept  json
// @Produce  json
// @Router /items [get]
func GetAllItem(ctx *gin.Context) {
	res, status := services.GetAllItem(ctx)
	if !res.Success {
		ctx.JSON(status, res)
		return
	}

	ctx.JSON(status, res)
}


