package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "up",
		})
	})
	return r
}
func main() {
	setupRouter().Run()
}
