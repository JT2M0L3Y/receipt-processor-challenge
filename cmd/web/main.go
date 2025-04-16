package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// define routes
	routes := router.Group("/receipts")
	{
		routes.GET("/:id/points", getPoints)
		routes.POST("/process", processReceipt)
	}

	router.Run(":8080")
}

func getPoints(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "get points",
	})
}

func processReceipt(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "Here is the receipt id: " + id,
	})
}
