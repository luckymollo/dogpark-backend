package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "up",
		})
	})

	v1 := r.Group("/v1")

	v1.GET("/register")
	return r
}
func main() {
	conf := &aws.Config{Region: aws.String("ca-central-1")}
	sess, err := session.NewSession(conf)
	if err != nil {
		panic(err)
	}

	setupRouter().Run()
}
