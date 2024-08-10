package cognito_client

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"os"
)

var Client CognitoInterface

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CognitoInterface interface {
	SignUp(user *User) (string, error)
	ConfirmAccount(user *UserConfirmation) error
	SignIn(user *UserLogin) (*AuthCreate, error)
	RefreshToken(refreshToken string) (string, error)
	GetUserByToken(token string) (*cognito.GetUserOutput, error)
}

type cognitoClient struct {
	cognitoClient *cognito.CognitoIdentityProvider
	appClientId   string
}

func NewCognitoClient(appClientId string) {
	config := &aws.Config{
		Region:                        aws.String(os.Getenv("AWS_REGION")),
		CredentialsChainVerboseErrors: aws.Bool(true),
	}
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}

	client := cognito.New(sess)

	Client = &cognitoClient{
		cognitoClient: client,
		appClientId:   appClientId,
	}
}

func (c *cognitoClient) SignUp(user *User) (sub string, err error) {
	userCognito := &cognito.SignUpInput{
		ClientId: aws.String(c.appClientId),
		Username: aws.String(user.Email),
		Password: aws.String(user.Password),
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  aws.String("given_name"),
				Value: aws.String(user.Name),
			},
			{
				Name:  aws.String("email"),
				Value: aws.String(user.Email),
			},
		},
	}
	out, err := c.cognitoClient.SignUp(userCognito)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return *out.UserSub, nil
}

type UserConfirmation struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func (c *cognitoClient) ConfirmAccount(user *UserConfirmation) error {
	confirmationInput := &cognito.ConfirmSignUpInput{
		Username:         aws.String(user.Email),
		ConfirmationCode: aws.String(user.Code),
		ClientId:         aws.String(c.appClientId),
	}
	_, err := c.cognitoClient.ConfirmSignUp(confirmationInput)
	if err != nil {
		return err
	}
	return nil
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthCreate struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func (c *cognitoClient) RefreshToken(refreshToken string) (string, error) {
	authInput := &cognito.InitiateAuthInput{
		AuthFlow: aws.String("REFRESH_TOKEN_AUTH"),
		AuthParameters: aws.StringMap(map[string]string{
			"REFRESH_TOKEN": refreshToken,
		}),
		ClientId: aws.String(c.appClientId),
	}
	result, err := c.cognitoClient.InitiateAuth(authInput)
	if err != nil {
		return "", err
	}
	return *result.AuthenticationResult.IdToken, nil
}

func (c *cognitoClient) SignIn(user *UserLogin) (*AuthCreate, error) {
	authInput := &cognito.InitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: aws.StringMap(map[string]string{
			"USERNAME": user.Email,
			"PASSWORD": user.Password,
		}),
		ClientId: aws.String(c.appClientId),
	}
	result, err := c.cognitoClient.InitiateAuth(authInput)
	if err != nil {
		return nil, err
	}
	return &AuthCreate{
		Token:        *result.AuthenticationResult.IdToken,
		RefreshToken: *result.AuthenticationResult.RefreshToken,
	}, nil
}

func (c *cognitoClient) GetUserByToken(token string) (*cognito.GetUserOutput, error) {
	input := &cognito.GetUserInput{
		AccessToken: aws.String(token),
	}
	result, err := c.cognitoClient.GetUser(input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type COGSelfUser struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}
