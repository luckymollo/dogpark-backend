package auth

import (
	"dogpark-backend/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Cognito struct {
	Client      *cognitoidentityprovider.CognitoIdentityProvider
	RegFlow     *RegFlow
	UserPoolId  string
	AppClientId string
}

type RegFlow struct {
	Username string
	Email    string
}

func (cog *Cognito) Register(context *gin.Context) {
	var user model.User
	err := context.BindJSON(&user)
	if err != nil {
		context.String(http.StatusBadRequest, err.Error())
	}

	registerUser := &cognitoidentityprovider.SignUpInput{
		Username: aws.String(user.UserName),
		Password: aws.String(user.Password),
		ClientId: aws.String(cog.AppClientId),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(user.Email),
			},
		},
	}

	_, err = cog.Client.SignUp(registerUser)

	if err != nil {
		panic(err)
	}

	context.String(http.StatusCreated, "user is created")

}
