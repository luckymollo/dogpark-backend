package main

import (
	"dogpark-backend/auth"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gin-gonic/gin"
	"os"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "up",
		})
	})

	conf := &aws.Config{Region: aws.String("ca-central-1")}
	sess, err := session.NewSession(conf)
	if err != nil {
		panic(err)
	}

	cog := auth.Cognito{
		Client:      cognitoidentityprovider.New(sess),
		RegFlow:     &auth.RegFlow{},
		UserPoolId:  os.Getenv("COGNITO_USER_POOL_ID"),
		AppClientId: os.Getenv("COGNITO_APP_CLIENT_ID"),
	}

	v1 := r.Group("/v1")

	v1.POST("/register", func(c *gin.Context) {
		cog.Register(c)
	})
	return r
}
func main() {

	setupRouter().Run()
}
