package service

import (
	domain "MathXplains/internal/domain/entity"
	cognito "MathXplains/internal/infrastructure/aws/cognito"
	"errors"
	"fmt"
	identity "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"regexp"
	"strings"
	"time"
)

const (
	emailPattern            = `[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	passwordSpecialsPattern = `[@$!%*?&\-_]`

	nameMinLength = 3
	nameMaxLength = 64

	emailMinLength = 6
	emailMaxLength = 128

	passwordMinLength = 8
	passwordMaxLength = 128
)

type UserDTO struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	IsAdmin       bool    `json:"is_admin"`
	EmailVerified bool    `json:"email_verified"`
	VerifiedAt    *string `json:"verified_at"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
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
		var usernameExistsException *identity.UsernameExistsException
		if errors.As(err, &usernameExistsException) {
			return ErrorUserExists
		} else {
			fmt.Println(err)
			return ErrorInternalServer
		}
	}

	now := time.Now()
	err = userRepo.Save(
		sub,
		strings.TrimSpace(user.Name),
		now.Unix(),
	)
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

func RefreshToken(token string) (string, *APIError) {
	cogClient := cognito.Client
	auth, err := cogClient.RefreshToken(token)
	if err != nil {
		return "", NewError(400, err.Error())
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

func GetUserById(sub string) (*UserDTO, *APIError) {
	user, found, err := userRepo.FindById(sub)
	if err != nil {
		fmt.Println(err)
		return nil, ErrorInternalServer
	}

	if !found {
		return nil, ErrorUserNotFound
	}
	return toUserDTO(user), nil
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
	var verifiedAt *string
	if u.VerifiedAt != nil {
		form := FormatEpoch(*u.VerifiedAt)
		verifiedAt = &form
	} else {
		verifiedAt = nil
	}

	return &UserDTO{
		ID:            u.ID,
		Name:          u.Name,
		IsAdmin:       u.Admin,
		EmailVerified: u.EmailVerified,
		VerifiedAt:    verifiedAt,
		CreatedAt:     FormatEpoch(u.CreatedAt),
		UpdatedAt:     FormatEpoch(u.UpdatedAt),
	}
}

func checkName(name string) *APIError {
	length := len(name)
	if length < nameMinLength || length > nameMaxLength {
		return ErrorInvalidNameRange
	}
	return nil
}

func checkEmail(email string) *APIError {
	length := len(email)
	if length < emailMinLength || length > emailMaxLength {
		return ErrorInvalidEmailRange
	}

	matched, err := regexp.MatchString(emailPattern, email)
	if err != nil || !matched {
		return ErrorInvalidEmailPattern
	}
	return nil
}

func checkPassword(password string) *APIError {
	if len(password) < passwordMinLength || len(password) > passwordMaxLength {
		return ErrorInvalidPasswordRange
	}

	if strings.ToLower(password) == password || strings.ToUpper(password) == password {
		return ErrorPasswordCase
	}

	matched, err := regexp.MatchString(passwordSpecialsPattern, password)
	if err != nil || !matched {
		return ErrorPasswordSpecialChar
	}
	return nil
}
