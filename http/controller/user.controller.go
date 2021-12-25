package controller

import (
	"aws-wallet/database/models"
	"aws-wallet/http/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LogIn controller
// @Summary LogIn with credentials.
// @Description A registered user can login with their credentials.
// @Tags LogIn
// @Accept  json
// @Produce  json
// @Param user body models.SignIn true "LogIn User"
// @Success 200 {object} models.SignIn
// @Failure 401 {object} object
// @Router /signin [post]
func SignIn(ctx *gin.Context) {
	var user models.SignIn

	if credErr := ctx.ShouldBindJSON(&user); credErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "Invalid input provided")
		return
	}

	res, status := services.VerifyUser(&user)
	if !res.Success {
		ctx.JSON(status, res)
		return
	}

	ctx.JSON(status, res)
}

// Sign Up controller
// @Summary Sign Up with credentials.
// @Description A new user can sign up with their email & password.
// @Tags Sign Up
// @Accept  json
// @Produce  json
// @Param user body models.User true "Sign Up User"
// @Success 200 {object} models.User
// @Failure 401 {object} object
// @Router /signup [post]
func SignUp(ctx *gin.Context) {
	var user models.User

	if credErr := ctx.ShouldBindJSON(&user); credErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "Invalid input provided")
		return
	}

	res, status := services.CreateUser(&user)
	if !res.Success {
		ctx.JSON(status, res)
		return
	}

	ctx.JSON(status, res)
}


// Refresh token controller
// @Summary Verify token & create a new token.
// @Description You need to signedIn and give a Token in headers then "Refresh Token" will execute.
// @Tags Refresh token
// @Accept  json
// @Produce  json
// @Router /refreshToken [post]
func RefreshToken(ctx *gin.Context) {

}


// Verify OTP controller
// @Summary Verify OTP
// @Description You need to provide a Token in body.
// @Tags Verify OTP
// @Accept  json
// @Produce  json
// @Param user body models.OTP true "Verify OTP"
// @Success 200 {object} models.OTP
// @Failure 400 {object} object
// @Router /verifyotp [post]
func VerifyOTP(ctx *gin.Context) {
	var otp models.OTP

	if credErr := ctx.ShouldBindJSON(&otp); credErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, models.Response{Success: false, Message: "Invalid input provided", Data: nil})
		return
	}

	res, status := services.VerifyOTP(&otp)
	if !res.Success {
		ctx.JSON(status, res)
		return
	}

	ctx.JSON(status, res)
}
