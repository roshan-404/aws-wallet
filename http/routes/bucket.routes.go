package routes

import (
	"aws-wallet/http/controller"
	"aws-wallet/http/middleware"

	"github.com/gin-gonic/gin"
)

func bucketRoutes(router *gin.RouterGroup) {
	router.POST("upload", middleware.VerifyUser(), controller.UploadItem)
	router.GET("items", middleware.VerifyUser(), controller.GetAllItem)
}
