package routes

import (
	"MathXplains/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetProfessors(c echo.Context) error {
	knownOnly := c.QueryParam("known") == "true"
	professors, err := service.GetProfessors(knownOnly)
	if err != nil {
		return c.JSON(err.Status, err)
	}

	resp := R{
		"professors": professors,
	}
	return c.JSON(http.StatusOK, &resp)
}
