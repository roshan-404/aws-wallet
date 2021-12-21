package main

import (
	"aws-wallet/database/config"
	"aws-wallet/http/routes"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// @title AWS Wallet API Documentation.
// @version 1.0.0
// @description A service where users can register themselves and can have their personal account to store their files.
// @termsOfService http://swagger.io/terms/

// @contact.name Roshan Kumar Ojha
// @contact.email roshankumarojha04@gmail.com

// @host localhost:3000
// @BasePath /api/v1

// var addr = flag.String("addr", ":8080", "http service address")

func main() {
	fmt.Println("AWS wallet starting ...")

	// load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// stabling the database
	_, dbErr := config.ConnectionDB()
	if dbErr != nil {
		panic("Error in stabling database connection")
	}

	// start server
	router := routes.RouterSetup()
	router.Run(":" + os.Getenv("PORT"))
}
