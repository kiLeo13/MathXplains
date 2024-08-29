package routes

import (
	cognito "MathXplains/internal/infrastructure/aws/cognito"
	"MathXplains/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func GetUsers(c echo.Context) error {
	userId := c.Request().Header.Get("Sub")
	admin := service.IsAdmin(userId)
	if !admin {
		return c.JSON(http.StatusForbidden, service.ErrorMissingAdmin)
	}

	users, err := service.GetAllUsers()
	if err != nil {
		return c.JSON(err.Status, err)
	}
	return c.JSON(http.StatusOK, &R{"users": users})
}

func CreateUser(c echo.Context) error {
	body := cognito.User{}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, service.ErrorMalformedJSON)
	}

	err := service.CreateUser(&body)
	if err != nil {
		return c.JSON(err.Status, err)
	}
	return c.NoContent(http.StatusOK)
}

func LoginUser(c echo.Context) error {
	req := cognito.UserLogin{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, service.ErrorMalformedJSON)
	}

	create, err := service.SignIn(&req)
	if err != nil {
		return c.JSON(err.Status, err)
	}
	return c.JSON(http.StatusOK, create)
}

func RefreshToken(c echo.Context) error {
	req := make(map[string]string)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, service.ErrorMalformedJSON)
	}
	token, ok := req["refresh_token"]
	if !ok {
		return c.JSON(http.StatusBadRequest, service.ErrorParamNotProvided("refresh_token"))
	}

	tokens, err := service.RefreshToken(token)
	if err != nil {
		return c.JSON(err.Status, err)
	}
	return c.JSON(http.StatusOK, tokens)
}

func ConfirmAccount(c echo.Context) error {
	body := cognito.UserConfirmation{}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, service.ErrorMalformedJSON)
	}

	err := service.CreateConfirmation(&body)
	if err != nil {
		return c.JSON(err.Status, err)
	}
	return c.NoContent(http.StatusOK)
}

func ResendConfirmation(c echo.Context) error {
	body := cognito.UserConfirmation{}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, service.ErrorMalformedJSON)
	}

	err := service.ResendConfirmation(&body)
	if err != nil {
		return c.JSON(err.Status, err)
	}
	return c.NoContent(http.StatusOK)
}

func GetUserByID(c echo.Context) error {
	userId := c.Request().Header.Get("Sub")
	id := c.Param("id")
	if id == "@me" {
		id = userId
	}

	if !service.IsAdmin(userId) && id != userId {
		return c.JSON(http.StatusForbidden, service.ErrorMissingAdmin)
	}

	user, err := service.GetUserById(id)
	if err != nil {
		return c.JSON(err.Status, err)
	}
	return c.JSON(http.StatusOK, user)
}

func DeleteSelfUser(c echo.Context) error {
	userId := c.Request().Header.Get("Sub")
	del := service.DeleteUserDTO{ID: userId}
	if err := c.Bind(&del); err != nil {
		return c.JSON(http.StatusBadRequest, service.ErrorMalformedJSON)
	}

	if strings.TrimSpace(del.AccessToken) == "" {
		return c.JSON(http.StatusBadRequest, service.ErrorParamNotProvided("access_token"))
	}

	err := service.DeleteUserByID(&del)
	if err != nil {
		return c.JSON(err.Status, err)
	}
	return c.NoContent(http.StatusOK)
}
