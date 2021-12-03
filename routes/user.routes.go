package routes

import (
	ctrl "aws-wallet/controller"

	"github.com/gin-gonic/gin"
)

func userRoutes(router *gin.RouterGroup) {
	router.POST("signin", ctrl.SignIn)
	router.POST("signup", ctrl.SignUp)
	router.POST("refresh", ctrl.RefreshToken)
}
