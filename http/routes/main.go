package routes

import (
	"aws-wallet/docs"
	"os"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RouterSetup() *gin.Engine {
	router := gin.Default()

	//home route
	router.GET("/", func(ctx *gin.Context){
		ctx.JSON(200, "Welcome to AWS WALLET services for next routes use prefix /api/v1")
	})

	// Api Prefix
	apiPrefix := os.Getenv("API_PREFIX")
	// version 
	apiVersion := os.Getenv("API_VERSION")

	basePath := "/" + apiPrefix + "/" + apiVersion
	
	prefix := router.Group(basePath)
	// user api routes
	userRoutes(prefix)
	bucketRoutes(prefix)

	//Swagger documentation route
	docs.SwaggerInfo.BasePath = basePath
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
