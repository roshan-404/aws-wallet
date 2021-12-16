package main

import (
	"aws-wallet/config"
	"aws-wallet/routes"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	// "github.com/joho/godotenv"
)

func main() {
	fmt.Println("AWS wallet starting ...")
	// changes from secondary branch

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
