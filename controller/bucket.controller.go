package controller

import (
	"aws-wallet/services"

	"github.com/gin-gonic/gin"
)

func UploadItem(ctx *gin.Context) {
	res, status := services.UploadFile(ctx)
	if !res.Success {
		ctx.JSON(status, res)
		return
	}

	ctx.JSON(status, res)
}

