package main

import (
	"aws-wallet/config"
	"aws-wallet/routes"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("AWS wallet starting ...")

	// load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
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