package main

import (
	"github.com/gin-gonic/gin"
	"go-restapi-jwt/controllers"
	"go-restapi-jwt/models"
)

func main() {
	models.ConnectDatabse()

	router := gin.Default()

	// public like base url "/api/v1"
	public := router.Group("/api/v1")

	public.POST("/register", controllers.Register)

	router.Run(":8080")

}
