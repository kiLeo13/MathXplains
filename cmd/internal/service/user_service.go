package service

import (
	domain "MathXplains/internal/domain/entity"
	cognito "MathXplains/internal/infrastructure/aws/cognito"
	"errors"
	"fmt"
	identity "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"regexp"
	"strings"
)

const (
	emailPattern            = `[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	passwordSpecialsPattern = `[@$!%*?&\-_]`
	passwordNumericPattern  = `\d`

	nameMinLength = 3
	nameMaxLength = 64

	emailMinLength = 6
	emailMaxLength = 128

	passwordMinLength = 8
	passwordMaxLength = 256
)

type UserDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsAdmin   bool   `json:"is_admin"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type DeleteUserDTO struct {
	AccessToken string `json:"access_token"`
	ID          string
}

func GetAllUsers() ([]*UserDTO, *APIError) {
	users, err := userRepo.FindAll()
	if err != nil {
		fmt.Println(err)
		return nil, ErrorInternalServer
	}

	var userList []*UserDTO
	for _, u := range users {
		userList = append(userList, toUserDTO(u))
	}
	return userList, nil
}

func CreateUser(user *cognito.User) *APIError {
	cogClient := cognito.Client
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)

	if err := checkName(user.Name); err != nil {
		return err
	}

	if err := checkEmail(user.Email); err != nil {
		return err
	}

	if err := checkPassword(user.Password); err != nil {
		return err
	}

	sub, err := cogClient.SignUp(user)
	if err != nil {
		if errors.Is(err, err.(*identity.UsernameExistsException)) {
			return ErrorUserExists
		} else {
			fmt.Println(err)
			return ErrorInternalServer
		}
	}

	now := NowUTC()
	newUser := &domain.User{
		ID:        sub,
		Name:      user.Name,
		Admin:     false,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err = userRepo.Save(newUser)
	if err != nil {
		return ErrorInternalServer
	}
	return nil
}

// SignIn TODO: Improve error handling
func SignIn(u *cognito.UserLogin) (*cognito.AuthCreate, *APIError) {
	cogClient := cognito.Client

	if err := checkEmail(u.Email); err != nil {
		return nil, err
	}

	if err := checkPassword(u.Password); err != nil {
		return nil, err
	}

	auth, err := cogClient.SignIn(u)
	if err != nil {
		return nil, NewError(400, err.Error())
	}
	return auth, nil
}

func GlobalSignOut(accessToken string) *APIError {
	cogClient := cognito.Client
	err := cogClient.GlobalSignOut(accessToken)
	if err != nil {
		return NewError(400, err.Error())
	}
	return nil
}

func RefreshToken(refreshToken string) (*cognito.TokenRefreshOut, *APIError) {
	cogClient := cognito.Client
	auth, err := cogClient.RefreshToken(refreshToken)
	if err != nil {
		return nil, NewError(400, err.Error())
	}
	return auth, nil
}

// CreateConfirmation TODO: Improve error handling
func CreateConfirmation(u *cognito.UserConfirmation) *APIError {
	cogClient := cognito.Client
	err := cogClient.ConfirmAccount(u)
	if err != nil {
		return NewError(400, err.Error())
	}
	return nil
}

func ResendConfirmation(u *cognito.UserConfirmation) *APIError {
	cogClient := cognito.Client

	if err := checkEmail(u.Email); err != nil {
		return err
	}

	err := cogClient.ResendConfirmation(u)
	if err != nil {
		return NewError(400, err.Error())
	}
	return nil
}

func GetUserById(sub string) (*UserDTO, *APIError) {
	user, err := userRepo.FindById(sub)
	if err != nil {
		fmt.Println(err)
		return nil, ErrorInternalServer
	}

	if user == nil {
		return nil, ErrorUserNotFound
	}
	return toUserDTO(user), nil
}

// DeleteUserByID TODO: Improve error handling
func DeleteUserByID(du *DeleteUserDTO) *APIError {
	cogClient := cognito.Client
	user, err := userRepo.FindById(du.ID)
	if err != nil {
		fmt.Println(err)
		return ErrorInternalServer
	}
	err = cogClient.DeleteUser(du.AccessToken)
	if err != nil {
		return NewError(400, err.Error())
	}

	if user != nil {
		err = userRepo.DeleteByID(user.ID)
		if err != nil {
			fmt.Println(err)
			return ErrorInternalServer
		}
	}
	return nil
}

func IsAdmin(userId string) bool {
	admin, err := userRepo.IsAdmin(userId)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return admin
}

func toUserDTO(u *domain.User) *UserDTO {
	return &UserDTO{
		ID:        u.ID,
		Name:      u.Name,
		IsAdmin:   u.Admin,
		CreatedAt: FormatEpoch(u.CreatedAt),
		UpdatedAt: FormatEpoch(u.UpdatedAt),
	}
}

func checkName(name string) *APIError {
	length := len(name)
	if length == 0 {
		return ErrorParamNotProvided("name")
	}

	if length < nameMinLength || length > nameMaxLength {
		return ErrorInvalidNameRange
	}
	return nil
}

func checkEmail(email string) *APIError {
	length := len(email)
	if length == 0 {
		return ErrorParamNotProvided("email")
	}

	if length < emailMinLength || length > emailMaxLength {
		return ErrorInvalidEmailRange
	}

	matched, err := regexp.MatchString(emailPattern, email)
	if err != nil || !matched {
		return ErrorInvalidPattern("email")
	}
	return nil
}

func checkPassword(password string) *APIError {
	length := len(password)
	if length == 0 {
		return ErrorParamNotProvided("password")
	}

	if length < passwordMinLength || length > passwordMaxLength {
		return ErrorInvalidPasswordRange
	}

	if strings.ToLower(password) == password || strings.ToUpper(password) == password {
		return ErrorPasswordCase
	}

	if matches, _ := regexp.MatchString(passwordNumericPattern, password); !matches {
		return ErrorPasswordNumbers
	}

	if matches, _ := regexp.MatchString(passwordSpecialsPattern, password); !matches {
		return ErrorPasswordSpecialChar
	}
	return nil
}
