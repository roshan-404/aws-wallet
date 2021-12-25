package models

type User struct {
	Id       string `dynamodbav:"id"`
	Username string `dynamodbav:"username" validate:"required,username,isValid"`
	Password string `dynamodbav:"password" validate:"required,password"`
	PhoneNumber string `dynamodbav:"phone_number" validate:"required,phone,isNum"`
	Verified    bool   `dynamodbav:"verified"`
}

type SignIn struct {
	Username string `dynamodbav:"username" validate:"required"`
	Password string `dynamodbav:"password" validate:"required"`
}


type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}


type TokenDetails struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AtExpires    int64  `json:"at_expires"`
	RtExpires    int64  `json:"rt_expires"`
}

type OTP struct {
	Id  string `json:"id"`
	OTP string `json:"otp"`
}

