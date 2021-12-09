package routes

import (
	"aws-wallet/controller"
	"aws-wallet/middleware"

	"github.com/gin-gonic/gin"
)

func bucketRoutes(router *gin.RouterGroup) {
	router.POST("upload", middleware.VerifyUser(), controller.UploadItem)
}
