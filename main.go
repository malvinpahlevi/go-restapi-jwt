package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// public like base url "/api/v1"
	public := router.Group("/api/v1")

	public.POST("/register", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "this is a register endpoint",
		})
	})

	router.Run(":8080")

}
