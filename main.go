package main

import (
	"example/wongnai-test/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	data := services.GetData()

	r.GET("/covid/summary", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Province": services.FilterByProvince(data),
			"AgeGroup": services.FilterByAge(data),
		})
	})

	r.Run(":3000")
}
