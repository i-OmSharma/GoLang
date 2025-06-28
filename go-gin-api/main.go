package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//basic GET route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Om from Gin",
		})
	})

	// Running the server
	router.Run(":8080")

}
