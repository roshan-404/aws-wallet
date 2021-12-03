package routes

import (
	ctrl "aws-wallet/controller"
	"aws-wallet/middleware"

	"github.com/gin-gonic/gin"
)

func bucketRoutes(router *gin.RouterGroup) {
	router.POST("upload", middleware.VerifyUser(), ctrl.UploadItem)
}
