package main

import (
	"github.com/gin-gonic/gin"
	"go-restapi-jwt/controllers"
	"go-restapi-jwt/middlewares"
	"go-restapi-jwt/models"
)

func main() {
	models.ConnectDatabse()

	router := gin.Default()

	// public like base url "/api/v1"
	public := router.Group("/api/v1")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := router.Group("/api/v1/admin")
	protected.Use(middlewares.JwtAuthMiddleware(), middlewares.CORSMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	router.Run(":8080")

}
