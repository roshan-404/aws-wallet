package services

import (
	"aws-wallet/database/models"
	"aws-wallet/http/repository"
	"aws-wallet/http/utils"
	"aws-wallet/http/validator"
	"aws-wallet/http/verification"
	"fmt"
	"log"
	"os"
	"strconv"

	uuid "github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Response models.Response

func VerifyUser(user *models.SignIn) (res Response, status int) {
	//validate user
	vErr := validator.SigninValidator(user)
	if vErr != nil {
		return Response{Success: false, Message: vErr.Error(), Data: nil}, 400
	}

	item, err := repository.GetItem(user.Username)
	if err != nil {
		return Response{Success: true, Message: err.Error(), Data: nil}, 400
	}

	if err = bcrypt.CompareHashAndPassword([]byte(fmt.Sprintf("%v", item["password"])), []byte(user.Password)); err != nil {
		return Response{Success: false, Message: "Wrong password!", Data: nil}, 401
	}

	token, err := utils.CreateToken(user.Username, fmt.Sprintf("%v", item["id"]))
	if err != nil {
		return Response{Message: "Something went wrong!", Data: nil, Success: false}, 500
	}

	
	tokens := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
		"access-token-expiry" : os.Getenv("ACCESS_TOKEN_EXPIRY"),
		"refresh-token-expiry" : os.Getenv("REFRESH_TOKEN_EXPIRY"),

	}

	return Response{Success: true, Message: "SignIn successful!", Data: tokens}, 200
}

func CreateUser(user *models.User) (res Response, status int) {
	//validate the fields 
	vErr := validator.SignUpValidator(user)
	if vErr != nil {
		return Response{Success: false, Message: vErr.Error(), Data: nil}, 400
	}
	// check if user exists
	exists, err := repository.FindItem(user.Username)
	if err != nil {
		return Response{Success: false, Message: err.Error(), Data: nil}, 500
	}
	if exists {
		return Response{Success: false, Message: "User already exists!", Data: nil}, 500
	}

	// generate new user id
	newId, _ := uuid.NewV4()
	// hash password
	hash, hashErr := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if hashErr != nil {
		return Response{Success: false, Message: "Something went wrong!", Data: nil}, 500
	}

	// assign id & hash password
	user.Id = newId.String()
	user.Password = string(hash)
	// u := splitString(user.Username, " ")
	bucketName := user.Username + "-buckets-" + user.Id
	fmt.Println(bucketName)
	// createing new bucket
	_, BucketErr := repository.CreateBucket(bucketName)
	if BucketErr != nil {
		return Response{Success: false, Message: BucketErr.Error(), Data: nil}, 400
	}

	// Add new user in database
	err = repository.PutItem(user)
	if err != nil {
		return Response{Success: false, Message: err.Error(), Data: nil}, 400
	}

	otp, err := utils.CreateOTP(user.Id)
	if err != nil {
		return Response{Success: false, Message: err.Error(), Data: nil}, 400
	}

	isSend := verification.SendSMS(user.PhoneNumber, "Welcome to awsCloud services. OTP - "+strconv.Itoa(otp))
	if !isSend {
		log.Println("Message failed to send - ", user.PhoneNumber)
	}

	return Response{Success: true, Message: "SignUp successful!", Data: user}, 200
}

func VerifyOTP(otp *models.OTP) (res Response, status int) {
	err := verification.VerifyOTP(otp.Id, otp.OTP)
	if err != nil {
		return Response{Success: false, Message: err.Error(), Data: nil}, 400
	}

	dberr := repository.UpadteItem(otp.Id, "true")
	if dberr != nil {
		return Response{Success: false, Message: dberr.Error(), Data: nil}, 400
	}

	return Response{Success: true, Message: "OTP varification successfully!", Data: nil}, 200
}
