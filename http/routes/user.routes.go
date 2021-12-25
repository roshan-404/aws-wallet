package routes

import (
	ctrl "aws-wallet/http/controller"

	"github.com/gin-gonic/gin"
)

func userRoutes(router *gin.RouterGroup) {
	router.POST("signin", ctrl.SignIn)
	router.POST("signup", ctrl.SignUp)
	router.POST("refreshToken", ctrl.RefreshToken)
	router.POST("verifyotp", ctrl.VerifyOTP)
}
