package utils

import (
	"aws-wallet/database/models"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var td = &models.TokenDetails{}

func CreateToken(username string, id string) (*models.TokenDetails, error) {
	
	td.AtExpires = time.Now().Add(time.Minute * 40).Unix()
	td.RtExpires = time.Now().Add(time.Minute * 120).Unix()

	var err error
	//Creating Access Token
	fmt.Println(username, id)
	atClaims := jwt.MapClaims{}
	atClaims["expiresAt"] = td.AtExpires
	atClaims["username"] = username
	atClaims["userId"] = id
	atClaims["authorized"] = true
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	

	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["expiresAt"] = td.RtExpires
	rtClaims["username"] = username
	rtClaims["userId"] = id
	rtClaims["authorized"] = true
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func verifyToken(data string, secret []byte) (*jwt.Token, string) {
	token, err := jwt.Parse(data, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, "Invalid Token!"
	}

	return token, ""
}

func VerifyAccessToken(ctx *gin.Context) (jwt.Claims, string) {
	if data := ctx.Request.Header.Get("Authorization"); data != "" {
		token, err := verifyToken(data, []byte(os.Getenv("ACCESS_SECRET")))
		if err != "" {
			return nil, err
		}

		// extract expiresTime from token
		ext := token.Claims.(jwt.MapClaims)
		expiresTime := ext["expiresAt"]
		username := ext["username"]
		userId := ext["userId"]

		ctx.Request.Header.Set("username", fmt.Sprintf("%v", username))
		ctx.Request.Header.Set("userId", fmt.Sprintf("%v", userId))

		if int64(expiresTime.(float64)) < time.Now().Unix() {
			return nil, "Token expired!"
		}
		return token.Claims, ""
	} else {
		return nil, "You are not logged In"
	}
}
