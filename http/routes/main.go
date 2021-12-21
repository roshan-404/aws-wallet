package routes

import (
	"os"

	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {
	router := gin.Default()

	//home route
	router.POST("/", func(ctx *gin.Context){
		ctx.JSON(200, "Welcome to AWS WALLET services for next routes use prefix /api/v1")
	})

	// Api Prefix
	apiPrefix := os.Getenv("API_PREFIX")
	// version 1
	apiVersion := os.Getenv("API_VERSION")

	prefix := router.Group("/" + apiPrefix + "/" + apiVersion)
	// user api routes
	userRoutes(prefix)
	bucketRoutes(prefix)

	return router
}
