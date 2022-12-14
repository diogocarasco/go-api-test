package configs

import (
	"github.com/diogocarasco/go-api-test/controllers"
	"github.com/diogocarasco/go-api-test/middlewares"

	docs "github.com/diogocarasco/go-api-test/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// GetServer returns the default application configuration
func GetServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	server := gin.New()

	docs.SwaggerInfo.BasePath = "/api/v1"
	server.MaxMultipartMemory = 8 << 20 // 8 MiB

	// default middlewares section
	server.Use(gin.Recovery())
	server.Use(middlewares.InstanaTracerMiddleware())

	server.GET("/", func(c *gin.Context) { c.JSON(200, "Hello!!") }) // HELLO!
	server.GET("/feira", controllers.GetFeiras)                      // Fetch all rows
	server.GET("/feira/:id", controllers.GetFeirasById)              // Fetch rows by ID
	server.DELETE("/feira/:id", controllers.DeleteFeiras)            // Delete row by ID
	server.POST("/feira", controllers.InsertFeiras)                  // Insert row
	server.PATCH("/feira/:id", controllers.UpdateFeiras)             // Update row
	server.POST("feira/upload", controllers.ImportFeiras)            // Import data from CSV file

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return server
}
