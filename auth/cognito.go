package auth

import (
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gin-gonic/gin"
)

type Cognito struct {
	Client      *cognitoidentityprovider.CognitoIdentityProvider
	RegFlow     *regFlow
	UserPoolId  string
	AppClientId string
}

type regFlow struct {
	Username string
	Email    string
}

func (cog *Cognito) Register(context *gin.Context) {

}
